package postgresql

import (
	"context"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CountriesModel struct {
	DB *pgxpool.Pool
}

func (m *CountriesModel) Insert(ctx context.Context, countriesName string, Flag string, ReligionID int, LanguagesID int,
	governmentFormID int, TerritorySizeID int) (int, error) {
	stmp := `INSERT INTO Countries (CountriesName, Flag, ReligionID, LanguagesID, GovernmentFormID, TerritorySizeID) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.Exec(ctx, stmp, countriesName, Flag, ReligionID, LanguagesID, governmentFormID,
		TerritorySizeID)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (m *CountriesModel) Get(ctx context.Context) ([]model.Countries, error) {
	stmp := `SELECT * FROM countries`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []model.Countries
	for rows.Next() {
		var c model.Countries
		values, err := rows.Values()
		if err != nil {
			return nil, nil // костыль
		}

		c.ID = int(values[0].(int16))
		c.CountriesName = values[1].(string)
		if values[2] != nil {
			c.Flag = values[2].(string)
		} else {
			c.Flag = ""
		}

		if values[3] != nil {
			c.ReligionID = int(values[3].(int16))
		} else {
			c.ReligionID = -1
		}
		c.LanguagesID = int(values[4].(int16))
		c.GovernmentFormID = int(values[5].(int16))
		c.TerritorySizeID = int(values[6].(int16))

		courses = append(courses, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
