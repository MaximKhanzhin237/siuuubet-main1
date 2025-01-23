package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Application struct {
}

type Application1 struct{}

type Application2 struct{}

var App = &Application{}

var App1 = &Application1{}

var App2 = &Application2{}

// основной сервер(выводит страницу для регистрации ставок пользователем, также удаляет выбранные ставки
func startServer1() {
	srv := &http.Server{
		Addr:    ":4040",
		Handler: App.routes(),
	}

	log.Println("Запуск сервера на 4040")
	err1 := srv.ListenAndServe()
	log.Fatal(err1)
}

// сервер для регистрации/авторизации
func startServer2() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: App1.routes(),
	}

	log.Println("Запуск сервера на 8080")
	err1 := srv.ListenAndServe()
	log.Fatal(err1)
}

// Сервер для высчитывания потенциальной суммы ставки и сохранения сделанной ставки в базу данных
func startServer3() {
	srv := &http.Server{
		Addr:    ":8081",
		Handler: App2.routes(),
	}

	log.Println("Запуск сервера на 8081")
	err1 := srv.ListenAndServe()
	log.Fatal(err1)
}

func main() {
	// Start each server in a separate goroutine
	go startServer1()
	go startServer2()
	go startServer3()

	// Give the servers some time to start before blocking the main goroutine
	time.Sleep(2 * time.Second)

	fmt.Println("Press Ctrl+C to stop the servers.")
	select {} // block forever
}
