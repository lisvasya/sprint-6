package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)

	// Создаем сервер с помощью функции из пакета server
	s := server.NewServer(logger)
	logger.Println("start the server")

	// Запускаем сервер и проверяем ошибки
	if err := s.ListenAndServe(); err != nil {
		// Если возникла ошибка при запуске сервера, выводим её с Fatal
		logger.Fatal("Server failed to start: ", err)
	}
}
