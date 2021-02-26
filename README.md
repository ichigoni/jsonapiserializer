# jsonapiserializer
## JSON API Serializer for GO

## Installation

Using `go get` to install package, it is common for a package to be referenced by its canonical path, you would use the full canonical path:

```
$ go get github.com/ichigoni/jsonapiserializer
```

Manual install package by modify `go.mod` file

```GO
module ...

go ...

require (

...
github.com/ichigoni/jsonapiserializer v0.0.1
...

)

```

Using `go mod vendor` to include the package files.

```
$ go mod vendor
```

## Sample Usage

Example 1 
---
```
{
  "type": "users",
  "id: "1",
  "attributes": {
    "first-name": "Alvin",
    "last-name": "Chong"
  }
}
```


```GO
package main

import (
	"log"
	jsonapi "github.com/ichigoni/jsonapiserializer"
)

// UserStruc User Structure
type UserStruc struct {
	FirstName	string	`json:"first-name"`
	LastName	string	`json:"last-name`
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
```

Example 2 - List
---
```
{
  "data":[
    {
      "type": "users",
      "id: "1",
      "attributes": {
        "first-name": "Alvin",
        "last-name": "Chong"
      }
    },
    {
      "type": "users",
      "id: "2",
      "attributes": {
        "first-name": "Ichigoni",
        "last-name": "Chong"
      }
    }
  ]
}
```


```GO
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

```

