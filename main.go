package main

import (
	"fmt"
	"projectapp/config"
)

func main() {
	db:= config.ConnectToDB()
	defer db.Close()
	fmt.Print("MENU UTAMA: ")
	fmt.Print("\n1.Add Accout\n2.Login\n3.Read Account\n4.Update Account\n5.Delete Account\n6.TopUp Account\n7.Transfer\n8.History TopUp\n9.History Transfer\n10.Lihat Profil user lain\n0.Exit\n")
	fmt.Println("Masukan pilihan anda: ")
	var pilihan int
	fmt.Scanln(&pilihan)
}
