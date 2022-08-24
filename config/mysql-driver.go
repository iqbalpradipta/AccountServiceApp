package config

import (
	"database/sql"
	"fmt"
	"log"

	// "os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	// dbConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", "root:@tcp(192.168.0.112:3306)/accountserviceapp")

	if err != nil {
		log.Fatal("error sql Open ", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("success connect to DB")
	}

	return db
}
