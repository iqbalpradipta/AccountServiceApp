package controllers

import (
	"database/sql"
	"projectapp/entities"
)

func GetIdUsersByTelp(db *sql.DB, Contact string) int {
	var query = "SELECT user_id FROM users WHERE telp = ?"
	var id int
	result := db.QueryRow(query, &Contact).Scan(&id)
	if result != nil {
		if result == sql.ErrNoRows {
			return -1
		}
		return -1
	}
	return id
}

func GetUserData(db *sql.DB) ([]entities.User, error) {
	var query = "SELECT Id, user_id, name, password, alamat, contact, saldo, update_at FROM user"
	result, errselect := db.Query(query)
	if errselect != nil {
		return nil, errselect
	}

	var userData []entities.User
	for result.Next() {
		var rowuser entities.User
		errScan := result.Scan(&rowuser.Id, &rowuser.User_id, &rowuser.Name, &rowuser.Password, &rowuser.Alamat, &rowuser.Contact, &rowuser.Saldo, &rowuser.Update_at)
		if errScan != nil {
			return nil, errScan
		}
		userData = append(userData, rowuser)
	}
	return userData, nil
}

func InsertUserData(db *sql.DB, inputUser entities.User) (int, error) {
	var query = "insert into User (id, User_id, Name, Password, Alamat, Contact, Saldo)values(?, ?, ?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errExec := statement.Exec(inputUser.Id, inputUser.User_id, inputUser.Name, inputUser.Password, inputUser.Alamat, inputUser.Contact, inputUser.Saldo)
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
func LoginUserData(db *sql.DB) ([]entities.User, error){
	var query = "SELECT contact, password FROM user"
	result, errselect := db.Query(query)
	if errselect != nil {
		return nil, errselect
	}
	var infoUser []entities.User
	for result.Next() {
		var rowuser entities.User
		errScan := result.Scan(&rowuser.Contact, &rowuser.Password)
		if errScan != nil {
			return nil, errScan
		}
		infoUser = append(infoUser, rowuser)
	}
	return infoUser, nil
}

func DeleteUserData(db *sql.DB, deleteUser entities.User) (int,error){
	var query = "delete from User WHERE id = ?"
	statement, err := db.Prepare(query)
	if err != nil {
		return -1 ,err
	}
	result, errExec := statement.Exec(deleteUser.Id)
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

func UpdateData(db *sql.DB, updateUser entities.User) (int, error)  {
	var query = "Update User SET Name = ? WHERE id = ?"
	statement, err := db.Prepare(query)
	if err != nil {
		return -1 ,err
	}
	result, errExec := statement.Exec(updateUser.Name, updateUser.Id)
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

func GetUserSaldo(db *sql.DB, idUser int) (int, error) {
	var query = "SELECT saldo FROM user WHERE id = ?"
	var saldo int
	result := db.QueryRow(query, &idUser).Scan(&saldo)
	if result != nil {
		if result == sql.ErrNoRows {
			return -1, result
		}
		return -1, result
	}
	return saldo, nil
}


func PostTambahSaldo(db *sql.DB, idUser int, nominal int) (int, error) {
	var query = "update user set saldo = (?) where id = (?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldo, errSaldo := GetUserSaldo(db, idUser)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var newSaldo = nominal + saldo
	result, err := statement.Exec(&newSaldo, &idUser)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func PostKurangSaldo(db *sql.DB, idUser int, jumlah_transfer int) (int, error) {
	var query = "update user set Saldo = (?) where user_id = (?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldo, errSaldo := GetUserSaldo(db, idUser)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var newSaldo = saldo - jumlah_transfer
	result, err := statement.Exec(&newSaldo, &idUser)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func ReadUserInfo(db *sql.DB, id string) entities.User {
	results := db.QueryRow("SELECT id, name, Contact, saldo from users where User_id = ?", &id)

	var dataUser entities.User
	err := results.Scan(&dataUser.Id, &dataUser.Name, &dataUser.Contact, &dataUser.Saldo)

	if err != nil {
		return entities.User{}
	}

	return dataUser
}