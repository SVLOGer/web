const uploadTitle = document.getElementById('title');
const uploadSubtitle = document.getElementById('subtitle');
const uploadAuthorName = document.getElementById('authorName');
const uploadPublishDate = document.getElementById('date');
const uploadAuthorPhoto = document.getElementById("photoAuthor");
const uploadTinyPostIMG = document.getElementById("photoPostTiny");
const uploadPostIMG = document.getElementById("photoPost");
const uploadcontent = document.getElementById("contentPost");

const removeAuthorPhoto = document.getElementById("removePhotoAuthor");
const removePostIMG = document.getElementById("removePostIMG");
const removeTinyPostIMG = document.getElementById("removeTinyPostIMG");

let AuthorPhoto;
let PostIMG;
let PostTinyIMG;

let authorIMGInputName;
let previewIMGInputName;
let postImgInputName;

let uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");

const publishButton = document.querySelector(".publish-button");

uploadTitle.addEventListener (
  "input" , 
  () => {
    let title = document.getElementById('title').value;
    let defaultTitle = 'New Post';
    if (title !== '') {
      document.getElementById('titleVisual').innerHTML = title;
      document.getElementById('titleVisualTiny').innerHTML = title;
    }
    else {
      document.getElementById('titleVisual').innerHTML = defaultTitle;
      document.getElementById('titleVisualTiny').innerHTML = defaultTitle;
    }
  }  
)

uploadSubtitle.addEventListener (
  "input" , 
  () => {
    let subtitle = document.getElementById('subtitle').value;
    let defaultSubtitle = 'Please, enter any description';
    if (subtitle !== ''){
      document.getElementById('subtitleVisual').innerHTML = subtitle;
      document.getElementById('subtitleVisualTiny').innerHTML = subtitle;
    } 
    else {
      document.getElementById('subtitleVisual').innerHTML = defaultSubtitle;
      document.getElementById('subtitleVisualTiny').innerHTML = defaultSubtitle;
    }
  }     
)

uploadAuthorName.addEventListener (
  "input" , 
  () => {
    let name = document.getElementById('authorName').value;
    let defaultName = 'Enter author name';
    if (name !== ''){
      document.getElementById('authorNameVisualTiny').innerHTML = name;
    }    
    else {
      document.getElementById('authorNameVisualTiny').innerHTML = defaultName;
    }
  }  
)

uploadPublishDate.addEventListener (
  "input" , 
  () => {
    let date = document.getElementById('date').value;
    let defaultDate = '01/05/2023';
    if (date !== ''){
      document.getElementById('dateVisuality').innerHTML = date;
    }    
    else {
      document.getElementById('dateVisuality').innerHTML = defaultDate;
    }
  }  
)

uploadAuthorPhoto.addEventListener (
  "input" , 
  () => {
    const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
    const previewInput = document.querySelector(".preview-author-photo");
    const file = document.querySelector(".author-photo-form__input").files[0];
    const reader = new FileReader();
      reader.addEventListener(
        "load",
        () => {
          previewPostCardAuthorPhoto.src = reader.result;
          previewInput.src = reader.result;
          AuthorPhoto = reader.result.replace("data:", "").replace(/^.+,/, "");
          authorIMGInputName = file.name;
        },
        false
      );
    document.querySelector(".author-photo-form__remove-button").classList.add("remove-button__remove-button-show");
    uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
    uploadAuthorPhotoButton.innerHTML = 'Upload New';
    uploadAuthorPhotoButton.classList.add("author-photo-form__upload-button-view");
    document.querySelector(".upload-button__icon").classList.add("upload-button__icon-view");
    if (file) {
      reader.readAsDataURL(file);
   }
  }
)

