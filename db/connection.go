package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func ConnectDb() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_Password")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@/%s", dbUsername, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to database success!!")

	Conn = db
}
