package main

import (
	"User/commands"
	"User/db"
	"User/pkg/modules"
	"database/sql"
	"fmt"
	"log"
	_"mattn/go-sqlite3"
)

func main() {
	dB, err := sql.Open("sqlite3", "dbUser")
	if err != nil {
		log.Fatal("Error", err)
	}
	db.CreateUser(dB)

	db.CreateAccount(dB)

	db.CreateCurrency(dB)

	// InsertUser(dB)

	curr := modules.Currency{
		Id:   0,
		Name: "UZB",
	}

	newUser := modules.User{
		Id:      0,
		Name:    "Sherzod",
		Surname: "Sodiqov",
		Age:     21,
		Gender:  "M",
		Remove:  false,
	}

	newAcc := modules.Account{
		Id:        0,
		UserId:    7,
		NumberAcc: 222333,
		Amount:    2500,
		Currency:  3,
		Remove:    false,
	}

	err = commands.AddCurr(dB, curr)
	if err != nil {
		fmt.Println(err)
	}

	err = commands.AddNewUser(dB, newUser)
	if err != nil{
		fmt.Println(err)
	}

	err = commands.AddNewAccount(dB, newAcc)
	if err != nil {
		fmt.Println(err)
	}
}



func InsertUser(dataBase *sql.DB) (err error) {
	const AddUser = `insert into users values(1,'Tolmas','Farmonov',22,'M',false)`
	_, err = dataBase.Exec(AddUser)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}