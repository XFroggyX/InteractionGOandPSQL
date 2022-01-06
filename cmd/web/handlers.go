package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		app.notFound(writer)
		return
	}

	files := []string{
		"./ui/html/index.page.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(writer, err)
		return
	}

	err = ts.Execute(writer, nil)
	if err != nil {
		app.serverError(writer, err)
	}
}

func (app *application) showTable(writer http.ResponseWriter, request *http.Request) {
	tableName := request.URL.Query().Get("name")
	if tableName == "" {
		app.notFound(writer)
		return
	}

	// writer.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(writer, "Отображение выбранной таблицы с NAME %s...", tableName)
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) insertTable(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}
	_, err := writer.Write([]byte("Insert page"))
	if err != nil {
		app.errorLog.Println(err)
	}
}
