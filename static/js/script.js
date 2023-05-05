const button = document.getElementById('button');
button.onclick = function(event) {
  let xhr = new XMLHttpRequest();

  xhr.open('GET', 'http://localhost:3000/admin');
  
  xhr.responseType = 'json';
  
  xhr.send();
  
  // тело ответа {"сообщение": "Привет, мир!"}
  xhr.onload = function() {
    let responseObj = xhr.response;
    alert(responseObj.message); // Привет, мир!
  };
};
