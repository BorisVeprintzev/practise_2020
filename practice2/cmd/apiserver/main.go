package main

import (
	"apiserver/internal/app/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string //указание откуда брать конфиг
)

// ключи
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
