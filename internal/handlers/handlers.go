package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "index.html is not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		http.Error(w, "error while opening index.html", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, "index.html", info.ModTime(), file)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error with file", http.StatusBadRequest)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error while reading file", http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
		http.Error(w, "file is empty", http.StatusInternalServerError)
		return
	}

	converted := service.Convert(string(data))
	if err != nil {
		http.Error(w, "error while writing file", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}
	name := strings.ReplaceAll(time.Now().UTC().String(), ":", "-") + ext

	out, err := os.Create(name)
	if err != nil {
		http.Error(w, "cannot create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := out.Write([]byte(converted)); err != nil {
		http.Error(w, "error while writing file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(converted))
}
