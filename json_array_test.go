package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Array
func TestEncodeArray(t *testing.T) {
	persons := Person{
		ID:      "1",
		Name:    "Akbar",
		Hobbies: []string{"Gaming", "Reading"},
	}
	byte, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestDecodeArray(t *testing.T) {
	dataJSON := `{"ID":"1", "Name":"Akbar", "Hobbies":["Gaming", "Reading"]}`
	dataByte := []byte(dataJSON)
	persons := &Person{}
	err := json.Unmarshal(dataByte, persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(persons)
	fmt.Println(persons.ID)
	fmt.Println(persons.Name)
	fmt.Println(persons.Hobbies)
}

// Array Complex
func TestEncodeArrayComplex(t *testing.T) {
	persons := Person{
		ID:      "1",
		Name:    "Akbar",
		Hobbies: []string{"Gaming", "Reading"},
		Address: []Addresses{{Country: "Kota Bima", Street: "jl. Yos Sudarso", PostalCode: 112233}},
	}
	byte, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestDecodeArrayComplex(t *testing.T) {
	personJSON := `{"ID":"1", "Name":"Akbar", "Hobbies":["Gaming","Reading"], "Address": [{"Country":"Kota Bima", "Street":"jl. Yos Sudarso", "PostalCode":112233}]}`
	personByte := []byte(personJSON)
	persons := &Person{}
	err := json.Unmarshal(personByte, persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(persons.ID)
	fmt.Println(persons.Name)
	fmt.Println(persons.Hobbies)
	fmt.Println(persons.Address)
}
