package api

import (
	"awesomeProject/internal/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (api *API) GetAllPersons(rw http.ResponseWriter, r *http.Request) (interface{}, error) {

	return api.v1Api.GetAllPersons()

}

func (api *API) GetPersonById(rw http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return api.v1Api.GetPersonById(int64(id))
}

func (api *API) AddPerson(rw http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	name := vars["name"]
	age, _ := strconv.Atoi(vars["age"])

	person := &model.Person{
		Id:   int64(id),
		Name: name,
		Age:  int64(age),
	}

	return api.v1Api.AddPerson(*person)
}
