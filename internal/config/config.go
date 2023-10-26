package config

import (
	"errors"
	"flag"
	"os"

	"github.com/rs/zerolog"
)

const (
	defaultPort = ":8080"
)

type Config struct {
	DSN      string
	Port     string
	logLevel string
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
	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevel = zerolog.InfoLevel.String()
	}
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	return &Config{
		DSN:      dsn,
		Port:     port,
		logLevel: logLevel,
	}, nil
}
