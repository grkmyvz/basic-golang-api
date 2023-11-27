package database

import (
	"fmt"
	"testing"
)

func TestDatabase(t *testing.T) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		dbname   = "randgo"
	)

	dbConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Connect(dbConfig)
}
