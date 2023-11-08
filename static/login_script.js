let loginButton = document.getElementById("login")

loginButton.onclick = () => {
    let emailData = document.getElementById("email").value
    let pwdData = document.getElementById("password").value
    const data = {
        email: emailData,
        password: pwdData
    }

    fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if(response.status == 200) {
            response.json().then(json => {
                localStorage.setItem('jwt_login_register', json.jwt)
                window.location.href = 'http://localhost:8080/user/template'
            })
        }
    
        if(response.status == 401) {
            alert("Incorrect password")
        }

        if(response.status == 404) {
            alert("User not found")
        }
    })
    .catch(error => {
        alert(error)
    })
}