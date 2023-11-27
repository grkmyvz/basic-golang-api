package main

import (
	"fmt"
	"randgo/api"
	"randgo/database"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "randgo"
)

func main() {
	dbConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db := database.Connect(dbConfig)
	database.Connection = db
	fmt.Println("Successfully connected Database!")
	api.APIServer()
}
