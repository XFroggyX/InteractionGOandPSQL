package models

import "errors"

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

var TabletNames = []string{
	"Countries",
	"Languages",
	"GovernmentForms",
	"TerritorySizes",
	"Religions",
	"Associations",
	"AssociationsOfCountries",
	"Сontinents",
	"СontinentsOfCountries",
}

type Countries struct {
	ID               int
	CountriesName    string
	Flag             string
	ReligionID       int
	LanguagesID      int
	GovernmentFormID int
	TerritorySizeID  int
}

type Languages struct {
	ID       int
	Language string
}

type GovernmentForms struct {
	ID   int
	Form string
}

type TerritorySizes struct {
	ID   int
	Type string
}

type Religions struct {
	ID    int
	Title string
}

type Associations struct {
	ID    int
	Title string
}

type AssociationsOfCountries struct {
	CountriesID    int
	AssociationsID int
}

type Сontinents struct {
	ID   int
	Name string
}

type СontinentsOfCountries struct {
	CountriesID  int
	ContinentsID int
}
