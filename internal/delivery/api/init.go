package api

import (
	v1 "awesomeProject/internal/usecase/v1"
	"github.com/gorilla/mux"
	"net/http"
)

func New(v1 v1.ApiV1) *API {
	api := &API{
		v1Api: v1,
	}

	return api.initHandlers()
}

func (api *API) initHandlers() *API {

	myRouter := mux.NewRouter()

	myRouter.Handle("/health", HandlerFunc(api.Health))
	myRouter.Handle("/getAllPersons", HandlerFunc(api.GetAllPersons)).Methods(http.MethodGet)
	myRouter.Handle("/getPersonById/{id}", HandlerFunc(api.GetPersonById)).Methods(http.MethodGet)
	myRouter.Handle("/addPersonById/{id}/{name}/{age}", HandlerFunc(api.AddPerson)).Methods(http.MethodGet)
	//log.Fatal(http.ListenAndServe(":8056", myRouter))
	http.Handle("/", myRouter)
	return api
}