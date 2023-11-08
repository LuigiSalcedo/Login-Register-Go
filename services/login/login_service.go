package login

import (
	"encoding/json"
	"html/template"
	"net/http"
	"webservice/repositories/auth"
	"webservice/security"
	"webservice/services"

	"github.com/golang-jwt/jwt/v5"
)

type LoginService struct{}

func loadFile(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html")

	t, err := template.ParseFiles("././static/login.html")

	if services.CheckError(err) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error loading file: login.html"))
		return
	}
	t.Execute(w, nil)
}

func authLogin(w http.ResponseWriter, r *http.Request) {
	authData := auth.AuthLogin{}

	err := json.NewDecoder(r.Body).Decode(&authData)

	if services.CheckError(err) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := auth.VerifyLoginData(authData)

	if err != nil {
		switch err {
		case auth.ErrInternal500:
			w.WriteHeader(http.StatusInternalServerError)
		case auth.ErrNotFound404:
			w.WriteHeader(http.StatusNotFound)
		case auth.ErrNotValid:
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.Write([]byte(err.Error()))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	signedToken, err := token.SignedString([]byte(security.Secret()))

	if services.CheckError(err) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(services.ToJSON(struct {
		JWT string `json:"jwt"`
	}{signedToken}))
}

func (ls *LoginService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loadFile(w)
	case http.MethodPost:
		authLogin(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}
