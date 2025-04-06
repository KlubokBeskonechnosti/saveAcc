package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	// "google.golang.org/api/vault/v1"
)

type Vault struct { // Хранилище аккаунтов, можно посмотреть когда была последняя запись
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault { // Создаёт новый обьект Vault
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Printf("Не удалось разобрать файл data.json: %v", err)
		return &Vault{Accounts: []Account{}, UpdatedAt: time.Now()} // Обработка ошибок JSOM
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc) // Добавляет аккаунт
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось разобрать файл dats.json")
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault) // Конвертирует в JSON
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccountByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}
