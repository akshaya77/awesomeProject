package main

import "fmt"

type Person struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int64	`json:"age"`
}

func (p Person) String() string  {
	return fmt.Sprintf(" %v is %v years old and has id %v\n", p.Name, p.Age, p.Id)
}

type Persons []Person

type MyError struct {
	message string `json:"message"`
}

func (myError MyError) Error() string {
	return myError.message
}