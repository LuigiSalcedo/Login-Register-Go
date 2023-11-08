package database

import (
	"database/sql"
	"fmt"
	"os"
	"webservice/services"

	_ "github.com/go-sql-driver/mysql"
)

const initUser = "CREATE TABLE IF NOT EXISTS USER(ID INTEGER PRIMARY KEY, NAME TEXT NOT NULL, AGE INTEGER NOT NULL)"

const initLogin = `
CREATE TABLE IF NOT EXISTS LOGIN(EMAIL VARCHAR(40), PASSWORD TEXT NOT NULL, ID_USER INTEGER NOT NULL, 
FOREIGN KEY(ID_USER) REFERENCES USER(ID))`

var conn *sql.DB

func init() {
	user := os.Args[1]
	pass := os.Args[2]
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/loginregister", user, pass))

	if services.CheckError(err) {
		panic(err)
	}

	_, err = db.Exec(initUser)

	if services.CheckError(err) {
		panic(err)
	}

	_, err = db.Exec(initLogin)

	if services.CheckError(err) {
		panic(err)
	}

	conn = db
}

// Get the database connection
func DB() *sql.DB {
	return conn
}
