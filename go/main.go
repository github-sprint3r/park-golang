package main

func getDefaultCashList() map[float64]int {
	defaultCashList := map[float64]int {
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

	return defaultCashList
}

func changeCalculate(parkingCosting int, cash int) map[float64]int{
	cashList := getDefaultCashList()

	return cashList
}