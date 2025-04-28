package handlers

import (
	"fmt"
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

	data, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "Failed to read index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

// UploadHandler отвечает за загрузку файла и конвертацию
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничиваем только метод POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Проверяем, что тип запроса multipart/form-data
	contentType := r.Header.Get("Content-Type")
	if contentType == "" || contentType[:19] != "multipart/form-data" {
		http.Error(w, "Content-Type isn't multipart/form-data", http.StatusBadRequest)
		return
	}

	// Парсим multipart форму
	err := r.ParseMultipartForm(10 << 20) // 10 МБ
	if err != nil {
		http.Error(w, "failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to get uploaded file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Вызываем функцию автоопределения типа и конвертации
	convertedStr, err := service.Convert(string(fileBytes))
	if err != nil {
		http.Error(w, "failed to convert file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Генерируем имя для нового файла
	extension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%s_converted%s", time.Now().UTC().Format("20060102_150405"), extension)

	// Создаем новый файл
	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "failed to create output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Записываем результат конвертации в новый файл
	_, err = newFile.WriteString(convertedStr)
	if err != nil {
		http.Error(w, "failed to write to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем результат конвертации пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(convertedStr))
}
