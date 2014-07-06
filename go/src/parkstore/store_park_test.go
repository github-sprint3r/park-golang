package parkstore

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
)

func TestStoreParkResponeOK(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://1.2.3.4/parkStore", nil)
	response := httptest.NewRecorder()

	ParkStore(response, request)

	if response.Code != 200 {
		t.Errorf("expect 200 but was %d", response.Code)
	}
}

func TestStoreParkReturnValue(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://1.2.3.4/parkStore", nil)
	response := httptest.NewRecorder()

	ParkStore(response, request)

	var responseJson = map[string]interface{} {
		"chargeRates": []ChargeRate{
			ChargeRate{Rate: "r1", Value: 0},
		},
	}

	jsonByte, _ := json.Marshal(responseJson)

	if !bytes.Equal([]byte(response.Body.String()), jsonByte)  {
		t.Errorf("expect String but return %s", response.Body.String())
	}
	
}
