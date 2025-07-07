package domain

import "Integrador/src/users/domain/entities"

type IUser interface {
	Save(user entities.User) error
	GetAll() ([]entities.User, error)
	GetById(id int) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	Edit(entities.User) error
	Delete(id int) error
}
