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
	const maxUploadSize = 10 << 20
	err := r.ParseMultipartForm(maxUploadSize)
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}
