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

func calculateRemain(remain int, cashType int) (int, int) {
	numberOfCash := 0
	remain_result := remain
	divide_result := remain / cashType
	if divide_result > 0 {
		numberOfCash = divide_result
		remain_result = remain % cashType
	}

	return numberOfCash, remain_result
}

func changeCalculate(parkingCosting int, cash int) map[float64]int{
	cashList := getDefaultCashList()

	remain := cash - parkingCosting

	cashList[50], remain = calculateRemain(remain, 50)
	cashList[20], remain = calculateRemain(remain, 20)
	cashList[10], remain = calculateRemain(remain, 10)

	return cashList
}