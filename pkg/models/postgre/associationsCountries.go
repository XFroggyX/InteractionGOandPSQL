package postgresql

import (
	"context"
	"database/sql"
	"errors"
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
