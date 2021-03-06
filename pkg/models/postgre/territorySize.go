package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TerritorySizesModel struct {
	DB *pgxpool.Pool
}

func (m *TerritorySizesModel) NameField() []string {
	return []string{"ID", "Type"}
}

func (m *TerritorySizesModel) Get(ctx context.Context) ([]model.TerritorySizes, error) {
	stmp := `SELECT * FROM TerritorySizes`
	rows, err := m.DB.Query(ctx, stmp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storage []model.TerritorySizes
	for rows.Next() {
		var c model.TerritorySizes
		values, err := rows.Values()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrNoRecord
			}
			return nil, err
		}

		c.ID = int(values[0].(int16))
		c.Type = values[1].(string)

		storage = append(storage, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (m *TerritorySizesModel) Insert(ctx context.Context, type_ string) error {
	stmp := `INSERT INTO TerritorySizes (type) 
	VALUES ($1)`

	_, err := m.DB.Exec(ctx, stmp, type_)
	if err != nil {
		return err
	}

	return nil
}

func (m *TerritorySizesModel) Delete(ctx context.Context, id int) error {
	stmp := `DELETE FROM TerritorySizes WHERE id = $1`
	_, err := m.DB.Exec(ctx, stmp, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNoRecord
		}
		return err
	}
	return nil
}

func (m *TerritorySizesModel) Update(ctx context.Context, id int, nameFields string, value string) error {
	stmp := fmt.Sprintf(`UPDATE TerritorySizes SET %s = `, nameFields)
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
