package repository

import "awesomeProject/internal/model"

type DBRepository interface {
	GetAllPersons() ([]model.Person, error)
	GetPersonById(id int64) (*model.Person, error)
	AddPerson(person model.Person) (model.Person, error)
}

