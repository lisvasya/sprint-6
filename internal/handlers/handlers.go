package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler отвечает за корневой путь и отдаёт HTML-файл
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("C:/Users/User/Desktop/lessons/sprint-6/sprint-6/index.html")

	if err != nil {
		http.Error(w, "Failed to read index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

// UploadHandler отвечает за загрузку файла и конвертацию
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const maxUploadSize = 10 << 20 // 10 MB
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "failed to parse multipart form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to get uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение данных из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Преобразование данных
	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "failed to convert data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Определение расширения и создание пути для сохранения файла
	ext := filepath.Ext(header.Filename)
	now := time.Now().UTC().Format("20060102_150405")
	localFileName := now + "_converted" + ext // Папка для хранения файлов

	// Сохраняем результат в файл
	f, err := os.Create(localFileName)
	if err != nil {
		http.Error(w, "Failed to create output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Записываем конвертированные данные в файл
	_, err = f.WriteString(converted)
	if err != nil {
		http.Error(w, "Failed to write to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ответ пользователю
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded and converted successfully"))
}
