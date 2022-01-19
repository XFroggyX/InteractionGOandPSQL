package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (m *LanguagesModel) Insert(ctx context.Context, language string) error {
	stmp := `INSERT INTO Languages (Language) 
	VALUES ($1)`

	_, err := m.DB.Exec(ctx, stmp, language)
	if err != nil {
		return err
	}

	return nil
}

func (m *LanguagesModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM Languages WHERE id = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *LanguagesModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE Languages SET %s = `, nameFields)
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
