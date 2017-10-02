var newOutput = document.getElementById(".inputName").addEventListener("submit", ()=>{
    url = "localhost:4000/hello?" + inputName;
    console.log(url)
    var promise = fetch(url);
    promise.then(function(response) {
        return response.json();
    })
})

output.textContent = "newOutput"
