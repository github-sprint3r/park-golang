package main

import (
	"reflect"
	"testing"
)

func TestParkingCost20AndCash20ShouldReturnAllZero(t *testing.T) {
	expectMap := getDefaultCashList()
	actualMap := changeCalculate(20, 20)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}

func TestParkingCost0Cash0ShouldReturn0(t *testing.T) {
	expectMap := getDefaultCashList()
	result := changeCalculate(0, 0)
	eq := reflect.DeepEqual(expectMap, result)
	if eq != true {
		t.Errorf("Failed")
	}
}
