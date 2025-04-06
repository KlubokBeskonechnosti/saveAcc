package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	// "bufio"
	// "os"
	// "strings"
)

type Account struct {
	Login     string    `json:"login"` // сериализуется в JSON
	Password  string    `json:"password"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	fmt.Println(acc.Login, acc.Password, acc.URL)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res { // Заполняет случайными символами
		res[i] = lettersRune[rand.IntN(len(lettersRune))]
	}
	acc.Password = string(res) // Сохраняет пароль
}

func NewAccountWithTimeStamp(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}

	newAcc := &Account{ // Создает аккаунт
		Login:     login,
		Password:  password,
		URL:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" { // Если пароль не введен
		newAcc.generatePassword(12) // Генерирует случайный
	}

	return newAcc, nil
}

var lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
