package account

import (
	"encoding/json"
	"fmt"
	"password/app-password/input"
	"password/app-password/output"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type VaultEncrypter interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc VaultEncrypter
}

func NewVault(db Db, enc VaultEncrypter) *VaultWithDb {
	var vault Vault
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	err = json.Unmarshal(data, &vault)

	output.PrintError(err)

	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) saveVaultToFile() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)

	output.PrintError(err)

	vault.db.Write(encData)
}

func (vault *VaultWithDb) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)

	vault.saveVaultToFile()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	return json.Marshal(vault)
}

func (vault *VaultWithDb) FindAccounts(checker func(acc Account) bool) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		isMatched := checker(acc)

		if isMatched {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func (vault *VaultWithDb) DeleteByUrl(url string) (bool, bool) {
	var accounts []Account
	isDeleted := false
	isFound := false

	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)

		if isMatched {
			fmt.Println("Найден аккаунт: ")
			acc.Output()
			answer := input.InputData("Подтвердите удаление (y/n)")

			switch answer {
			case "Y":
			case "y":
				isDeleted = true
				continue
			case "N":
			case "n":
				accounts = append(accounts, acc)
				isFound = true
				continue
			default:
				output.PrintError("Введена неверная команда")
				answer = input.InputData("Подтвердите удаление (y/n)")
			}
		}

		accounts = append(accounts, acc)
		continue
	}

	vault.Accounts = accounts
	vault.saveVaultToFile()

	return isDeleted, isFound
}
