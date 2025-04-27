package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "server: ", log.LstdFlags)

	srw := server.NewServer(logger)

	logger.Println("starting server on :8080")
	err := srw.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
