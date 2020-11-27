package db

import (
	"database/sql"
	"fmt"
)

func CreateUser(dataBase *sql.DB) {
	const UserTable = `create table if not exists users(
		id 			integer primary key autoincrement,
		nameUser 	text 	not null,
		surname 	text 	not null,
		age 		integer not null,
		gender 		text 	not null,
		remove 		boolean not null 	default false )`

	_, err := dataBase.Exec(UserTable)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateAccount(dataBase *sql.DB) {
	const AccTable = `create table if not exists accounts(
		id 			integer primary key autoincrement,
		userId 		integer references users(Id) not null,
		numberAcc 	integer not null,
		amount 		integer not null,
		currency 	integer references currency(Id),
		remove 		boolean not null default false )`

	_, err := dataBase.Exec(AccTable)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateCurrency(dataBase *sql.DB) {
	const CurrTable = `create table if not exists currency(
		id 			integer primary key autoincrement,
		nameCurr 	text )`

	_, err := dataBase.Exec(CurrTable)
	if err != nil {
		fmt.Println(err)
	}
}




