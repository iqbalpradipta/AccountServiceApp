package controllers

import (
	"database/sql"
	"fmt"
	"projectapp/controllers/user"
	"projectapp/entities"
)

func GetAllTopUp(db *sql.DB) ([]entities.Topup, error) {
	var query = "select Topup , user_id , jumlah_top_up from topup"
	result, errselect := db.Query(query)
	if errselect != nil {
		return nil, errselect
	}

	var userData []entities.Topup
	for result.Next() {
		var rowuser entities.Topup
		errScan := result.Scan()
		if errScan != nil {
			return nil, errScan
		}
		userData = append(userData, rowuser)
	}
	return userData, nil
}

func GetHistoryTopUpById(db *sql.DB, idUser int) ([]entities.Topup, error) {
	var query = "select id , user_id , jumlah_top_up from Topup where id = ? order by id desc"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return []entities.Topup{}, errPrepare
	}
	result, err := statement.Query(&idUser)
	if err != nil {
		return []entities.Topup{}, err
	}
	var historyTopUp = []entities.Topup{}
	for result.Next() {
		var topup = entities.Topup{}
		err := result.Scan(&topup)
		if err != nil {
			return []entities.Topup{}, err
		}
		historyTopUp = append(historyTopUp, topup)
	}
	return historyTopUp, nil
}

func PostTopUp(db *sql.DB, idUser int, jumlah_top_up int) (int, error) {
	var query = "insert into Topup (id,jumlah_top_up) values (?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	_, errSaldo := controllers.PostTambahSaldo(db, idUser, jumlah_top_up)
	if errSaldo != nil {
		return 0, errSaldo
	}
	result, err := statement.Exec(&idUser, &jumlah_top_up)
	if err != nil {
		return 0, err
	} else {
		rowTopUp, _ := result.RowsAffected()
		fmt.Println("Saldo anda bertambah")
		return int(rowTopUp), nil
	}
}