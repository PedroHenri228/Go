document.getElementById("teste").addEventListener('click', function() {
    var xhr = new XMLHttpRequest()
    xhr.open('GET', 'http://localhost:8080/data')
    xhr.onload = function() {
        if(xhr.status ===  200) {
            document.getElementById("result").innerText = xhr.responseText
        } else {
            document.getElementById("result").innerText = 'Error' + xhr.status
        }
    }
    xhr.onerror = function() {
        document.getElementById("result").innerHTML = "Erro na requisição"
        
    }
    xhr.send()
})