removeAuthorPhoto.addEventListener (
  "click" , 
  () => {
    document.querySelector(".author-photo-form__remove-button").classList.remove("remove-button__remove-button-show");
    uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
    uploadAuthorPhotoButton.innerHTML = 'Upload';
    uploadAuthorPhotoButton.classList.remove("author-photo-form__upload-button-view");
    document.querySelector(".upload-button__icon").classList.remove("upload-button__icon-view");
    const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
    const previewInput = document.querySelector(".preview-author-photo");
    defaultAuthorPhoto = '/static/img/author0.png';
    previewPostCardAuthorPhoto.src = defaultAuthorPhoto;
    previewInput.src = defaultAuthorPhoto;
  }
)
uploadTinyPostIMG.addEventListener (
  "input" , 
  () => {
    const previewPostCardPhoto = document.querySelector(".post-card__photo");
    const previewInput = document.querySelector(".upload-place-tiny__img");
    const file = document.querySelector(".input-hero-image-tiny__input").files[0];
    const reader = new FileReader();
    reader.addEventListener(
      "load",
      () => {
        previewPostCardPhoto.src = reader.result;
        previewInput.src = reader.result;
        document.querySelector(".tiny-img-buttons").classList.add("tiny-img-buttons-show");
        document.querySelector(".input-hero-image-tiny__sign").classList.add("input-hero-image-tiny__sign-remove");
        PostTinyIMG = reader.result.replace("data:", "").replace(/^.+,/, "");
        previewIMGInputName = file.name;
      },
      false
    );

    if (file) {
      reader.readAsDataURL(file);
    }
  }
)

removeTinyPostIMG.addEventListener (
  "click" , 
  () => {
    const defaultPostTinyIMG = "/static/img/postphototiny0.png";
    const previewPostCardPhoto = document.querySelector(".post-card__photo");
    const previewInput = document.querySelector(".upload-place-tiny__img");
    previewPostCardPhoto.src = defaultPostTinyIMG;
    previewInput.src = defaultPostTinyIMG;
    document.querySelector(".tiny-img-buttons").classList.remove("tiny-img-buttons-show");
    document.querySelector(".input-hero-image-tiny__sign").classList.remove("input-hero-image-tiny__sign-remove");
  }
)

uploadPostIMG.addEventListener (
  "input" , 
  () => {
    const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
    const previewInput = document.querySelector(".upload-place__img");
    const file = document.querySelector(".input-hero-image__input").files[0];
    const reader = new FileReader();
    reader.addEventListener(
      "load",
      () => {
        previewPostCardPhoto.src = reader.result;
        previewInput.src = reader.result;
        document.querySelector(".img-buttons").classList.add("img-buttons-show");
        document.querySelector(".input-hero-image__sign").classList.add("input-hero-image__sign-remove");
        PostIMG = reader.result.replace("data:", "").replace(/^.+,/, "");
        postImgInputName = file.name;
      },
      false
    );

    if (file) {
      reader.readAsDataURL(file);
    }
  }
)

removePostIMG.addEventListener (
  "click" , 
  () => {
    const defaultPostIMG = "/static/img/postphoto0.png";
    const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
    const previewInput = document.querySelector(".upload-place__img");
    document.querySelector(".img-buttons").classList.remove("img-buttons-show");
    document.querySelector(".input-hero-image__sign").classList.remove("input-hero-image__sign-remove");
    previewPostCardPhoto.src = defaultPostIMG;
    previewInput.src = defaultPostIMG;
  }
)  

// вывод json  
publishButton.addEventListener(
  "click",
  () => {
    const data = {
      title: uploadTitle.value,
      subtitle: uploadSubtitle.value,
      authorName: uploadAuthorName.value, 
      publishDate: uploadPublishDate.value,
      content: uploadcontent.value ,
      postIMG: PostIMG, 
      authorIMG: AuthorPhoto,
      previewIMG: PostTinyIMG,
      postIMGName: postImgInputName,
      authorIMGName: authorIMGInputName,
      previewIMGName: previewIMGInputName,
    }
    console.log(JSON.stringify(data, null, "\t"));
    doFecth(data);
  }    
)

async function doFecth(data) {
  const response = await fetch("/api/post", {
    method: "POST",
    headers: {
      "Content-Type": "application/json;charset=utf-8"
    },
    body: JSON.stringify(data)
  });
  if (!response.ok) {
    alert("Ошибка HTTP: " + response.status);
  }
}