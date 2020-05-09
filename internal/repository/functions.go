package repository

import (
	"awesomeProject/internal/model"
)

func (D *DBConnection) GetAllPersons() ([]model.Person, error) {
	return D.redis.GetAllPersons()
}

func (D *DBConnection) GetPersonById(id int64) (*model.Person, error) {
	//for _, v := range Persons {
	//	if v.Id == id {
	//		return &v, nil
	//	}
	//}
	//return nil,errors.New("no person with id present")
	return D.redis.GetPersonById(id)
}

func (D *DBConnection) AddPerson(person model.Person) (model.Person, error) {
	//Persons = append(Persons, person)
	//return person, nil
	return D.redis.AddPerson(person)
}