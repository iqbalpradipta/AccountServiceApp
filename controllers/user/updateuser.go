package user

import (
	"database/sql"
	"fmt"
	"projectapp/entities"
)

func UpdateUser(db *sql.DB) (entities.User, error) {
	fmt.Println("Pilih Menu Update: ")
	fmt.Print("1.Update nama\n2.Update Email\n3.Update Password\n")
	var menu,UpdateNama,UpdateJenisKelamin,UpdatePassword string
	fmt.Scanln(&menu)
	var user entities.User

	switch menu {
		case "1":
		{
			fmt.Println("Masukan Nama Baru: ")
			fmt.Scanln(&UpdateNama)
			
			user.Name = UpdateNama
		}
		case "2":
		{
			fmt.Println("Update Jenis Kelamin: ")
			fmt.Scanln(&UpdateJenisKelamin)

			user.Jenis_kelamin = UpdateJenisKelamin
		}
		case "3":
		{
			fmt.Println("Masukan Password Baru: ")
			fmt.Scanln(&UpdatePassword)

			user.Password = UpdatePassword
		}
	}
	return user, nil
}