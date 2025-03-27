package account

import (
	"errors"
	"fmt"
	"net/url"
	"time"
	"math/rand/v2"
	// "bufio"
	// "os"
	// "strings"
)

type Account struct {
    Login     string    `json:"login"`
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
    for i := range res {
        res[i] = lettersRune[rand.IntN(len(lettersRune))]
    }
    acc.Password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*Account, error) {
    if login == "" {
        return nil, errors.New("INVALID_LOGIN")
    }
    _, err := url.ParseRequestURI(urlString)
    if err != nil {
        return nil, errors.New("INVALID URL")
    }
    
    newAcc := &Account{
        Login:     login,
        Password:  password,
        URL:       urlString,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    
    if password == "" {
        newAcc.generatePassword(12)
    }
    
    return newAcc, nil
}

var lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")


