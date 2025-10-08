package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*#")

type Account struct {
	CreatedAt time.Time `json:"createdAt"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url       string    `json:"url"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Неверный URL")
	}

	if login == "" {
		return nil, errors.New("Неверный логин")
	}

	accountInstance := &Account{
		CreatedAt: time.Now(),
		Login:     login,
		Password:  password,
		UpdatedAt: time.Now(),
		Url:       urlString,
	}

	if password == "" {
		accountInstance.generatePassword(10)
	}

	return accountInstance, nil
}

func (acc Account) Output() {
	color.Green(acc.Login)
	color.Green(acc.Password)
	color.Green(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	acc.Password = string(res)
}
