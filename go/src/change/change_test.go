package change

import (
	"reflect"
	"testing"
)

func TestParkingCost20AndCash1990ShouldReturn1000WithOneAnd500WithOneAnd100WithFourAnd50WithOneAnd20WithOne(t *testing.T) {
	expectMap := getDefaultCashList()
	expectMap[1000] = 1
	expectMap[500] = 1
	expectMap[100] = 4
	expectMap[50] = 1
	expectMap[20] = 1
	expectMap[10] = 0
	actualMap := changeCalculate(20, 1990)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}

func TestParkingCost20AndCash1000ShouldReturn500WithOneAnd100With4And50WithOneAnd20WithOneAndTenWithOne(t *testing.T) {
	expectMap := getDefaultCashList()
	expectMap[500] = 1
	expectMap[100] = 4
	expectMap[50] = 1
	expectMap[20] = 1
	expectMap[10] = 1
	actualMap := changeCalculate(20, 1000)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}

func TestParkingCost20AndCash500ShouldReturn100With4And50With1And20WithOneAnd10WithOne(t *testing.T) {
	expectMap := getDefaultCashList()
	expectMap[100] = 4
	expectMap[50] = 1
	expectMap[20] = 1
	expectMap[10] = 1
	actualMap := changeCalculate(20, 500)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}

func TestParkingCost20AndCash100ShouldReturn50WithOneAnd20WithOneAnd10WithOne(t *testing.T) {
	expectMap := getDefaultCashList()
	expectMap[50] = 1
	expectMap[20] = 1
	expectMap[10] = 1
	actualMap := changeCalculate(20, 100)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}

func TestParkingCost20AndCash50ShouldReturn20WithOneAnd10With10WithOne(t *testing.T) {
	expectMap := getDefaultCashList()
	expectMap[20] = 1
	expectMap[10] = 1
	actualMap := changeCalculate(20, 50)
	assertResult := reflect.DeepEqual(expectMap, actualMap)

	if !assertResult {
		t.Errorf("Fail")
	}
}

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
