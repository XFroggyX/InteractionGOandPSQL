package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (m *ReligionsModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM Religions WHERE id = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *ReligionsModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE Religions SET %s = `, nameFields)
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
