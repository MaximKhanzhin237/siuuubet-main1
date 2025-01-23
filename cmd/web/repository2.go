package main

func InsertStavki(Result string, BetSum float64, PotentialSum float64, user int) error {
	mut1.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "insert into Stavki(result, bet_sum, potential_sum, id_pol) values($1, $2, $3, $4)"
	_, err := db.Exec(stmt, Result, BetSum, PotentialSum, user)
	if err != nil {
		mut1.Unlock()
		return err
	}
	mut1.Unlock()
	return nil
}

func UpdatePolzovatel(balance float64, user int) error {
	mut1.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "update Polzovatel set balance=$1 where id = $2" //если в Polzovatel есть запись о балансе, то значение баланса обновляется
	_, err := db.Exec(stmt, balance, user)
	if err != nil {
		mut1.Unlock()
		return err
	}
	mut1.Unlock()
	return nil
}
