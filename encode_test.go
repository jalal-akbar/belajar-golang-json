package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJSON(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	LogJSON("Jalal")
	LogJSON(1)
	LogJSON(false)
	LogJSON([]string{"Jalaluddin", "Muh", "Akbar"})
}
