package jsonapiserializer

// StructuredError struct for structured Error
type StructuredError struct {
	Details string `json:"details,omitempty"`
	ID		string `json:"id,omitempty"`
	Source	string `json:"source,omitempty"`
	Status	string `json:"status,omitempty"`
	Title	string `json:"title,omitempty"`
}

// ResponseError is the json error response
type ResponseError struct {
    Errors []StructuredError `json:"errors"`
}

// JSONStruc JSONAPI Format
type JSONStruc struct {
    Type  		string      `json:"type"`
    ID    		string      `json:"id"`
    Attributes  interface{} `json:"attributes,omitempty"`
}

// ListStruc List Structure
type ListStruc struct {
    Data []JSONStruc `json:"data"`
}
