package database

import (
	"arrieup/collocom/serverside/user"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var cfg mysql.Config

func DBsetup() {
	cfg = mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "collocom",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func CreateUser(usr user.User) (int64, error) {
	result, err := db.Exec("INSERT INTO user (id, username, email, password, create_time) VALUES (?, ?, ?, ?, ?)", usr.Id, usr.Username, usr.Email, usr.Password, usr.Create_time)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}
	return id, nil
}

func ReadUserByID(id int) (user.User, error) {
	var usr user.User

	row := db.QueryRow("SELECT * FROM user WHERE ID = ?", id)
	var T string
	if err := row.Scan(&usr.Id, &usr.Username, &usr.Email, &usr.Password, &T); err != nil {
		if err == sql.ErrNoRows {
			usr.Create_time, _ = time.Parse("2006-01-02 15:04:05", T)
			return usr, fmt.Errorf("ReadUserByID %d: no such user", id)
		}
		usr.Create_time, _ = time.Parse("2006-01-02 15:04:05", T)
		return usr, fmt.Errorf("ReadUserByID %d: %v", id, err)
	}
	usr.Create_time, _ = time.Parse("2006-01-02 15:04:05", T)
	return usr, nil
}

func ReadUserByUsername(username string) (user.User, error) {
	var usr user.User

	row := db.QueryRow("SELECT * FROM user WHERE USERNAME = ?", username)
	var T string
	if err := row.Scan(&usr.Id, &usr.Username, &usr.Email, &usr.Password, &T); err != nil {
		if err == sql.ErrNoRows {
			usr.Create_time, _ = time.Parse("2006-01-02 15:04:05", T)
			return usr, fmt.Errorf("ReadUserByUsername %s: no such user", username)
		}
		usr.Create_time, _ = time.Parse("2006-01-02 15:04:05", T)
		return usr, fmt.Errorf("ReadUserByUsername %s: %v", username, err)
	}
	usr.Create_time, _ = time.Parse("2006-01-02 15:04:05", T)
	return usr, nil
}
