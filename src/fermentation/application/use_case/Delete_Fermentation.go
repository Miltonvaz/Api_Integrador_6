package use_case

import "Integrador/src/fermentation/domain"

type Delete_Fermentation struct {
	db domain.FermentationRepository
}

func NewDelete_Fermentation(db domain.FermentationRepository) *Delete_Fermentation {
	return &Delete_Fermentation{db: db}
}

func (dt *Delete_Fermentation) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
