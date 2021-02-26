package main

import (
	"log"
	jsonapi "github.com/ichigoni/jsonapiserializer"
)

// UserStruc User Structure
type UserStruc struct {
	FirstName	string	`json:"first-name"`
	LastName	string	`josn:"last-name`
}

func main() {
	var userData1 UserStruc
	var userData2 UserStruc

	var jsonData1 jsonapi.JSONStruc
	var jsonData2 jsonapi.JSONStruc

	userData1.FirstName = "Alvin"
	userData1.LastName = "Chong"

	jsonData1.Type = "users"
	jsonData1.ID = "1"
	jsonData1.Attributes = userData1

	userData2.FirstName = "Ichigoni"
	userData2.LastName = "Don"

	jsonData2.Type = "users"
	jsonData2.ID = "2"
	jsonData2.Attributes = userData2

	// userList := []jsonapi.JSONStruc{jsonData1, jsonData2}
	// log.Println(string(jsonapi.SerializerList(userList)))
	 
 	log.Println(string(jsonapi.SerializerList([]jsonapi.JSONStruc{jsonData1, jsonData2})))

}
