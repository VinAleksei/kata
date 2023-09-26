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
		fmt.Println("Тестовый калькулятор v 0.1: для выхода используйте 'exit' ")
		fmt.Println("Введите арифметическое выражение, используйте операнды от 1 до 10 или римские числа (I-X).")
		in, _ := reader.ReadString('\n')
		// обработка красной строки
		in = strings.TrimSpace(in)
		// выход из цикла по команде "exit"
		if in == "exit" {
			break
		}
		operator, operant1str, operant1int, operant2int, err := InParts(in)
		if err != nil {
			fmt.Println("", err)
			continue
		}
		result, err := calculate(operator, operant1int, operant2int)
		if err != nil {
			fmt.Println("", err)
			continue
		}

		printResult(result, isRoman(operant1str))

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
func InParts(in string) (string, string, int, int, error) {
	parts := strings.Split(in, " ")
	if len(parts) > 3 {
		return "", "", 0, 0, fmt.Errorf("error: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	} else if len(parts) < 3 {
		return "", "", 0, 0, fmt.Errorf("error: строка не является математической операцией")
	}

	operator := strings.TrimSpace(parts[1])
	operant1 := strings.TrimSpace(parts[0])
	operant2 := strings.TrimSpace(parts[2])

	var operant1str string
	var operant1int, operant2int int
	var err error

	if isRoman(operant1) != isRoman(operant2) {
		return "", "", 0, 0, fmt.Errorf("error: калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	if isRoman(operant1) {
		operant1str = operant1
		operant1int, err = romanToArabic(operant1)
		if err != nil {
			return "", "", 0, 0, fmt.Errorf("error: %v", err)
		}
	} else {
		operant1str = operant1
		operant1int, err = strconv.Atoi(operant1)
		if err != nil {
			return "", "", 0, 0, fmt.Errorf("error: проверьте первый операнд, используйте только целые числа")
		}
	}

	if isRoman(operant2) {
		operant2int, err = romanToArabic(operant2)
		if err != nil {
			return "", "", 0, 0, fmt.Errorf("error: %v", err)
		}
	} else {
		operant2int, err = strconv.Atoi(operant2)
		if err != nil {
			return "", "", 0, 0, fmt.Errorf("error: проверьте второй операнд, используйте только целые числа")
		}
	}

	if (1 <= operant1int && operant1int <= 10) && (1 <= operant2int && operant2int <= 10) {
	} else {
		return "", "", 0, 0, fmt.Errorf("error: операнды должны быть в пределах от 1 до 10")
	}
	return operator, operant1str, operant1int, operant2int, nil
}

// romanToArabic конвертирует римское число в арабское число.
func romanToArabic(roman string) (int, error) {
	var RomanNumerals = map[string]int{
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
		value := RomanNumerals[string(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 {
		return "", fmt.Errorf("error: в римской системе нет отрицательных чисел и нуля")
	} else if arabic > 100 {
		return "", fmt.Errorf("error: в римской системе поддерживаются только числа от 1 до 100")
	}
	var romanNumerals = map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}

	var roman string
	if roman, ok := romanNumerals[arabic]; ok {
		return roman, nil
	}

	return roman, nil
}

func printResult(result int, isRoman bool) {
	fmt.Print("Результат: ")
	if isRoman {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("error: ошибка конвертации", err)
			return // Вернуться без вывода результата в случае ошибки конвертации
		}
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
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
