package store

// Config - конфигурация ДБ
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig - создание конф
func NewConfig() *Config {
	return &Config{}
}
