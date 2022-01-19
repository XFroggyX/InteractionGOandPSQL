package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GovernmentFormsModel struct {
	DB *pgxpool.Pool
}

func (m *GovernmentFormsModel) NameField() []string {
	return []string{"ID", "Form"}
}

func (m *GovernmentFormsModel) Get(ctx context.Context) ([]model.GovernmentForms, error) {
	stmp := `SELECT * FROM GovernmentForms`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.GovernmentForms
	for rows.Next() {
		var c model.GovernmentForms
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int16))
		c.Form = values[1].(string)

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (m *GovernmentFormsModel) Insert(ctx context.Context, form string) error {
	stmp := `INSERT INTO GovernmentForms (Form) 
	VALUES ($1)`

	_, err := m.DB.Exec(ctx, stmp, form)
	if err != nil {
		return err
	}

	return nil
}

func (m *GovernmentFormsModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM GovernmentForms WHERE id = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *GovernmentFormsModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE GovernmentForms SET %s = `, nameFields)
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
