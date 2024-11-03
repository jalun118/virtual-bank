package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"virtual-bank/internal/app"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/repository"
	"virtual-bank/internal/usecase"
)

func CreateTypeCurrentAccount(ctx context.Context, accountTypeUsecase *usecase.AccountTypeUsecase) {
	_, resNume := accountTypeUsecase.CreateType(ctx, models.CreateAccountType{
		Name: "current_account",
	})
	switch resNume {
	case constants.DuplicateRecord:
		fmt.Print("\n[!] Duplicate Type\n")
		return
	case constants.DatabaseError:
		fmt.Print("\n[!] Database Error\n")
		return
	}
	fmt.Print("\n[√] Success Insert\n")
}

func CreateTypeSavings(ctx context.Context, accountTypeUsecase *usecase.AccountTypeUsecase) {
	_, resNume := accountTypeUsecase.CreateType(ctx, models.CreateAccountType{
		Name: "savings",
	})
	switch resNume {
	case constants.DuplicateRecord:
		fmt.Print("\n[!] Duplicate Type\n")
		return
	case constants.DatabaseError:
		fmt.Print("\n[!] Database Error\n")
		return
	}
	fmt.Print("\n[√] Success Insert\n")
}

func main() {
	viper := app.NewViper()

	db := app.NewDatabase(viper)

	log := helpers.NewLogger(os.Stdout)

	accountTypeRepo := repository.NewAccountTypeRepository()
	accountTypeUsecase := usecase.NewAccountTypeUsecase(db, accountTypeRepo, log)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n+-----------------+\n| Installer Tools |\n+-----------------+\n\nSelect Option to run\n\n0. Exit CLI\n1. Create account type \"Current Account\"\n2. Create account type \"Savings\"\n3. Run options 1 and 2\n\nSelect your option: ")
	selection, _ := reader.ReadString('\n')
	if selection == "\r\n" {
		fmt.Print("\n[!] No Option To Select\n")
		return
	}
	selection = strings.Trim(selection, "\r\n")
	if selection == "0" {
		return
	}

	fmt.Print("Are you sure run to command: (Y/n) ")
	confirm, _ := reader.ReadString('\n')
	lowConfirm := strings.ToLower(strings.Trim(confirm, "\r\n"))

	switch lowConfirm {
	case "n":
		fmt.Print("\nHappy Nice Day!!\n")
		return
	default:
		lowConfirm = "y"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch selection {
	case "1":
		CreateTypeCurrentAccount(ctx, accountTypeUsecase)
	case "2":
		CreateTypeSavings(ctx, accountTypeUsecase)
	case "3":
		CreateTypeCurrentAccount(ctx, accountTypeUsecase)
		CreateTypeSavings(ctx, accountTypeUsecase)
	default:
		fmt.Println("\n[!] Select Not Found")
	}

}
