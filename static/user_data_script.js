const updateData = (id, name, age) => {
    ID = document.getElementById("id")
    NAME = document.getElementById("name")
    AGE = document.getElementById("age")

    ID.innerText = id
    NAME.innerText = name
    AGE.innerText = age
}

const fetchData = () => {
    fetch('http://localhost:8080/user', {
        method: 'GET',
        headers: {
            'Authorization': localStorage.getItem('jwt_login_register')
        }
    })
    .then(response => {
        if(response.status == 200) {
            response.json().then(r => {
                updateData(r.id, r.name, r.age)
            })
        }
    })
    .catch(err => {
        alert(err)
    })
}

const homeButton = document.getElementById("home")

homeButton.onclick = () => {
    window.location.href = 'http://localhost:8080/'
}

fetchData()