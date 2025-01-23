package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	Id int `json:"id"`
}

func sendMessage(server2URL string, message Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("не получилось обработать сообщение: %w", err)
	}

	resp, err := http.Post(server2URL, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		return fmt.Errorf("не получилось отправить сообщение на сервер: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("сервер вернул ошибку: %d, body: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (app *Application1) handleRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ts, err := template.ParseFiles("./ui/html/register.html")

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	if r.Method == http.MethodPost {
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}
		if GetPolzovatel(newUser) > 0 {
			respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Пользователь уже существует"})
			return
		}

		Id := InsertPolzovatel(newUser)
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Успешно зарегистрирован"})
		server2URL := "http://localhost:4040/"
		message1 := Message{Id: Id}
		err = sendMessage(server2URL, message1)
		if err != nil {
			log.Println("Ошибка отправки сообщения на server2:", err)
		} else {
			log.Println("Сообщение отправлено на server2")
		}
	}
}

func (app *Application1) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ts, err := template.ParseFiles("./ui/html/authorize.html")

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	if r.Method == http.MethodPost {
		var loginUser User
		fmt.Println("login post")
		err := json.NewDecoder(r.Body).Decode(&loginUser)
		if err != nil {
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}
		Id := GetPolzovatel(loginUser)
		if Id == 0 {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Неверные данные"})
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Успешный вход"})
		server2URL := "http://localhost:4040/"
		message1 := Message{Id: Id}
		err = sendMessage(server2URL, message1)
		if err != nil {
			log.Println("Ошибка отправки сообщения на server2:", err)
		} else {
			log.Println("Сообщение отправлено на server2")
		}
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
