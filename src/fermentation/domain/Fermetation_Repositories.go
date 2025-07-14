package domain

import "Integrador/src/fermentation/domain/entities"

type FermentationRepository interface {
	Save(fermentation entities.Fermentation) (entities.Fermentation, error)
	Update(fermentation entities.Fermentation) (entities.Fermentation, error)
	GetByID(id, userID int) (entities.Fermentation, error)
	GetAll(userID int) ([]entities.Fermentation, error)
	Delete(id, userID int) error
}
