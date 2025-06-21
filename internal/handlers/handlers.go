package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	dir := filepath.Dir("index.html")
	projectRoot := filepath.Join(dir, "..")
	path, err := filepath.Abs(projectRoot)
	newPath := filepath.Join(path, "index.html")
	if err != nil {
		log.Println("ошибка определения пути html-документа")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.Open(newPath)
	if err != nil {
		log.Println("ошибка открытия html-документа")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println("ошибка копирования html-документа")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("ошибка парсинга")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Форма запарсилась")
	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Println("ошибка получения файла из формы")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("получен файл из формы")
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("ошибка чтения файла")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("файл прочитан")
	convertedData := service.Convert(string(fileBytes))
	log.Println("данные конвертированы")
	outputFilename := fmt.Sprintf("result_%s%s",
		time.Now().UTC().Format("20060102_150405"),
		filepath.Ext(header.Filename))
	newFile, err := os.Create(outputFilename)
	if err != nil {
		log.Println("ошибка создания файла")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("файл создан")
	defer newFile.Close()
	err = os.WriteFile(outputFilename, []byte(convertedData), 0755)
	if err != nil {
		log.Println("ошибка записи файла")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("файл записан")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedData))
	log.Println("---------------------------------")
}
