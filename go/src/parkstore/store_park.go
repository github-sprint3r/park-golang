package parkstore

import (
	"net/http"
	"fmt"
)

type ChargeRate struct {
	Rate string `json:"rate"`
	Value int `json:"value"`
}

func ParkStore(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, `{"chargeRates":[{"rate":"r1","value":0}]}`)
}
