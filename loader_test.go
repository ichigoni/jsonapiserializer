package jsonapiserializer

import (
	// "bytes"
	// "os"
	// "os/exec"
	"regexp"
	"testing"
)

const (
	errorDetails	= "Error Details"
	errorTitle		= "Error Title"
	errorSource		= "Error Source"
	errorStatus		= "Error Status"

	test1JsonType = "users"
	test1FirstName = "Alvin"
	test1LastName = "Chong"
	test2FirstName = "Ichigoni"
	test2LastName = "Don"
)
// UserStruc User Structure
type UserStruc struct {
	FirstName	string	`json:"first-name"`
	LastName	string	`json:"last-name`
}

func TestError(t *testing.T) {
	e := StructuredError{
		Details: 	errorDetails,
		Title: 		errorTitle,
		Source: 	errorSource,
		Status:		errorStatus}
	regexps := []*regexp.Regexp{
		regexp.MustCompile(errorDetails),
		regexp.MustCompile(errorTitle),
		regexp.MustCompile(errorSource),
		regexp.MustCompile(errorStatus),
	}
	output := string(Error(e))
	if !regexps[0].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", errorDetails)
	}
	if !regexps[1].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", errorTitle)
	}
	if !regexps[2].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", errorSource)
	}
	if !regexps[3].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", errorStatus)
	}
}

func TestSerializer(t *testing.T) {
	var userData UserStruc
	var jsonData JSONStruc

	userData.FirstName = test1FirstName
	userData.LastName = test1LastName

	jsonData.Type = test1JsonType
	jsonData.ID = "1"
	jsonData.Attributes = userData

	output := string(Serializer(jsonData))
	regexps := []*regexp.Regexp{
		regexp.MustCompile(test1FirstName),
		regexp.MustCompile(test1LastName),
		regexp.MustCompile(test1JsonType),
		regexp.MustCompile("attributes"),
		regexp.MustCompile("id"),
	}
	if !regexps[0].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1FirstName)
	}
	if !regexps[1].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1LastName)
	}
	if !regexps[2].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1JsonType)
	}
	if !regexps[3].MatchString(output) {
		t.Errorf("The output doesn't contain attributes")
	}
	if !regexps[4].MatchString(output) {
		t.Errorf("The output doesn't contain ID")
	}
}

func TestSerializer_undefineType(t *testing.T) {
	var userData UserStruc
	var jsonData JSONStruc

	userData.FirstName = test1FirstName
	userData.LastName = test1LastName

	jsonData.Type = test1JsonType
	jsonData.ID = "1"
	// jsonData.Attributes = userData

	output := string(Serializer(jsonData))
	regexps := []*regexp.Regexp{
		regexp.MustCompile("attributes"),
		regexp.MustCompile("id"),
	}
	if regexps[0].MatchString(output) {
		t.Errorf("The output didn't return the expected error. Missing attributes")
	}
}

func TestSerializer_undefineAttributes(t *testing.T) {
	var userData UserStruc
	var jsonData JSONStruc

	userData.FirstName = test1FirstName
	userData.LastName = test1LastName

	jsonData.ID = "1"
	jsonData.Attributes = userData

	output := string(Serializer(jsonData))
	regexps := []*regexp.Regexp{
		regexp.MustCompile(test1FirstName),
		regexp.MustCompile(test1LastName),
		regexp.MustCompile(test1JsonType),
		regexp.MustCompile("attributes"),
		regexp.MustCompile("id"),
	}
	if !regexps[0].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1FirstName)
	}
	if !regexps[1].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1LastName)
	}
	if !regexps[3].MatchString(output) {
		t.Errorf("The output doesn't contain attributes")
	}
	if !regexps[4].MatchString(output) {
		t.Errorf("The output doesn't contain ID")
	}
	if regexps[2].MatchString(output) {
		t.Errorf("The output didn't return the expected error. Missing TYPE: %s", test1JsonType)
	}
}

func TestSerializerList(t *testing.T) {
	var userData1 UserStruc
	var userData2 UserStruc
	var jsonData1 JSONStruc
	var jsonData2 JSONStruc

	userData1.FirstName = test1FirstName
	userData1.LastName = test1LastName

	userData2.FirstName = test2FirstName
	userData2.LastName = test2LastName

	jsonData1.Type = test1JsonType
	jsonData1.ID = "1"
	jsonData1.Attributes = userData1

	jsonData2.Type = "users"
	jsonData2.ID = "2"
	jsonData2.Attributes = userData2

	output := string(SerializerList([]JSONStruc{jsonData1, jsonData2}))
	regexps := []*regexp.Regexp{
		regexp.MustCompile(test1FirstName),
		regexp.MustCompile(test2FirstName),
		regexp.MustCompile(test1JsonType),
		regexp.MustCompile("attributes"),
		regexp.MustCompile("id"),
	}
	if !regexps[0].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1FirstName)
	}
	if !regexps[1].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test2FirstName)
	}
	if !regexps[2].MatchString(output) {
		t.Errorf("The output doesn't contain the expected %s for the output", test1JsonType)
	}
	if !regexps[3].MatchString(output) {
		t.Errorf("The output doesn't contain attributes")
	}
	if !regexps[4].MatchString(output) {
		t.Errorf("The output doesn't contain ID")
	}
}
