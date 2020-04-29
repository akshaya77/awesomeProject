package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func main()  {
	handleRequests()
}

func handleRequests() {
	fmt.Println("Starting Server - Started")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/getAllPersons", getAllPersons).Methods("GET")
	myRouter.HandleFunc("/getPersonById/{id}", getPersonById).Methods("POST")
	log.Fatal(http.ListenAndServe(":8056", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"welcome to home page")
}

func getAllPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllPersons() endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPersonById(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("getPersonById() endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("query with id: ", id)
	var person Person
	for i:=0;i<4;i++ {
		if int64(id) == persons[i].Id {
			person = persons[i]
			break
		}
	}
	if reflect.DeepEqual(person, Person{}) {
		fmt.Println(fmt.Sprintf("No person with id %v",id))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MyError{message: fmt.Sprintf("No person with id %v",id)})
		return
	}
	json.NewEncoder(w).Encode(person)
}

var persons = Persons{
	Person{
		Id:   22,
		Name: "Akshaya",
		Age:  23,
	},
	Person{
		Id:   47,
		Name: "Likitha",
		Age:  24,
	},
	Person{
		Id:   77,
		Name: "Devi",
		Age:  20,
	},
	Person{
		Id:   18,
		Name: "Kohli",
		Age:  31,
	},
}