package usecase

import "awesomeProject/internal/model"

type UseCase interface {
	GetAllPersons() ([]model.Person, error)
	GetPersonById(id int64) (*model.Person, error)
	AddPerson(person model.Person) (model.Person, error)
}