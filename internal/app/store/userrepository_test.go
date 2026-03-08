package store_test

//содержит простейшие тесты
import (
	"os"
	"testing"

	"test/internal/app/model"
	"test/internal/app/store"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseUrl := os.Getenv("DATABASE_URL") // read env vars
	if databaseUrl == "" {
		databaseUrl = "host=localhost dbname=restapi_test sslmode=disable"
	}

	testStore, tearDown := store.TestStore(t, databaseUrl)
	defer tearDown("users") //отложенный вызов фукнции. Выполнится последней после всех операций в функции

	user, err := testStore.User().CreateUser(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_findByEmail(t *testing.T) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost dbname=restapi_test sslmode=disable"
	}

	testStore, tearDown := store.TestStore(t, databaseUrl)
	defer tearDown("users") //отложенный вызов фукнции. Выполнится последней после всех операций в функции

	email := "user@example.org"
	_, err := testStore.User().FindByEmail(email)

	assert.Error(t, err)

	testStore.User().CreateUser(&model.User{
		Email: "user@example.org",
	})

	user, err := testStore.User().FindByEmail("ccccc")
	assert.NoError(t, err)
	assert.NotNil(t, user)

}
