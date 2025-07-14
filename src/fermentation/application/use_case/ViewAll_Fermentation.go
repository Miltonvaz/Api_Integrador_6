package use_case

import (
	"Integrador/src/fermentation/domain"
	"Integrador/src/fermentation/domain/entities"
)

type Get_All_Fermentation struct {
	db domain.FermentationRepository
}

func NewGet_All_Fermentation(db domain.FermentationRepository) *Get_All_Fermentation {
	return &Get_All_Fermentation{db: db}
}

func (gt *Get_All_Fermentation) Execute(userID int) ([]entities.Fermentation, error) {
	return gt.db.GetAll(userID)
}
