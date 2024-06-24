package storages

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "library"
	dbUser     = "go_user"
	dbPassword = "P@ssw0rd"
)

var db *sql.DB

func InitDB() error {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = dbHost
	}
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if port == 0 {
		port = dbPort
	}
	dbname := os.Getenv("POSTGRES_DATABASE")
	if dbname == "" {
		dbname = dbName
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = dbUser
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = dbPassword
	}

	conString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", conString)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}

	return nil
}
