package config

import (
	"errors"
	"os"
)

const (
	defaultPort = ":8080"
)

type Config struct {
	DSN  string
	Port string
}

func New() (*Config, error) {
	dsn, ok := os.LookupEnv("DB_DSN")
	if !ok || dsn == "" {
		return nil, errors.New("database source variable not found or empty")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}
	return &Config{
		DSN:  dsn,
		Port: port,
	}, nil
}
