package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

var mut sync.Mutex

// метод для открытия базы данных
func CreateBD() *sql.DB {
	connStr := "user=postgres password=0 dbname=bet sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

// метод для проверки наличия данных о ставках в базе данных
func CheckStavki(user int) (int, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt1 := "select count(*) from Stavki where id_pol = $1"
	result1 := db.QueryRow(stmt1, user) //проверка: есть ли запись о балансе в таблице Polzovatel
	var n int
	er := result1.Scan(&n)
	if er != nil {
		mut.Unlock()
		return 0, er
	}
	mut.Unlock()
	return n, nil
}

// вставка информации о ставке в базу данных
func InsertStavki(Result string, BetSum float64, PotentialSum float64, user int) error {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "insert into Stavki(result, bet_sum, potential_sum, id_pol) values($1, $2, $3, $4)"
	_, err := db.Exec(stmt, Result, BetSum, PotentialSum, user)
	if err != nil {
		mut.Unlock()
		return err
	}
	mut.Unlock()
	return nil
}

// проверка наличия информации о балансе пользователя
func CheckPolzovatel(user int) (int, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt1 := "select count(*) from Polzovatel where id = $1"
	result1 := db.QueryRow(stmt1, user) //проверка: есть ли запись о балансе в таблице Polzovatel
	var n int
	er := result1.Scan(&n)
	if er != nil {
		mut.Unlock()
		return 0, er
	}
	mut.Unlock()
	return n, nil
}

// обновление баланса пользователя в базе данных
func UpdatePolzovatel(balance float64, user int) error {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "update Polzovatel set balance=$1 where id = $2" //если в Polzovatel есть запись о балансе, то значение баланса обновляется
	_, err := db.Exec(stmt, balance, user)
	if err != nil {
		mut.Unlock()
		return err
	}
	mut.Unlock()
	return nil
}

// извлечение информации о балансе пользователя из базы данных
func GetPolzovatel(user int) (float64, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stm := "select balance from Polzovatel where id = $1"
	var num1 float64
	e := db.QueryRow(stm, user).Scan(&num1)
	if e != nil {
		mut.Unlock()
		return 0, e
	}
	mut.Unlock()
	return num1, nil
}

func GetStavki(user int) (string, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stm := "select id, result, bet_sum, potential_sum from Stavki where id_pol = $1" //если в Stavki есть данные о ставках, то эти данные записываются в выходную структуру ListOfBets
	rows, err := db.Query(stm, user)
	if err != nil {
		mut.Unlock()
		return "", err
	}
	r := ""
	for rows.Next() {
		var s string
		var s1 string
		var s2 string
		var s3 string
		if err := rows.Scan(&s, &s1, &s2, &s3); err != nil {
			mut.Unlock()
			return "", err
		}
		s = s + " " + s1 + " " + s2 + " " + s3
		r += s + "\n"
	}
	mut.Unlock()
	return r, nil
}

func DeleteStavki(id int) error {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "delete from Stavki where id=$1" //удаление ставки из Stavki по id
	db.QueryRow(stmt, id)
	mut.Unlock()
	return nil
}
