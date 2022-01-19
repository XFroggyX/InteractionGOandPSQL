package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AssociationsOfCountriesModel struct {
	DB *pgxpool.Pool
}

func (m *AssociationsOfCountriesModel) NameField() []string {
	return []string{"CountriesID", "AssociationsID"}
}

func (m *AssociationsOfCountriesModel) Get(ctx context.Context) ([]model.AssociationsOfCountries, error) {
	stmp := `SELECT * FROM AssociationsOfCountries`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.AssociationsOfCountries
	for rows.Next() {
		var c model.AssociationsOfCountries
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.CountriesID = int(values[0].(int32))
		c.AssociationsID = int(values[1].(int32))

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (m *AssociationsOfCountriesModel) Insert(ctx context.Context, countriesID, associationsID int) error {
	stmp := `INSERT INTO AssociationsOfCountries (countriesID, associationsID) 
	VALUES ($1, $2)`

	_, err := m.DB.Exec(ctx, stmp, countriesID, associationsID)
	if err != nil {
		return err
	}

	return nil
}

func (m *AssociationsOfCountriesModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM AssociationsOfCountries WHERE countriesID = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *AssociationsOfCountriesModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE AssociationsOfCountries SET %s = `, nameFields)
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
