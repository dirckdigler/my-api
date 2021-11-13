package models

type Person struct {
	ID        int    `json:"ID,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
}
