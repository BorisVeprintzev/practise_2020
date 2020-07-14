package apiserver

import "apiserver/internal/store"

// Config - хранится конфигурация сервера
// BindAddr - порт
// LogLevel - logrus библеотека для логирования работы
// сам LogLevel - Уровень логирования
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig - возвращает конфиг сервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
