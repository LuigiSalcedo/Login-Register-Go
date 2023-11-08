let loginButton = document.getElementById("login")
let registerButton = document.getElementById("register")

loginButton.onclick = () => {
    window.location.href = 'http://localhost:8080/login'
}

registerButton.onclick = () => {
    window.location.href = 'http://localhost:8080/register'
}