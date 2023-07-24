package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman, arabic bool

// Проверка используемых чисел
func RorA(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		if arabic == true {
			fmt.Println("Не используйте римские и арабские числа вместе", err)
			os.Exit(1)
		}
		roman = true
		switch in {
		case "I":
			return 1
		case "II":
			return 2
		case "III":
			return 3
		case "IV":
			return 4
		case "V":
			return 5
		case "VI":
			return 6
		case "VII":
			return 7
		case "VIII":
			return 8
		case "IX":
			return 9
		case "X":
			return 10
		default:
			fmt.Println("Неизвестное число")
			os.Exit(1)
		}
	}
	if out > 10 || out < 1 {
		fmt.Println("Вводите числа от 1 до 10")
		os.Exit(1)
	}
	if roman == true {
		fmt.Println("Не используйте римские и арабские числа вместе", err)
		os.Exit(1)
	}
	arabic = true
	return out
}

// Преобразования арабских чисел в римские
func ArabicToRoman(result int) string {
	switch {
	case result < 0:
		return "в римской системе нет отрицательных чисел."
	case result > 100:
		return "Результат больше 100, не подходит под условия задачи"
	case result == 0:
		return ""
	case result < 4:
		return "I" + ArabicToRoman(result-1)
	case result == 4:
		return "IV"
	case result < 9:
		return "V" + ArabicToRoman(result-5)
	case result == 9:
		return "IX"
	case result < 40:
		return "X" + ArabicToRoman(result-10)
	case result < 50:
		return "XL" + ArabicToRoman(result-40)
	case result < 90:
		return "L" + ArabicToRoman(result-50)
	case result < 100:
		return "XC" + ArabicToRoman(result-90)
	case result == 100:
		return "C"
	}
	return ""
}

// Проведение вычислений
func NumOpertion(a int, operand string, b int) string {
	var result int
	switch operand {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Неизвестная операция")
		os.Exit(1)
	}
	if roman == true && result < 1 {
		fmt.Println("Римские цифры не могут быть с отрицательным значением")
		os.Exit(1)
	}
	if roman == true {
		return ArabicToRoman(result)
	}
	return strconv.Itoa(result)
}

func main() {
	//Ввод строки с пробелами
	fmt.Print("Введите математическую операцию (+, -, *, /), разделяя операнды и оператор пробелами: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		os.Exit(1)
	}
	input = strings.TrimSpace(input)
	fmt.Println(input)

	// Разбиваем строку на операнды и оператор
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат операции")
		return
	}
	// вызов вычислений и вывод
	fmt.Printf("%v %v %v = %v", parts[0], parts[1], parts[2], NumOpertion(RorA(parts[0]), parts[1], RorA(parts[2])))
}
