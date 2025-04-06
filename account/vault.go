package account

import (
	"encoding/json"
	"time"
)
type Vault struct { // Хранилище аккаунтов, можно посмотреть когда была последняя запись
	Accounts []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault { // Создаёт новый обьект Vault
	return &Vault{
		Accounts: []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc) // Добавляет аккаунт
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault) // Конвертирует в JSON
	if err != nil {
		return nil, err
	}
	return file, nil
}