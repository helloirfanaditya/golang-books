package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func SetDB() {
	U := "postgres"
	PS := "123456"
	H := "localhost"
	PORT := "5432"
	DBNAME := "books"

	setDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", H, PORT, U, PS, DBNAME)
	Db, err = sql.Open("postgres", setDB)

	if err != nil {
		panic(err)
	}

	if err = Db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("connected ")
}

func Close() {
	Db.Close()
}
