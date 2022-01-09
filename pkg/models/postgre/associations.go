package postgresql

import (
	"context"
	"database/sql"
	"errors"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AssociationsModel struct {
	DB *pgxpool.Pool
}

func (m *AssociationsModel) NameField() []string {
	return []string{"ID", "Title"}
}

func (m *AssociationsModel) Get(ctx context.Context) ([]model.Associations, error) {
	stmp := `SELECT * FROM Associations`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.Associations
	for rows.Next() {
		var c model.Associations
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int32))
		c.Title = values[1].(string)

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}
