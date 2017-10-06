"use strict";
document.querySelector("#form").addEventListener("submit", (evt)=>{
    evt.preventDefault();
    var inputName = document.querySelector("#input-name").value;
    var url = "http://localhost:4000/hello?name=" + inputName;
    fetch(url)
        .then((response) => {
            if(response.ok) {
                return response.text();
            }
            return response.text().then((t) => Promise.reject(t));
        })
        .then((data) => {
            console.log(data);
        })
        .catch((err) => {
            console.error(err);
        });
})

document.querySelector("#memory").addEventListener("click", ()=>{
    var url = "localhost:4000/memory"
    setInterval(() => {
        fetch(url)
        .then((response) => {
            return response.json();
        });
    }, 1000);
})

document.querySelector("#city-form").addEventListener("submit", (evt) => {
    evt.preventDefault();
    var inputCity = document.querySelector("#input-city").value;
    var url = "http://localhost:4000/zips/" + inputCity;
    fetch(url) 
        .then((response) => {
            if(response.ok) {
                return response.text();
            }
            return reponse.text().then((t) => Promise.reject(t));
        })
        .then((data) => {
            var output = document.querySelector('#zip-output');
            output.textContent = data;
        })
        .catch((err) => {
            console.error(err);
        })
})

output.textContent = "newOutput"
