package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Object
type Person struct {
	ID      string
	Name    string
	Age     int
	Married bool
	Hobbies []string
	Address []Addresses
}

type Addresses struct {
	Country    string
	Street     string
	PostalCode int
}

func TestJSONObject(t *testing.T) {
	persons := Person{
		ID:      "1",
		Name:    "AKbar",
		Age:     26,
		Married: false,
	}
	byte, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}
