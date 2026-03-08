package store

import (
	"fmt"
	"strings"
	"testing"
)

// тестовая функция. принимает строку подключения и возвращает тестовое хранилище
// funcв в аргументах - этоcallback функция дл очистки таблиц
func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper() // this is test-method
	config := NEwConfig()
	config.DatabaseUrl = databaseUrl
	store := New(config)
	err := store.Open()
	if err != nil {
		t.Fatal(err)
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			_, err2 := store.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE",
				strings.Join(tables, ",")))
			if err2 != nil {
				t.Fatal(err2)
			}
		}

		store.Close()
	}
}
