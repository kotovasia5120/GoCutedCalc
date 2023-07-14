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
func ArabicToRoman(arabicnumeral int) string {
	arabicmap := map[int]string{
		500: "D",
		100: "C",
		50:  "L",
		10:  "X",
		5:   "V",
		1:   "I",
	}
	result := ""
	for num, symbol := range arabicmap {
		for arabicnumeral >= num {
			result += symbol
			arabicnumeral -= num
		}
	}
	return result
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
	fmt.Println(NumOpertion(RorA(parts[0]), parts[1], RorA(parts[2])))
	fmt.Printf("%v %v %v = %v", parts[0], parts[1], parts[2], NumOpertion(RorA(parts[0]), parts[1], RorA(parts[2])))
}
