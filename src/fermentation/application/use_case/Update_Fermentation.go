package use_case

import (
	"Integrador/src/fermentation/domain"
	"Integrador/src/fermentation/domain/entities"
)

type UpdateFermentation struct {
	db domain.FermentationRepository
}

func NewUpdateFermentation(db domain.FermentationRepository) *UpdateFermentation {
	return &UpdateFermentation{db: db}
}

func (u *UpdateFermentation) Execute(f entities.Fermentation) (interface{}, error) {
	return u.db.Update(f)
}
