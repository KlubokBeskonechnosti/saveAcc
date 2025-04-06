package main

import (
	"demo/password/files"
	"demo/password/account"
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Приложение паролей")
	vault := account.NewVault()

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			CreateAccount(vault)
		case 2:
			FindAccount(vault)
		case 3:
			DeleteAccount()
		default:
			break Menu
		}
	}

}

func promptData(prompt string) string {
    fmt.Print(prompt + ": ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
}

func getMenu() int {
    var variant int
    for {
        fmt.Println("Выберите вариант:")
        fmt.Println("1. Создать аккаунт")
        fmt.Println("2. Найти аккаунт")
        fmt.Println("3. Удалить аккаунт")
        fmt.Println("4. Выход")
        
        _, err := fmt.Scan(&variant)
        if err != nil {
            fmt.Println("Ошибка ввода, попробуйте еще раз")
            // Очищаем буфер ввода
            var discard string
            fmt.Scanln(&discard)
            continue
        }
        // Очищаем буфер ввода после чтения числа
        var discard string
        fmt.Scanln(&discard)
        
        if variant >= 1 && variant <= 4 {
            return variant
        }
        fmt.Println("Неверный вариант, попробуйте еще раз")
    }
}

func CreateAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount) // Добавляем аккаунт
	
	data, err := vault.ToBytes() // Конвернтируем в JSON
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(data, "data.json") // Сохраняем в файл
}

func FindAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		fmt.Println("Аккаунт не найден")
	}
	for _, account := range accounts{
		account.Output()
	}
}

func DeleteAccount() {

}
