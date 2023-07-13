package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	fmt.Print("Введите математическую операцию (+, -, *, /), разделяя операнды и оператор пробелами: ")

	fmt.Println(input)

	// Разбиваем строку на операнды и оператор
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат операции")
		return
	}

	// Преобразуем операнды в числа
	num1, err1 := strconv.ParseFloat(parts[0], 64)
	num2, err2 := strconv.ParseFloat(parts[2], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Неверный формат операндов")
		return
	}

	// Выполняем операцию
	var result float64
	switch parts[1] {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Неверный оператор")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f", num1, parts[1], num2, result)
}
