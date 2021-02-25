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

Sample1
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
```
