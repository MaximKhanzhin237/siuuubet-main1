package main

import "fmt"

func Count(balance, odds, bet_sum float64, res string, user int) (int, error) {
	fmt.Println("stavka")
	var b float64
	if Check1(balance, bet_sum) {
		b = DecBal1(balance, bet_sum)
	} else {
		return 1, nil
	}
	potSum := PotSum1(odds, bet_sum)
	mut2.Lock()
	err := InsertStavki(res, bet_sum, potSum, user)
	mut2.Unlock()
	if err != nil {
		return 1, err
	}
	fmt.Println("update")
	mut2.Lock()
	err = UpdatePolzovatel(b, user)
	mut2.Unlock()
	if err != nil {
		return 1, err
	}
	return 0, nil
}

func PotSum1(odds, bet_sum float64) float64 {
	return odds * bet_sum
}

// вычисление баланса после регистрации ставки
func DecBal1(balance, bet_sum float64) float64 {
	return balance - bet_sum
}

// проверка корректности введенных пользователем данных о ставке
func Check1(balance, bet_sum float64) bool {
	if (balance >= 0) && (bet_sum >= 0) && (balance >= bet_sum) {
		return true
	} else {
		return false
	}
}
