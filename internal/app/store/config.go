package store

type Config struct {
	DatabaseUrl string `toml:"database_url"`
}

func NEwConfig() *Config {
	return &Config{}
}
