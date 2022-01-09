package main

import (
	"errors"
	"fmt"
	"github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		app.notFound(writer)
		return
	}

	list, err := app.countries.Get(app.ctx)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(writer)
		} else {
			app.serverError(writer, err)
		}
		return
	}

	data := &templateData{
		BD:          list,
		TabletNames: models.TabletNames,
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
	err = ts.Execute(writer, data)
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

	if tableName == "countries" {
		list, err := app.countries.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data := &templateData{
			BD:          list,
			TabletNames: models.TabletNames,
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
		err = ts.Execute(writer, data)
		if err != nil {
			app.serverError(writer, err)
		}
	} else {
		_, err := fmt.Fprintf(writer, "Отображение выбранной таблицы с NAME %s...", tableName)
		if err != nil {
			app.errorLog.Println(err)
		}
	}
}

func (app *application) insertTable(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}

	var (
		CountriesName    = "Польша"
		Flag             = ""
		ReligionID       = 2
		LanguagesID      = 1
		GovernmentFormID = 1
		TerritorySizeID  = 1
	)
	err := app.countries.Insert(app.ctx, CountriesName, Flag, ReligionID, LanguagesID, GovernmentFormID, TerritorySizeID)
	if err != nil {
		app.serverError(writer, err)
		return
	}
}

func (app *application) updateTable(writer http.ResponseWriter, request *http.Request) {
	tableName := request.URL.Query().Get("name")
	userID := request.URL.Query().Get("id")
	tableValue := request.URL.Query().Get("value")
	nameFields := request.URL.Query().Get("fields")
	if tableName == "" || userID == "" || tableValue == "" || nameFields == "" {
		app.notFound(writer)
		return
	}

	if tableName == "countries" {
		convertUserID, err := strconv.Atoi(userID)
		if err != nil {
			app.serverError(writer, err)
			return
		}

		err = app.countries.Update(app.ctx, convertUserID, nameFields, tableValue)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	}
}
