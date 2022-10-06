package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func TestEncodeJSONTag(t *testing.T) {
	products := Product{
		ID:       "1",
		Name:     "Amer",
		Price:    150000,
		Quantity: 10,
	}
	byte, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestDecodeJSONTag(t *testing.T) {
	dataJSON := `{"id":"1","name":"Amer","price":150000,"quantity":100}`
	dataByte := []byte(dataJSON)
	products := &Product{}
	err := json.Unmarshal(dataByte, products)
	if err != nil {
		panic(err)
	}
	fmt.Println(products.ID)
	fmt.Println(products.Name)
	fmt.Println(products.Price)
	fmt.Println(products.Quantity)

}
