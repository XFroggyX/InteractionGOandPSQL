package postgresql

import (
	"context"
	"database/sql"
	"errors"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ReligionsModel struct {
	DB *pgxpool.Pool
}

func (m *ReligionsModel) NameField() []string {
	return []string{"ID", "Title"}
}

func (m *ReligionsModel) Get(ctx context.Context) ([]model.Religions, error) {
	stmp := `SELECT * FROM Religions`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.Religions
	for rows.Next() {
		var c model.Religions
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int16))
		c.Title = values[1].(string)

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (m *ReligionsModel) Insert(ctx context.Context, title string) error {
	stmp := `INSERT INTO Religions (title) 
	VALUES ($1)`

	_, err := m.DB.Exec(ctx, stmp, title)
	if err != nil {
		return err
	}

	return nil
}
