package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost dbname=restapi sslmode=disable",
	}
}

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		db: nil,
	}
}

func (s *Store) Open(databaseURL string) error {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// функция -getter для репозитория. Реализация инкапсуляции
// либо возвращает готовое, либо создает репозиторий -singleton
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
