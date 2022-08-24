package main

import (
	"fmt"
	"projectapp/config"
	"projectapp/controllers/user"
	"projectapp/entities"
	"strconv"
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
			RowsAffected, err := user.InsertUserData(db, inputUser)
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
		{
			var updateUser = entities.User{}
			var AccountID int
			var str = strconv.Itoa(AccountID)
			fmt.Print("Masukan Nomor Telp anda untuk melanjutkan: ")
			fmt.Scanln(&updateUser.Contact)
			fmt.Println("Masukan Password anda: ")
			fmt.Scanln(&updateUser.Password)
			AccountID = user.LoginUserData(db, updateUser.Contact)
			passwordAccount := user.LoginUserData(db, str)
			validPass := user.LoginUserData(db, updateUser.Password)
			if AccountID == 0 {
				fmt.Println("Account anda tidak ditemukan")
			} else if passwordAccount != validPass {
				fmt.Println("Password anda salah")
			} else {
				fmt.Println("Account Terdaftar. Terima kasih sudah login ^-^")
			}
		}
	case 3:
		{
			result, err := user.GetUserData(db)
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
		var update string
		result, err := user.UpdateUser(db)
		if err != nil {
			update = "Nomor Handphone tidak ditemukan"
		} else {
			update = fmt.Sprint("Data User: ", "\n", result.Contact, "\n", "berhasil diupdate")
		}
		fmt.Println(update)
	}
}
