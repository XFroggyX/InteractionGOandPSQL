package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	files := []string{
		"./ui/html/index.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(writer, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Server, Error", 500)
	}
}

func showTable(writer http.ResponseWriter, request *http.Request) {
	tableName := request.URL.Query().Get("name")
	if tableName == "" {
		http.NotFound(writer, request)
		return
	}

	// writer.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(writer, "Отображение выбранной таблицы с NAME %s...", tableName)
	if err != nil {
		log.Fatal(err)
	}
}

func insertTable(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Метод запрещен", 405)
		return
	}
	_, err := writer.Write([]byte("Insert page"))
	if err != nil {
		log.Fatal(err)
	}
}
