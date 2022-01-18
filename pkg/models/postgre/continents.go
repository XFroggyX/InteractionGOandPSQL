package postgresql

import (
	"context"
	"database/sql"
	"errors"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type СontinentsModel struct {
	DB *pgxpool.Pool
}

func (m *СontinentsModel) NameField() []string {
	return []string{"ID", "Name"}
}

func (m *СontinentsModel) Get(ctx context.Context) ([]model.Сontinents, error) {
	stmp := `SELECT * FROM Сontinents`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.Сontinents
	for rows.Next() {
		var c model.Сontinents
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int32))
		c.Name = values[1].(string)

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}
