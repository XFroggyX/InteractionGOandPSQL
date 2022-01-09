package postgresql

import (
	"context"
	"database/sql"
	"errors"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type СontinentsOfCountriesModel struct {
	DB *pgxpool.Pool
}

func (m *СontinentsOfCountriesModel) NameField() []string {
	return []string{"CountriesID", "ContinentsID"}
}

func (m *СontinentsOfCountriesModel) Get(ctx context.Context) ([]model.СontinentsOfCountries, error) {
	stmp := `SELECT * FROM СontinentsOfCountries`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.СontinentsOfCountries
	for rows.Next() {
		var c model.СontinentsOfCountries
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.CountriesID = int(values[0].(int32))
		c.ContinentsID = int(values[1].(int32))

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}
