package auth

import (
	"errors"
	"webservice/database"
	"webservice/services"

	"golang.org/x/crypto/bcrypt"
)

type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	searchLoginData = `SELECT ID_USER, PASSWORD FROM LOGIN WHERE EMAIL = ? LIMIT 1`
)

var (
	ErrInternal500 = errors.New("internal server error")
	ErrNotFound404 = errors.New("user not found")
	ErrNotValid    = errors.New("incorrect password")
)

func VerifyLoginData(data AuthLogin) (int64, error) {
	r, err := database.DB().Query(searchLoginData, data.Email)

	if services.CheckError(err) {
		return -1, ErrInternal500
	}

	if !r.Next() {
		return -1, ErrNotFound404
	}

	var id int64
	var hash string

	err = r.Scan(&id, &hash)

	if services.CheckError(err) {
		return -1, ErrInternal500
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))

	if services.CheckError(err) {
		return -1, ErrNotValid
	}

	return id, nil
}
