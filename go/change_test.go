package main

import (
	"testing"
	"reflect"
)

func TestParkingCost20AndCash20ShouldReturnAllZero(t *testing.T) {
	expectMap := map[float64]int {
		1000: 0,
		500: 0,
		100: 0,
		50: 0,
		20: 0,
		10: 0,
		5: 0,
		2: 0,
		1: 0,
		0.50: 0,
		0.25: 0,
	}

	actualMap := changeCalculate(20, 20)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}