package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Streaming Encode
func TestStreamingEncode(t *testing.T) {
	writer, _ := os.Create("personOutput.json")
	encoder := json.NewEncoder(writer)

	p := Person{
		ID:      "1",
		Name:    "Akbar",
		Age:     26,
		Married: false,
		Hobbies: []string{"Gaming", "Reading"},
		Address: []Addresses{{Country: "Kota Bima", Street: "jln. Yos Sudarso", PostalCode: 112233}},
	}
	encoder.Encode(p)

	fmt.Println(p)

}

// Streaming Decode
func TestStreamingDecode(t *testing.T) {
	reader, err := os.Open("person.json")
	if err != nil {
		panic(err)
	}
	p := &Person{}
	decoder := json.NewDecoder(reader)
	decoder.Decode(p)
	fmt.Println(p)
}
