package apiserver

import "test/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` // toml tag
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "info",
		Store:    store.NEwConfig(),
	}
}
