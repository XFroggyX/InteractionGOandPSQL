package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (m *AssociationsModel) Insert(ctx context.Context, title string) error {
	stmp := `INSERT INTO Associations (Title) 
	VALUES ($1)`

	_, err := m.DB.Exec(ctx, stmp, title)
	if err != nil {
		return err
	}

	return nil
}

func (m *AssociationsModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM Associations WHERE id = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *AssociationsModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE Associations SET %s = `, nameFields)
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
