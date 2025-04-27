package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Server *http.Server
	Logger *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	// Создаем новый мультиплексор (роутер)
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.IndexHandler)        // Страница индекса
	mux.HandleFunc("/upload", handlers.UploadHandler) // Обработка загрузки файла

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		Server: s,
	}

}
func (s *Server) Run() error {
	s.Logger.Println("Server started on :8080")
	return s.Server.ListenAndServe()
}
