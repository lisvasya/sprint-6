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
	srv := server.NewServer(logger)

	// Запускаем сервер и проверяем ошибки
	if err := srv.Run(); err != nil {
		// Если возникла ошибка при запуске сервера, выводим её с Fatal
		logger.Fatal("Server failed to start: ", err)
	}
}
