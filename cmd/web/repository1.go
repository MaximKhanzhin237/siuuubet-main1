package main

import (
	"database/sql"
	"sync"
)

var mut1 sync.RWMutex

func CreateBD() *sql.DB {
	connStr := "user=postgres password=0 dbname=bet sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func GetPolzovatel(u User) int {
	mut1.Lock()
	db := CreateBD()
	defer db.Close()
	login := u.Username
	password := u.Password
	stmt1 := "select id from Polzovatel where login = $1 and password = $2"
	result1 := db.QueryRow(stmt1, login, password) //проверка: есть ли запись о балансе в таблице Polzovatel
	var n int
	er := result1.Scan(&n)
	if er != nil {
		mut1.Unlock()
		return 0
	}
	mut1.Unlock()
	return n
}

func InsertPolzovatel(u User) int {
	mut1.Lock()
	db := CreateBD()
	defer db.Close()
	login := u.Username
	password := u.Password
	stmt1 := "insert into Polzovatel (login, password, balance) values ($1, $2, 1000)"
	_, err := db.Exec(stmt1, login, password)
	if err != nil {
		mut1.Unlock()
		return 0
	}
	stmt2 := "select id from Polzovatel where login = $1 and password = $2"
	result2 := db.QueryRow(stmt2, login, password)
	var n int
	result2.Scan(&n)
	mut1.Unlock()
	return n
}
