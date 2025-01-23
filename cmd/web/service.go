package main

import (
	"awesomeProject2/cmd/web/Builder"
	"awesomeProject2/cmd/web/Strategy"
	"fmt"
	"sync"
)

// создание контекста использования функций сервиса(в данном случае используются настоящие функции)
var ctx = new(Strategy.Context)

func init() {
	ctx.Algorithm(&Strategy.BetM{})
}

var mut2 sync.Mutex

// метод для вычисления потенциальной суммы выигрыша ставки и суммы баланса после создания ставки

// вычисление потенциальной суммы выигрыша ставки
func PotSum(odds, bet_sum float64) float64 {
	return odds * bet_sum
}

// вычисление баланса после регистрации ставки
func DecBal(balance, bet_sum float64) float64 {
	return balance - bet_sum
}

// проверка корректности введенных пользователем данных о ставке
func Check(balance, bet_sum float64) bool {
	if (balance >= 0) && (bet_sum >= 0) && (balance >= bet_sum) {
		return true
	} else {
		return false
	}
}

// вычисление баланса после удаления ставки
func InckBal(balance, bet_sum float64) float64 {
	return balance + bet_sum
}

// метод для удаления данных о ставке в базе данных
func Decrease(id int, bet_sum, balance float64, user int) (int, error) {
	var b float64
	if Check(balance, bet_sum) {
		b = InckBal(balance, bet_sum)
	} else {
		return 1, nil
	}
	mut2.Lock()
	err := ctx.Strategy.UpdatePolzovatel(b, user)
	mut2.Unlock()
	if err != nil {
		return 1, err
	}
	mut2.Lock()
	err = ctx.Strategy.DeleteStavki(id)
	mut2.Unlock()
	if err != nil {
		return 1, err
	}
	return 0, nil
}

// метод для извлечения данных о всех сделанных ставках из базы данных
func Get(user int) (Builder.ListOfBets, error) {
	result := Builder.ListOfBets{}
	mut2.Lock()
	b, err3 := ctx.Strategy.GetPolzovatel(user)
	mut2.Unlock()
	fmt.Println(b)
	if err3 != nil {
		return Builder.ListOfBets{}, err3
	}
	mut2.Lock()
	n, err4 := ctx.Strategy.CheckStavki(user)
	mut2.Unlock()
	if err4 != nil {
		return Builder.ListOfBets{}, err4
	}
	var r string
	if n != 0 {
		mut2.Lock()
		r, _ = ctx.Strategy.GetStavki(user)
		mut2.Unlock()
	} else {
		r = ""
	}
	result.Bets = r
	result.Balance = b
	return result, nil
}
