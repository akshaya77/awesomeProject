package v1

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
)

type ApiV1 struct {
	db repository.DBRepository
}

func (a ApiV1) GetAllPersons() ([]model.Person, error) {
	return a.db.GetAllPersons()
}

func (a ApiV1) GetPersonById(id int64) (*model.Person, error) {
	return a.db.GetPersonById(id)
}

func (a ApiV1) AddPerson(person model.Person) (model.Person, error) {
	return a.db.AddPerson(person)
}