package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

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
	Image       string `db:"image_url"`
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
	Image    string `db:"image_url"`
	Content  string `db:"content"`
}

type adminPage struct {
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
			image_url,
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
        	image_url,
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
	ts, err := template.ParseFiles("admin/admin.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := adminPage{
		Title:    "Escape",
		Subtitle: "Biba i Boba",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}
