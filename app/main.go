package main

import (
	"deployRunner/app/telegram"
	"deployRunner/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	telegram.NewListener(cfg).Listen()
}
