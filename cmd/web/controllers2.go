package main

import (
	"awesomeProject2/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusNoContent)
}

func (app *Application2) inf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("функция Count")
	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		handleOptions(w, r)
	}
	if r.Method == "POST" {
		fmt.Println("post inf")
		var bet models.Bet
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&bet)
		fmt.Println("decode")
		if err != nil {
			fmt.Println("error")
			http.Error(w, "Internal Server Error", 500)
			return
		}
		fmt.Println("count")
		flag, err = Count(bet.Balance, bet.Odds, bet.BetSum, bet.Result, message.Id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
	} else {
		fmt.Println("fkjrekofre")
		http.Error(w, "Method not allowed", 405)
	}
}
