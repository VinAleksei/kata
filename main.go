package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	// Запускаем цикл ввода данных от пользователя, прерываем выполнение командой "exit",
	// обязательно обрабатываем ввод с учетом красной строки
	for {
		fmt.Println("Тестовый калькулятор v 0.1:")
		fmt.Println("Введите арифметическое выражение, используйте операнды от 1 до 10 или римские числа (I-X).")
		in, _ := reader.ReadString('\n')
		// обработка красной строки
		in = strings.TrimSpace(in)
		// выход из цикла по команде "exit"
		if in == "exit" {
			break
		}
		operator, operant1int, operant2int, err := InParts(in)
		if err != nil {
			fmt.Println("", err)
			continue
		}
		result, err := calculate(operator, operant1int, operant2int)
		if err != nil {
			fmt.Println("", err)
			continue
		}

		fmt.Println("Результат:", result)
	}
}

// isRoman проверяет, является ли строка римским числом.
func isRoman(s string) bool {
	romanSymbols := "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(romanSymbols, char) {
			return false
		}
	}
	return true
}

// InParts функция на проверку условий ввода данных и разбиение строки на части,
// для проведения арифметических выражений.
func InParts(in string) (string, int, int, error) {
	parts := strings.Split(in, " ")
	if len(parts) > 3 {
		return "", 0, 0, fmt.Errorf("error: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	} else if len(parts) < 3 {
		return "", 0, 0, fmt.Errorf("error: строка не является математической операцией")
	}

	operator := strings.TrimSpace(parts[1])
	operant1 := strings.TrimSpace(parts[0])
	operant2 := strings.TrimSpace(parts[2])

	var operant1int, operant2int int
	var err error

	if isRoman(operant1) {
		operant1int, err = romanToArabic(operant1)
		if err != nil {
			return "", 0, 0, fmt.Errorf("error: %v", err)
		}
	} else {
		operant1int, err = strconv.Atoi(operant1)
		if err != nil {
			return "", 0, 0, fmt.Errorf("error: проверьте первый операнд, используйте только целые числа")
		}
	}

	if isRoman(operant2) {
		operant2int, err = romanToArabic(operant2)
		if err != nil {
			return "", 0, 0, fmt.Errorf("error: %v", err)
		}
	} else {
		operant2int, err = strconv.Atoi(operant2)
		if err != nil {
			return "", 0, 0, fmt.Errorf("error: проверьте второй операнд, используйте только целые числа")
		}
	}

	if (1 <= operant1int && operant1int <= 10) && (1 <= operant2int && operant2int <= 10) {
	} else {
		return "", 0, 0, fmt.Errorf("error: операнды должны быть в пределах от 1 до 10")
	}
	return operator, operant1int, operant2int, nil
}

// romanToArabic конвертирует римское число в арабское число.
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[string(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result, nil
}

// calculate функция математических операций
func calculate(operator string, operant1int, operant2int int) (int, error) {
	switch operator {
	case "+":
		return operant1int + operant2int, nil
	case "-":
		return operant1int - operant2int, nil
	case "*":
		return operant1int * operant2int, nil
	case "/":
		if operant2int == 0 {
			return 0, fmt.Errorf("error: деление на ноль")
		}
		return operant1int / operant2int, nil
	default:
		return 0, fmt.Errorf("error: недопустимый оператор")
	}
}
