package user

import (
	"database/sql"
	"projectapp/entities"
)

func GetUserData(db *sql.DB) ([]entities.User, error) {
	var query = "SELECT Id, user_id, name, password, alamat, jenis_kelamin, contact, saldo, update_at FROM user"
	result, errselect := db.Query(query)
	if errselect != nil {
		return nil, errselect
	}

	var userData []entities.User
	for result.Next() {
		var rowuser entities.User
		errScan := result.Scan(&rowuser.Id, &rowuser.User_id, &rowuser.Name, &rowuser.Password, &rowuser.Alamat, &rowuser.Jenis_kelamin, &rowuser.Contact, &rowuser.Saldo, &rowuser.Update_at)
		if errScan != nil {
			return nil, errScan
		}
		userData = append(userData, rowuser)
	}
	return userData, nil
}

func InsertUserData(db *sql.DB, inputUser entities.User) (int, error) {
	var query = "insert into User (User_id, Name, Password, Alamat, Jenis_kelamin, Contact, Saldo, update_at)values(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errExec := statement.Exec(inputUser.User_id, inputUser.Name, inputUser.Password, inputUser.Alamat, inputUser.Jenis_kelamin, inputUser.Contact, inputUser.Saldo, inputUser.Update_at)
	if errExec != nil {
		return -1, errExec
	}else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, nil
		}
		return int(row), nil
	}
}
