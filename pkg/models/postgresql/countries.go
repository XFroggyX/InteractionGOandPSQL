package postgresql

import (
	model "github.com/XFroggyX/InteractionGOandPSQL/pkg/models"
	"github.com/jackc/pgx/pgxpool"
)

type CountriesModel struct {
	DB *pgxpool.Pool
}

func (m *CountriesModel) Insert(CountriesName string, Flag string, ReligionID int, LanguagesId int,
	GovernmentFormID int, TerritorySizeID int) (int, error) {
	return 0, nil
}

func (m *CountriesModel) Get(id int) (*model.Countries, error) {
	return nil, nil
}
