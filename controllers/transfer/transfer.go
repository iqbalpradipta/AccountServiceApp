package controllers
import (
	"database/sql"
	"fmt"
	_user "projectapp/controllers/user"
	"projectapp/entities"
)

func GetTransferData(db *sql.DB) ([]entities.Transfer, error){
	var query = "select id, user_id_pengirim, user_id_penerima, jumlah_transfer from Transfer"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return []entities.Transfer{}, errPrepare
	}
	result, err := statement.Query()
	if err != nil {
		return []entities.Transfer{}, err
	}
	var historyTransfer []entities.Transfer
	for result.Next() {
		var transfer entities.Transfer
		err := result.Scan(&transfer.Id, &transfer.User_id_pengirim, &transfer.User_id_penerima, &transfer.Jumlah_transfer)
		if err != nil {
			return nil, err
		}
		historyTransfer = append(historyTransfer, transfer)
	}
	return historyTransfer, nil
}

func PostTransfer(db *sql.DB, idPengirim int, idPenerima int, jumlah_transfer int) (int, error) {
	var query = "insert into transfer(user_id_pengirim, user_id_penerima, jumlah_transfer) values (?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldoPemberi, errSaldo := _user.GetUserSaldo(db, idPengirim)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var sisaSaldo int
	if saldoPemberi > jumlah_transfer && saldoPemberi > 10000 {
		sisaSaldo = saldoPemberi - jumlah_transfer
		_, errPostKurang := _user.PostKurangSaldo(db, idPengirim, jumlah_transfer)
		if errPostKurang != nil {
			return 0, errPostKurang
		}
	} else {
		fmt.Println("Saldo tidak mencukupi")
		sisaSaldo = saldoPemberi
		jumlah_transfer = 0
	}
	result, err := statement.Exec(&idPengirim, &idPenerima, jumlah_transfer, &sisaSaldo)
	_, errPostTambah := _user.PostTambahSaldo(db, idPenerima, jumlah_transfer)
	if errPostTambah != nil {
		return 0, errPostTambah
	}
	if err != nil {
		return 0, err
	} else {
		rowTopUp, _ := result.RowsAffected()
		return int(rowTopUp), nil
	}
}