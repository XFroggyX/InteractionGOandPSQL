package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (m *СontinentsModel) Insert(ctx context.Context, name string) error {
	stmp := `INSERT INTO Сontinents (Name) 
	VALUES ($1)`

	_, err := m.DB.Exec(ctx, stmp, name)
	if err != nil {
		return err
	}

	return nil
}

func (m *СontinentsModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM Сontinents WHERE id = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *СontinentsModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE Сontinents SET %s = `, nameFields)
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
