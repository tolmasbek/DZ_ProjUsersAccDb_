package commands

import (
	"User/pkg/modules"
	"database/sql"
	"fmt"
)

func AddNewUser(dataBase *sql.DB, users modules.User) (err error) {
	_, err = dataBase.Exec(`insert into users (nameUser, surname, age, gender, remove) values(($1),($2),($3),($4),($5))`, users.Name, users.Surname, users.Age, users.Gender, users.Remove)
	if err !=nil{
		fmt.Println(err)
	}
	return err
}

func AddNewAccount(dataBase *sql.DB, accounts modules.Account)(err error){
	_, err = dataBase.Exec(`insert into accounts (userId, numberAcc, amount, currency, remove) values(($1),($2),($3),($4),($5))`, accounts.UserId, accounts.NumberAcc, accounts.Amount, accounts.Currency, accounts.Remove)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func AddCurr(dataBase *sql.DB, currency modules.Currency)(err error){
	_, err = dataBase.Exec(`insert into currency(nameCurr)values (($1))`, currency.Name)
	if err != nil{
		fmt.Println(err)
	}
	return
}
