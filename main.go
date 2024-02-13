package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/simple-golang-api/api"
	"github.com/simple-golang-api/database"
	"github.com/simple-golang-api/util"
)

func main() {
	log.Println("=====> API is Up and Running!")
	log.Println("=====> API running on 8080 port!")
	util.VerifyEnv()
	database.ConnDatabase()
	api.InitApi()
}
