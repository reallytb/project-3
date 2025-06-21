package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER:", log.LstdFlags)
	server := server.NewServer(logger)
	err := server.Start()
	if err != nil {
		logger.Fatalf("ошибка при запуске сервера: %v", err)
	}
}
