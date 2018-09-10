package test

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
}
