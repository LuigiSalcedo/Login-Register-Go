package main

import (
	"html/template"
	"log"
	"net/http"
	"webservice/services/login"
	"webservice/services/register"
	users_services "webservice/services/users_service"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/login", &login.LoginService{})
	http.Handle("/register", &register.RegisterService{})
	http.HandleFunc("/user", users_services.FetchUser)
	http.HandleFunc("/user/template", users_services.LoadTemplate)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		t, err := template.ParseFiles("index.html")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error loading the index.html file"))
			log.Println(err)
			return
		}

		err = t.Execute(w, nil)

		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Server running . . . ")
	http.ListenAndServe(":8080", nil)
}
