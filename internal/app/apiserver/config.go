package apiserver

import (
	"apiserver/internal/app/store"
)

type Config struct {
	BinAddr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8082",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
