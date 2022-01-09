package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CountriesModel struct {
	DB *pgxpool.Pool
}

func (m *CountriesModel) Insert(ctx context.Context, countriesName string, flag string, religionID int, languagesID int,
	governmentFormID int, territorySizeID int) error {
	stmp := `INSERT INTO Countries (CountriesName, Flag, ReligionID, LanguagesID, GovernmentFormID, TerritorySizeID) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.Exec(ctx, stmp, countriesName, flag, religionID, languagesID, governmentFormID,
		territorySizeID)
	if err != nil {
		return err
	}

	return nil
}

func (m *CountriesModel) Get(ctx context.Context) ([]model.Countries, error) {
	stmp := `SELECT * FROM countries`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.Countries
	for rows.Next() {
		var c model.Countries
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int32))
		c.CountriesName = values[1].(string)
		if values[2] != nil {
			c.Flag = values[2].(string)
		}
		if values[3] != nil {
			c.ReligionID = int(values[3].(int16))
		}
		c.LanguagesID = int(values[4].(int16))
		c.GovernmentFormID = int(values[5].(int16))
		c.TerritorySizeID = int(values[6].(int16))

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (m *CountriesModel) Update(ctx context.Context, id int, nameFields string,
	value string) error {
	stmp := fmt.Sprintf(`UPDATE Countries SET %s = `, nameFields)
	stmp = stmp + `$1 WHERE id = $2`
	_, err := m.DB.Exec(ctx, stmp, value, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}
