package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Введите математическую операцию (+, -, *, /), разделяя операнды и оператор пробелами: ")
	fmt.Scanln(&input)

	// Разбиваем строку на операнды и оператор
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат операции")
		return
	}
}
