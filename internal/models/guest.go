package models

type Guest struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Confirmed bool `json:"confirmed"`
}

// Guests is variable public
var Guests = []Guest{
	Guest{Name: "Samantha", Age: 25, Confirmed: true},
	Guest{Name: "Marcos", Age: 25, Confirmed: true},
	Guest{Name: "Ana", Age: 22, Confirmed: false},
}