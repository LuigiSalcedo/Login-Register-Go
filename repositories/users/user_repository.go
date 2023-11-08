package users

import (
	"log"
	"webservice/database"
	"webservice/repositories/auth"
	"webservice/services"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type RegisterUser struct {
	User
	auth.AuthLogin
}

const (
	sqlInsertUser      = "INSERT INTO USER VALUES(?, ?, ?)"
	sqlInsertUserLogin = "INSERT INTO LOGIN VALUES(?, ?, ?)"
	sqlFetchUser       = "SELECT * FROM USER WHERE ID = ? LIMIT 1"
)

// Function to insert an user in the db
func InsertUser(u RegisterUser) error {
	tx, err := database.DB().Begin()

	defer func() {
		if tx == nil {
			log.Fatal("Error: Something went wrong with database connection.")
			return
		}

		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	if services.CheckError(err) {
		return err
	}

	_, err = tx.Exec(sqlInsertUser, u.Id, u.Name, u.Age)

	if services.CheckError(err) {
		return err
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	_, err = tx.Exec(sqlInsertUserLogin, u.Email, pwd, u.Id)

	if services.CheckError(err) {
		return err
	}

	return nil
}

// Search an user in the database
func SearchUser(id int64) (*User, error) {
	r, err := database.DB().Query(sqlFetchUser, id)

	if services.CheckError(err) {
		return nil, err
	}

	if !r.Next() {
		return nil, nil
	}

	user := &User{}

	err = r.Scan(&user.Id, &user.Name, &user.Age)

	if services.CheckError(err) {
		return nil, err
	}

	return user, nil
}
