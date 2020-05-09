package redis

import (
	"awesomeProject/internal/model"
	"encoding/json"
	"fmt"
)

type RedisDb struct {

}

var RedisClient *RedisConn

func RedisInit()  {

	RedisClient = InitRedis("localhost:6379")

	err := RedisClient.PingRedis()

	if err != nil {
		fmt.Println("redis connection not possible")
	} else {
		fmt.Println("redis connection established.. Hurray!!!")
	}

}

func (r RedisDb) GetAllPersons() ([]model.Person, error) {
	values := RedisClient.GetConnection("my-redis").Connection().HVals("myMap").Val()
	//json.NewEncoder()
	fmt.Println(values)

	list := make([]model.Person,0,10)

	for _, i2 := range values {
		var p model.Person
		json.Unmarshal([]byte(i2), &p)
		list = append(list, p)
	}

	return list, nil

}

func (r RedisDb) GetPersonById(id int64) (*model.Person, error) {
	buf,err := RedisClient.GetConnection("my-redis").Connection().HGet("myMap", string(id)).Bytes()
	if err != nil {
		return nil, err
	}

	var p model.Person
	json.Unmarshal(buf, &p)

	return &p,nil
}

func (r RedisDb) AddPerson(person model.Person) (model.Person, error) {
	buf,err := json.Marshal(person)
	if err != nil {
		return model.Person{}, err
	}
	RedisClient.GetConnection("my-redis").Connection().HSet("myMap", string(person.Id), buf)
	return person,nil
}


