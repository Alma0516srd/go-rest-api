package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl := os.Getenv("DATABASE_URL") // read from env-vars
	if databaseUrl == "" {
		databaseUrl = "host=localhost port=5435 dbname=restapi_dev user=postgres password=postgres sslmode=disable"
	}

	os.Exit(m.Run())
}
