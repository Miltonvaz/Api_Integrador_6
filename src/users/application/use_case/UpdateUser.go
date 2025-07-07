package use_case

import (
	"Integrador/src/users/domain"
	"Integrador/src/users/domain/entities"
)

type EditUser struct {
	db domain.IUser
}

func NewEditUser(db domain.IUser) *EditUser {
	return &EditUser{db: db}
}

func (ec *EditUser) Execute(client entities.User) error {
	return ec.db.Edit(client)
}
