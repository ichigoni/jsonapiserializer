package jsonapiserializer

import (
	"encoding/json"
	"log"
)

// Error is jsonapi serializer
func Error(theError StructuredError) []byte {
	resp := &ResponseError{
		Errors: []StructuredError{theError}}

	return Serializer(resp)
}

// Serializer is jsonapi serializer
func Serializer(v interface{}) []byte {
	JSON, err := json.Marshal(v)
	if err != nil {
		log.Println("Structured logger: Logger JSON Marshal failed !", err.Error())
	}

	return []byte(string(JSON));	
}
