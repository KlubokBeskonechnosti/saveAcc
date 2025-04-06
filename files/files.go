package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("files.txt") // Чтение файла
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func WriteFile(content string, name string){
	file, err := os.Create(name) // Создаёт файл
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // Закрывает файл
	_, err = file.WriteString(content) // Записывает строку
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
