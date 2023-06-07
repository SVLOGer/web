package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type createPostRequest struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	PostIMG     string `json:"postIMG"`
	PostName    string `json:"postIMGName"`
	Author      string `json:"authorName"`
	AuthorIMG   string `json:"authorIMG"`
	AuthorName  string `json:"authorIMGName"`
	PreviewIMG  string `json:"previewIMG"`
	PreviewName string `json:"previewIMGName"`
	PublishDate string `json:"publishDate"`
	Content     string `json:"content"`
}

type indexPage struct {
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type featuredPostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"modifier"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	Alt         string `db:"title"`
	AuthorAlt   string `db:"author"`
	PostID      string `db:"post_id"`
}

type mostRecentPostData struct {
	Image       string `db:"post_img"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	Alt         string `db:"title"`
	AuthorAlt   string `db:"author"`
	PostID      string `db:"post_id"`
}

type postData struct {
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	Image    string `db:"preview_img"`
	Content  string `db:"content"`
}

type adminPage struct {
	Title    string
	Subtitle string
}

type loginPage struct {
	Title    string
	Subtitle string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		mostRecentPostsData, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			FeaturedPosts:   featuredPostsData,
			MostRecentPosts: mostRecentPostsData,
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			preview_img,
			content
		FROM
		    post
		WHERE
			post_id = ?
	`

	var post postData

	err := db.Get(&post, query, postID)
	if err != nil {
		return postData{}, err
	}

	return post, nil
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
        	subtitle,
        	author,
			author_url,
        	modifier,
        	publish_date,
			post_id
		FROM
			post
		WHERE featured = 1
	`

	var posts []featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
		    post_img,
        	title,
        	subtitle,
        	author,
        	author_url,
        	publish_date,
			post_id
		FROM
			post
		WHERE featured = 0
	`

	var most []mostRecentPostData

	err := db.Select(&most, query)
	if err != nil {
		return nil, err
	}

	return most, nil
}

func admin(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := adminPage{
		Title:    "Escape",
		Subtitle: "Biba i Boba",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := loginPage{
		Title:    "Escape",
		Subtitle: "Biba i Boba",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		log.Println("Request completed successfully")
	}
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	authorImgName, err := saveImg(req.AuthorName, req.AuthorIMG)
	if err != nil {
		log.Println(err)
		return err
	}
	postImgName, err := saveImg(req.PostName, req.PostIMG)
	if err != nil {
		log.Println(err)
		return err
	}
	previewImgName, err := saveImg(req.PreviewName, req.PreviewIMG)
	if err != nil {
		log.Println(err)
		return err
	}
	const query = `
		INSERT INTO post
		(
			title,
			subtitle,
			preview_img,
			post_img,
			author,
			author_url,
			publish_date,
			content
		)
		VALUES
		(
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)
    `
	_, err = db.Exec(query, req.Title, req.Subtitle, previewImgName, postImgName, req.Author, authorImgName, req.PublishDate, req.Content)

	return err
}

func saveImg(imgName string, imgContent string) (string, error) {
	image, err := base64.StdEncoding.DecodeString(imgContent)
	if err != nil {
		log.Println(err)
		return "", err
	}

	imageFile, err := os.Create("static/img/" + imgName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = imageFile.Write(image)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "/static/img/" + imgName, err
}
