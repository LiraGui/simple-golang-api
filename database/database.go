package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnDatabase() {
	urlConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_IP"), "usavings")
	db, err := sql.Open("mysql", urlConnection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	log.Println("=====> Database connected!")
}
