package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("=====> API is Up and Running!")
	log.Println("=====> API running on 8080 port!")
	verifyEnv()
	database()
	api()
}

func verifyEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to get .env --> ", err)
	}

	log.Println("=====> .env Loaded!")
}

func api() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func database() {
	urlConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_IP"), "usavings")
	db, err := sql.Open("mysql", urlConnection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	log.Println("=====> Database connected!")
}
