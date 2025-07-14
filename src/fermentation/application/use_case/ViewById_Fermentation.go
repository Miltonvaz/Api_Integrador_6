package use_case

import "Integrador/src/fermentation/domain"

type GetByID_Fermentation struct {
	db domain.FermentationRepository
}

func NewGetByID_Fermentation(db domain.FermentationRepository) *GetByID_Fermentation {
	return &GetByID_Fermentation{db: db}
}

func (u *GetByID_Fermentation) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
