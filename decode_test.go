package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"ID":"1", "Name":"Akbar", "Age":26, "Married":false}`
	jsonByte := []byte(jsonRequest)
	data := Person{}
	err := json.Unmarshal(jsonByte, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.ID)
	fmt.Println(data.Name)
	fmt.Println(data.Age)
	fmt.Println(data.Married)
}
