package main

import (
	"errors"
	"github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	postgresql "github.com/XFroggyX/InteractionGOandPSQL/pkg/models/postgre"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) createPage(data *templateData, files []string, writer http.ResponseWriter) {
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

func (app *application) home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		app.notFound(writer)
		return
	}

	files := []string{
		"./ui/html/index.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/content.page.html",
		"./ui/html/insert.page.html",
	}

	model := app.listTables["Countries"].(postgresql.CountriesModel)
	list, err := model.Get(app.ctx)
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
		NamesField:  model.NameField(),
		NameBD:      "Countries",
	}

	app.createPage(data, files, writer)
}

var tableName = ""

func (app *application) showTable(writer http.ResponseWriter, request *http.Request) {
	tableName = request.URL.Query().Get("name")
	if tableName == "" {
		app.notFound(writer)
		return
	}

	files := []string{
		"./ui/html/index.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/content.page.html",
		"./ui/html/insert.page.html",
	}

	data := &templateData{TabletNames: models.TabletNames}

	if tableName == "Countries" {
		model := app.listTables["Countries"].(postgresql.CountriesModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "Countries"
	} else if tableName == "Languages" {
		model := app.listTables["Languages"].(postgresql.LanguagesModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "Languages"
	} else if tableName == "GovernmentForms" {
		model := app.listTables["GovernmentForms"].(postgresql.GovernmentFormsModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "GovernmentForms"
	} else if tableName == "TerritorySizes" {
		model := app.listTables["TerritorySizes"].(postgresql.TerritorySizesModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "TerritorySizes"
	} else if tableName == "Religions" {
		model := app.listTables["Religions"].(postgresql.ReligionsModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "Religions"
	} else if tableName == "Associations" {
		model := app.listTables["Associations"].(postgresql.AssociationsModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "Associations"
	} else if tableName == "AssociationsOfCountries" {
		model := app.listTables["AssociationsOfCountries"].(postgresql.AssociationsOfCountriesModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "AssociationsOfCountries"
	} else if tableName == "Сontinents" {
		model := app.listTables["Сontinents"].(postgresql.СontinentsModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "Сontinents"
	} else if tableName == "СontinentsOfCountries" {
		model := app.listTables["СontinentsOfCountries"].(postgresql.СontinentsOfCountriesModel)

		table, err := model.Get(app.ctx)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(writer)
			} else {
				app.serverError(writer, err)
			}
			return
		}

		data.BD = table
		data.NamesField = model.NameField()
		data.NameBD = "СontinentsOfCountries"
	} else {
		app.notFound(writer)
	}

	app.createPage(data, files, writer)
}

func (app *application) insertTable(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}

	err := request.ParseForm()
	if err != nil {
		app.serverError(writer, err)
		return
	}

	newItem := request.Form

	tableName := newItem["TableName"][0]

	if tableName == "Countries" {
		model := app.listTables["Countries"].(postgresql.CountriesModel)
		countriesName := newItem["CountriesName"][0]
		flag := newItem["Flag"][0]

		religionID, err := strconv.Atoi(newItem["ReligionID"][0])
		if err != nil {
			app.serverError(writer, err)
		}

		languagesID, err := strconv.Atoi(newItem["LanguagesID"][0])
		if err != nil {
			app.serverError(writer, err)
		}

		governmentFormID, err := strconv.Atoi(newItem["GovernmentFormID"][0])
		territorySizeID, err := strconv.Atoi(newItem["TerritorySizeID"][0])

		err = model.Insert(app.ctx, countriesName, flag, religionID, languagesID, governmentFormID, territorySizeID)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "Languages" {
		model := app.listTables["Languages"].(postgresql.LanguagesModel)
		language := newItem["Language"][0]

		err = model.Insert(app.ctx, language)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "GovernmentForms" {
		model := app.listTables["GovernmentForms"].(postgresql.GovernmentFormsModel)
		form := newItem["Form"][0]

		err = model.Insert(app.ctx, form)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "TerritorySizes" {
		model := app.listTables["TerritorySizes"].(postgresql.TerritorySizesModel)
		type_ := newItem["Type"][0]

		err = model.Insert(app.ctx, type_)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "Religions" {
		model := app.listTables["Religions"].(postgresql.ReligionsModel)
		title := newItem["Title"][0]

		err = model.Insert(app.ctx, title)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "Associations" {
		model := app.listTables["Associations"].(postgresql.AssociationsModel)
		title := newItem["Title"][0]

		err = model.Insert(app.ctx, title)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "AssociationsOfCountries" {
		model := app.listTables["AssociationsOfCountries"].(postgresql.AssociationsOfCountriesModel)
		countriesID, err := strconv.Atoi(newItem["CountriesID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		associationsID, err := strconv.Atoi(newItem["AssociationsID"][0])
		if err != nil {
			app.serverError(writer, err)
		}

		err = model.Insert(app.ctx, countriesID, associationsID)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "Continents" {
		model := app.listTables["Сontinents"].(postgresql.СontinentsModel)
		name := newItem["Name"][0]

		err = model.Insert(app.ctx, name)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else if tableName == "СontinentsOfCountries" {
		model := app.listTables["СontinentsOfCountries"].(postgresql.СontinentsOfCountriesModel)
		countriesID, err := strconv.Atoi(newItem["CountriesID"][0])
		if err != nil {
			app.serverError(writer, err)
		}

		continentsID, err := strconv.Atoi(newItem["ContinentsID"][0])
		if err != nil {
			app.serverError(writer, err)
		}

		err = model.Insert(app.ctx, countriesID, continentsID)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	} else {
		app.notFound(writer)
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

	if tableName == "Countries" {
		convertUserID, err := strconv.Atoi(userID)
		if err != nil {
			app.serverError(writer, err)
			return
		}

		model := app.listTables["Countries"].(postgresql.CountriesModel)
		err = model.Update(app.ctx, convertUserID, nameFields, tableValue)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	}
}

func (app *application) deleteTable(writer http.ResponseWriter, request *http.Request) {

}
