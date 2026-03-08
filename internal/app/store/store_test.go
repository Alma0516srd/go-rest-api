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
		databaseUrl = "host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
