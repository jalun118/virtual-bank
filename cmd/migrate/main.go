package main

import (
	"fmt"
	"virtual-bank/internal/app"
	"virtual-bank/internal/entity"
)

func main() {
	viper := app.NewViper()

	db := app.NewDatabase(viper)

	errMigate := db.AutoMigrate(
		&entity.Session{},
		&entity.User{},
		&entity.AccountType{},
		&entity.Account{},
		&entity.Transaction{},
	)

	if errMigate != nil {
		fmt.Println(errMigate)
	}
}
