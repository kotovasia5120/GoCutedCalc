package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanian = [11]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func showError(errorMessage string) {
	err := fmt.Errorf(errorMessage)
	fmt.Println(err)
}

// Проверка на содержание
func isContain(el string, list []string) bool {
	for _, v := range list {
		if v == el {
			return true
		}
	}
	return false
}

// Проверка используемых чисел
func validateNum(num string) bool {
	nums := [20]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	return isContain(num, nums[:])
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

// Проверка операции
func validateOperation(op string) bool {
	operations := [4]string{"+", "-", "*", "/"}
	return isContain(op, operations[:])
}

// Проведение вычислений

func arabicCalc(x int, y int, oper string) int {
	var res int
	switch oper {
	case "+":
		res = x + y
	case "-":
		res = x - y

	case "*":
		res = x * y

	case "/":
		res = x / y
	}
	return res
}

func romanianCalc(x string, y string, oper string) string {
	x_conv := decode(x)
	y_conv := decode(y)
	result := arabicCalc(x_conv, y_conv, oper)
	if result <= 0 {
		return "Римские числа немогут быть отрицательны."
	}
	return ArabicToRoman(result)
}
func decode(el string) int {
	// convert romanian to arabic
	for i, v := range romanian {
		if v == el {
			return i
		}
	}
	return -1
}
func main() {
	//Ввод строки с пробелами
	fmt.Print("Введите математическую операцию (+, -, *, /), разделяя операнды и оператор пробелами: ")
	readerline := bufio.NewReader(os.Stdin)
	line, err := readerline.ReadString('\n')
	if err != nil {
		errors.New("Ошибка ввода: ")
	}
	parts := strings.Fields(line)
	fmt.Println(line)

	for {
		// Разбиваем строку на операнды и оператор
		if len(parts) != 3 || !validateOperation(parts[1]) {
			showError("Формат ввода не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			break
		} else if !validateNum(parts[0]) || !validateNum(parts[2]) {
			showError("Только арабские или римские числа от 1 до 10 (I - X)")
			break
		} else {
			// вызов вычислений и вывод
			num1, num1_err := strconv.Atoi(strings.TrimSpace(parts[0]))
			operation := strings.TrimSpace(parts[1])
			num2, num2_err := strconv.Atoi(strings.TrimSpace(parts[2]))

			if num2_err == nil && num1_err == nil {
				result := arabicCalc(num1, num2, operation)
				fmt.Println(result)
				break
			} else if num2_err != nil && num1_err != nil {
				result := romanianCalc(parts[0], parts[2], parts[1])
				if len(result) > 6 {
					showError(result)
					break
				} else {
					fmt.Println(result)
					break
				}
			} else {
				showError("Операция между римскими и арабскиими числами невозможна")
				break
			}
		}
	}

}
