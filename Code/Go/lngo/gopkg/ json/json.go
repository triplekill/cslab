package main

import (
	"encoding/json"
	"fmt"

	"github.com/exfly/lngo/model"
)

func unmarshalt() {
	var jsonBlob = []byte(`[
        {"_id":"2ojafp3324pih","Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)

	var animals []model.Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
func any() {
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &f)
	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int ", vv)
		case float64:
			fmt.Println(k, "is float64 ", vv)
		case []interface{}:
			fmt.Println(k, "is array:")
			for i, j := range vv {
				fmt.Println(i, j)
			}
		}
	}
}
func marshalt() {
	var animals []model.Animal
	animals = append(animals, model.Animal{Name: "Platypus", Order: "Monotremata"})
	animals = append(animals, model.Animal{Name: "Quoll", Order: "Dasyuromorphia"})

	jsonStr, err := json.Marshal(animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(jsonStr))
}

func anystar() {
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &f)
	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int ", vv)
		case float64:
			fmt.Println(k, "is float64 ", vv)
		case []interface{}:
			fmt.Println(k, "is array:")
			for i, j := range vv {
				fmt.Println(i, j)
			}
		}
	}
}
func main() {
	// unmarshalt()
	// any()
	// marshalt()
	anystar()
}
