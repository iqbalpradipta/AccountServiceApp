package main

import (
	"fmt"
	"projectapp/config"
	_user "projectapp/controllers/user"
	_topup "projectapp/controllers/topup"
	"projectapp/entities"
	// "strconv"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()
	fmt.Print("MENU UTAMA: ")
	fmt.Print("\n1.Add Accout\n2.Login\n3.Read Account\n4.Update Account\n5.Delete Account\n6.TopUp Account\n7.Transfer\n8.History TopUp\n9.History Transfer\n10.Lihat Profil user lain\n0.Exit\n")
	fmt.Println("Masukan pilihan anda: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			inputUser := entities.User{}
			fmt.Println("Input id ")
			fmt.Scanln(&inputUser.Id)
			fmt.Println("Input user_id: ")
			fmt.Scanln(&inputUser.User_id)
			fmt.Println("Input Nama anda: ")
			fmt.Scanln(&inputUser.Name)
			fmt.Println("Input Password anda: ")
			fmt.Scanln(&inputUser.Password)
			fmt.Println("Input Alamat anda: ")
			fmt.Scanln(&inputUser.Alamat)
			fmt.Println("Input Jenis Kelamin: ")
			fmt.Println(&inputUser.Jenis_kelamin)
			fmt.Println("Input Nomor Telpon anda: ")
			fmt.Scanln(&inputUser.Contact)
			RowsAffected, err := _user.InsertUserData(db, inputUser)
			if err != nil {
				fmt.Println("Error Register data", err)
			} else {
				if RowsAffected == 0 {
					fmt.Println("Gagal Register data. RowsAffected = 0")
				} else {
					fmt.Println("Register data Berhasil. RowsAffected = ", RowsAffected)
				}
			}
		}
	case 2:
		loginUser := entities.User{}
		fmt.Println("Input Nomor Telpon anda: ")
		fmt.Scanln(&loginUser.Contact)
		fmt.Println("Input Password anda: ")
		fmt.Scanln(&loginUser.Password)
		result, err := _user.LoginUserData(db)
		if err != nil {
			fmt.Println("Error Login", err)
		} else {
			if err != nil {
				fmt.Println("login gagal")
			} else {
				for _, v := range result {
					fmt.Println("contact:", v.Contact)
				}
				fmt.Println("login Sukses. selamat datang")
			}
		}
	case 3:
		{
			result, err := _user.GetUserData(db)
			if err != nil {
				fmt.Println("Error membaca data dari database", err)
			} else {
				for _, v := range result {
					fmt.Println("id:", v.Id)
					fmt.Println("user_id:", v.User_id)
					fmt.Println("name:", v.Name)
					fmt.Println("password:", v.Password)
					fmt.Println("alamat:", v.Alamat)
					fmt.Println("jenis_kelamin:", v.Jenis_kelamin)
					fmt.Println("contact:", v.Contact)
					fmt.Println("saldo:", v.Saldo)
					fmt.Println("update_at:", v.Update_at)
				}
			}
		}
	case 4:
		{
			var updateData = entities.User{}
			fmt.Println("Name ID:")
			fmt.Scanln(&updateData.Id)
			fmt.Println("Update nama: ")
			fmt.Scanln(&updateData.Name)
			update, err := _user.UpdateData(db, updateData)
			if err != nil {
				fmt.Println("Error Update data", err)
			} else {
				if update == 0 {
					fmt.Println("Gagal update data. RowsAffected = 0")
				} else {
					fmt.Println("Update data Berhasil. RowsAffected = ", update)
				}
			}
		}
	case 5:
		{
			var deleteUser = entities.User{}
			fmt.Println("Delete Account dengan id:")
			fmt.Scanln(&deleteUser.Id)
			delete, err := _user.DeleteUserData(db, deleteUser)
			if err != nil {
				fmt.Println("Error Delete data", err)
			} else {
				if delete == 0 {
					fmt.Println("Gagal Delete data. RowsAffected = 0")
				} else {
					fmt.Println("Deleted data Berhasil. RowsAffected = ", delete)
				}
			}
		}
	case 6:
		var idAccount int
		var jumlah_top_up uint
			fmt.Print("Silahkan Masukkan Nominal Top Up: ")
			fmt.Scan(&jumlah_top_up)
			fmt.Print("\n")
			_, err := _topup.PostTopUp(db, idAccount, int(jumlah_top_up))
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Top Up Berhasil")
			}
			fmt.Print("\n")
	}
}
