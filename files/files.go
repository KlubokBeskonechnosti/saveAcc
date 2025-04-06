package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name) // Чтение файла
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string){
	file, err := os.Create(name) // Создаёт файл
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // Закрывает файл
	_, err = file.Write(content) // Записывает строку
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
