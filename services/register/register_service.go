package register

import (
	"encoding/json"
	"html/template"
	"net/http"
	"webservice/repositories/users"
	"webservice/services"
)

type RegisterService struct{}

func loadFile(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html")

	t, err := template.ParseFiles("././static/register.html")

	if services.CheckError(err) {
		w.WriteHeader(500)
		w.Write([]byte("Error loading file: register.html"))
		return
	}

	t.Execute(w, nil)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	u := users.RegisterUser{}

	err := json.NewDecoder(r.Body).Decode(&u)

	if u.Age < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Age not valid."))
		return
	}

	if u.Id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id not valid."))
		return
	}

	if len(u.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name not valid"))
	}

	if len(u.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password not valid"))
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Formato de entrada no valido"))
		return
	}

	err = users.InsertUser(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error with user register"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (rs *RegisterService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loadFile(w)
	case http.MethodPost:
		registerUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
	}
}
