package main

import (
	"password/app-password/account"
	"password/app-password/encrypter"
	"password/app-password/files"
	"password/app-password/input"
	"password/app-password/output"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		output.PrintError("Переменные не найднены")
	}
	vault := account.NewVault(files.NewVaultDb("data.vault"), encrypter.NewEncrypter())

Menu:
	for {
		menuVariant := input.InputData(
			"__Меню__",
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите пункт меню",
		)
		switch menuVariant {
		case "1":
			createAccount(vault)
		case "2":
			searchAccounts(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break Menu
		default:
			output.PrintError("Неверный пункт меню")
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := input.InputData("Введите логин")
	password := input.InputData("Введите пароль")
	url := input.InputData("Введите URL")

	newAccount, err := account.NewAccount(login, password, url)

	output.PrintError(err)

	vault.AddAccount(newAccount)
}

func searchAccounts(vault *account.VaultWithDb) {
	search := input.InputData("Введите логин или URL для поиска аккаунта")
	accounts := vault.FindAccounts(func(acc account.Account) bool {
		return strings.Contains(acc.Url, search) || strings.Contains(acc.Login, search)
	})

	if len(accounts) == 0 {
		output.PrintError("Аккаунты не найдены")
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := input.InputData("Введите URL для поиска аккаунта")

	isDeleted, isFound := vault.DeleteByUrl(url)

	if isDeleted {
		output.PrintSuccess("Выбранные аккаунты успешно удалены")
		return
	}

	if !isDeleted && isFound {
		output.PrintWarning("Не выбраны аккаунты для удаления")
		return
	}

	output.PrintError("Аккаунты не найдены")
}
