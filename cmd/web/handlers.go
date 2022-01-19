package main

import (
	"errors"
	"github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	postgresql "github.com/XFroggyX/InteractionGOandPSQL/pkg/models/postgre"
	"html/template"
	"log"
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

//nolint:funlen
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

//nolint:funlen
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

//nolint:funlen
func (app *application) deleteTable(writer http.ResponseWriter, request *http.Request) {
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

	switch {
	case tableName == "Countries":
		model := app.listTables["Countries"].(postgresql.CountriesModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Languages":
		model := app.listTables["Languages"].(postgresql.LanguagesModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "GovernmentForms":
		model := app.listTables["GovernmentForms"].(postgresql.GovernmentFormsModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "TerritorySizes":
		model := app.listTables["TerritorySizes"].(postgresql.TerritorySizesModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Religions":
		model := app.listTables["Religions"].(postgresql.ReligionsModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Associations":
		model := app.listTables["Associations"].(postgresql.AssociationsModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "AssociationsOfCountries":
		model := app.listTables["AssociationsOfCountries"].(postgresql.AssociationsOfCountriesModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Сontinents":
		model := app.listTables["Сontinents"].(postgresql.СontinentsModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "СontinentsOfCountries":
		model := app.listTables["СontinentsOfCountries"].(postgresql.СontinentsOfCountriesModel)
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		err = model.Delete(app.ctx, id)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	default:
		app.notFound(writer)
	}
}

//nolint:funlen
func (app *application) updateTable(writer http.ResponseWriter, request *http.Request) {
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

	log.Println(newItem)

	tableName := newItem["TableName"][0]

	switch {
	case tableName == "Countries":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		var operations string
		model := app.listTables["Countries"].(postgresql.CountriesModel)
		if newItem["CountriesName"][0] != "" {
			operations = "CountriesName"
		}
		if newItem["Flag"][0] != "" {
			operations = "Flag"
		}
		religionID, err := strconv.Atoi(newItem["ReligionID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		if religionID != 0 {
			operations = "ReligionID"
		}
		languagesID, err := strconv.Atoi(newItem["LanguagesID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		if languagesID != 0 {
			operations = "LanguagesID"
		}
		governmentFormID, err := strconv.Atoi(newItem["GovernmentFormID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		if governmentFormID != 0 {
			operations = "GovernmentFormID"
		}
		territorySizeID, err := strconv.Atoi(newItem["TerritorySizeID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		if territorySizeID != 0 {
			operations = "TerritorySizeID"
		}
		err = model.Update(app.ctx, id, operations, newItem[operations][0])
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Languages":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["Languages"].(postgresql.LanguagesModel)
		language := newItem["Language"][0]
		err = model.Update(app.ctx, id, "Language", language)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "GovernmentForms":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["GovernmentForms"].(postgresql.GovernmentFormsModel)
		form := newItem["Form"][0]
		err = model.Update(app.ctx, id, "Form", form)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "TerritorySizes":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["TerritorySizes"].(postgresql.TerritorySizesModel)
		type_ := newItem["Type"][0]
		err = model.Update(app.ctx, id, "Type", type_)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Religions":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["Religions"].(postgresql.ReligionsModel)
		title := newItem["Title"][0]
		err = model.Update(app.ctx, id, "Title", title)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Associations":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["Associations"].(postgresql.AssociationsModel)
		title := newItem["Title"][0]
		err = model.Update(app.ctx, id, "Title", title)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "AssociationsOfCountries":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["AssociationsOfCountries"].(postgresql.AssociationsOfCountriesModel)
		var operations string
		var value string
		countriesID, err := strconv.Atoi(newItem["CountriesID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		associationsID, err := strconv.Atoi(newItem["AssociationsID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		if countriesID == 0 {
			operations = "CountriesID"
			value = newItem["CountriesID"][0]
		}
		if associationsID == 0 {
			operations = "AssociationsID"
			value = newItem["AssociationsID"][0]
		}
		err = model.Update(app.ctx, id, operations, value)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "Continents":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["Сontinents"].(postgresql.СontinentsModel)
		name := newItem["Name"][0]
		err = model.Update(app.ctx, id, "Name", name)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	case tableName == "СontinentsOfCountries":
		id, err := strconv.Atoi(newItem["ID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		model := app.listTables["СontinentsOfCountries"].(postgresql.СontinentsOfCountriesModel)
		var operations string
		var value string
		countriesID, err := strconv.Atoi(newItem["CountriesID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		continentsID, err := strconv.Atoi(newItem["ContinentsID"][0])
		if err != nil {
			app.serverError(writer, err)
		}
		if countriesID == 0 {
			operations = "CountriesID"
			value = newItem["CountriesID"][0]
		}
		if continentsID == 0 {
			operations = "AssociationsID"
			value = newItem["ContinentsID"][0]
		}
		err = model.Update(app.ctx, id, operations, value)
		if err != nil {
			app.serverError(writer, err)
			return
		}
	default:
		app.notFound(writer)
	}
}
