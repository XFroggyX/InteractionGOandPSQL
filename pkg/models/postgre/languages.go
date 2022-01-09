package postgresql

import (
	"context"
	"database/sql"
	"errors"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type LanguagesModel struct {
	DB *pgxpool.Pool
}

func (m *LanguagesModel) NameField() []string {
	return []string{"ID", "Language"}
}

func (m *LanguagesModel) Get(ctx context.Context) ([]model.Languages, error) {
	stmp := `SELECT * FROM languages`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.Languages
	for rows.Next() {
		var c model.Languages
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int16))
		c.Language = values[1].(string)

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}
