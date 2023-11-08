package users_services

import (
	"html/template"
	"net/http"
	"webservice/repositories/users"
	"webservice/security"
	"webservice/services"

	"github.com/golang-jwt/jwt/v5"
)

func LoadTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("./static/user_data.html")

	if services.CheckError(err) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error loading user_data.html"))
		return
	}
	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, nil)
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	recivedToken := r.Header.Get("Authorization")

	token, err := jwt.Parse(recivedToken, func(token *jwt.Token) (any, error) {
		return []byte(security.Secret()), nil
	})

	if services.CheckError(err) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing JWT data"))
	}

	data := token.Claims.(jwt.MapClaims)

	id, ok := data["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JWT not valid"))
	}

	user, err := users.SearchUser(int64(id.(float64)))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(services.ToJSON(user))
}
