package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, "failed to parse multipart form: "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "failed to get uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to get uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "failed to convert data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	now := time.Now().UTC().Format("20060102_150405")

	localFileName := now + "_converted" + ext

	f, err := os.Create(localFileName)
	if err != nil {
		http.Error(w, "failed to write to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = f.WriteString(converted)
	if err != nil {
		http.Error(w, "failed to write to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = w.Write([]byte(converted))
	if err != nil {
		http.Error(w, "server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
