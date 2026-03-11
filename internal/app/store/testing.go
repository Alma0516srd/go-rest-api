package store

import (
	"database/sql"
	"strings"
	"testing"
)

// тестовая функция. принимает строку подключения и возвращает тестовое хранилище
// funcв в аргументах - этоcallback функция дл очистки таблиц
func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper() // this is test-method
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE ", strings.Join(tables, " "))
		}
		db.Close()
	}
}
