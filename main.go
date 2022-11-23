package main

import (
	"gin-mysql-jwt/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDb()

	r := gin.Default()
	Routes(r)
	r.Run(":8080")
}
