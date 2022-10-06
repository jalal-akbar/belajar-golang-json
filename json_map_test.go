package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	var product = map[string]interface{}{
		"id":   "1",
		"name": "Akbar",
	}
	byte, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestMapDecode(t *testing.T) {
	dataJSON := `{"id":"1", "name":"Akbar"}`
	dataByte := []byte(dataJSON)
	var result map[string]interface{}
	err := json.Unmarshal(dataByte, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
}
