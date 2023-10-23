package main

import (
	"fmt"

	"github.com/Kartochnik010/test-effectivemobile/internal/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}
