package models

import "errors"

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Countries struct {
	ID               int
	CountriesName    string
	Flag             string
	ReligionID       int
	LanguagesID      int
	GovernmentFormID int
	TerritorySizeID  int
}
