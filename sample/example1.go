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
	var userData UserStruc
	var jsonData jsonapi.JSONStruc

	userData.FirstName = "Alvin"
	userData.LastName = "Chong"

	jsonData.Type = "users"
	jsonData.ID = "1"
	jsonData.Attributes = userData

 	log.Println(string(jsonapi.Serializer(jsonData)))
}
