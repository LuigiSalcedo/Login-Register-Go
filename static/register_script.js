let registerButton = document.getElementById("register")

registerButton.onclick = () => {
    let idData = parseInt(document.getElementById("id").value) 
    let nameData = document.getElementById("name").value
    let ageData = parseInt(document.getElementById("age").value)
    let emailData = document.getElementById("email").value
    let pwdData = document.getElementById("pwd").value

    const data = {
        id: idData,
        name: nameData,
        age: ageData,
        email: emailData,
        password: pwdData
    }

    fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        alert("User registered successfully")
        window.location.href = 'http://localhost:8080/'
    })
    .catch(error => {
        alert(error)
    })
}