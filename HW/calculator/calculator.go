package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func IsDigit(elem string) bool {
	for _, letter := range "0123456789." {
		if elem == string(letter) {
			return true
		}
	}
	return false
}

func Calc(s string) (float64, error) {
	s = strings.ReplaceAll(s, " ", "")

	// Проверка на некорректные символы в конце строки
	if len(s) == 0 || strings.ContainsAny(s[len(s)-1:], "+-*/") {
		return 0, errors.New("expression is not valid")
	}

	var mix_s []interface{}
	var current_num string

	// Разбиваем строку на числа и операторы
	for i := 0; i < len(s); i++ {
		if IsDigit(string(s[i])) {
			current_num += string(s[i])
		} else {
			if current_num != "" {
				// Проверка на начало числа с нуля, исключая дробные числа
				if len(current_num) > 1 && current_num[0] == '0' && current_num[1] != '.' {
					return 0, errors.New("expression is not valid")
				}
				num, err := strconv.ParseFloat(current_num, 64)
				if err != nil {
					return 0, errors.New("ошибка преобразования числа: " + err.Error())
				}
				mix_s = append(mix_s, num)
				current_num = ""
			}
			// Проверка на последовательные операторы
			if i > 0 && strings.ContainsAny(string(s[i]), "+-*/") && strings.ContainsAny(string(s[i-1]), "+-*/") {
				return 0, errors.New("expression is not valid")
			}
			mix_s = append(mix_s, string(s[i]))
		}
	}

	// Если остались цифры в конце, добавляем их
	if current_num != "" {
		if len(current_num) > 1 && current_num[0] == '0' && current_num[1] != '.' {
			return 0, errors.New("expression is not valid")
		}
		num, err := strconv.ParseFloat(current_num, 64)
		if err != nil {
			return 0, errors.New("ошибка преобразования числа: " + err.Error())
		}
		mix_s = append(mix_s, num)
	}

	// Проходим по интерфейсу, заполняем стеки и выполняем мат. действия
	stek1 := make([]float64, 0)
	stek2 := make([]string, 0)

	for i := 0; i < len(mix_s); i++ {
		if num, ok := mix_s[i].(float64); ok {
			stek1 = append(stek1, num)
		} else if mix_s[i] == "(" {
			stek2 = append(stek2, mix_s[i].(string))
		} else if mix_s[i] == ")" {
			// Выполняем операции до открывающей скобки
			for len(stek2) > 0 && stek2[len(stek2)-1] != "(" {
				op := stek2[len(stek2)-1]
				stek2 = stek2[:len(stek2)-1]
				b := stek1[len(stek1)-1]
				stek1 = stek1[:len(stek1)-1]
				a := stek1[len(stek1)-1]
				stek1 = stek1[:len(stek1)-1]
				if op == "+" {
					stek1 = append(stek1, a+b)
				} else if op == "-" {
					stek1 = append(stek1, a-b)
				} else if op == "*" {
					stek1 = append(stek1, a*b)
				} else if op == "/" {
					if b == 0 {
						return 0, errors.New("expression is not valid")
					}
					stek1 = append(stek1, a/b)
				}
			}
			stek2 = stek2[:len(stek2)-1] // Убираем (
		} else {
			// Пока верхний оператор в стеке имеет такой же или более высокий приоритет
			for len(stek2) > 0 && (stek2[len(stek2)-1] == "*" || stek2[len(stek2)-1] == "/" || (stek2[len(stek2)-1] == "+" || stek2[len(stek2)-1] == "-") && (mix_s[i] == "+" || mix_s[i] == "-")) {
				op := stek2[len(stek2)-1]
				stek2 = stek2[:len(stek2)-1]
				b := stek1[len(stek1)-1]
				stek1 = stek1[:len(stek1)-1]
				a := stek1[len(stek1)-1]
				stek1 = stek1[:len(stek1)-1]
				if op == "+" {
					stek1 = append(stek1, a+b)
				} else if op == "-" {
					stek1 = append(stek1, a-b)
				} else if op == "*" {
					stek1 = append(stek1, a*b)
				} else if op == "/" {
					if b == 0 {
						return 0, errors.New("expression is not valid")
					}
					stek1 = append(stek1, a/b)
				}
			}
			// Добавляем текущий оператор в стек операторов
			stek2 = append(stek2, mix_s[i].(string))
		}
	}
	// Выполняем оставшиеся операции
	for len(stek2) > 0 {
		op := stek2[len(stek2)-1]
		stek2 = stek2[:len(stek2)-1]
		b := stek1[len(stek1)-1]
		stek1 = stek1[:len(stek1)-1]
		a := stek1[len(stek1)-1]
		stek1 = stek1[:len(stek1)-1]
		if op == "+" {
			stek1 = append(stek1, a+b)
		} else if op == "-" {
			stek1 = append(stek1, a-b)
		} else if op == "*" {
			stek1 = append(stek1, a*b)
		} else if op == "/" {
			if b == 0 {
				return 0, errors.New("expression is not valid")
			}
			stek1 = append(stek1, a/b)
		}
	}
	return stek1[0], nil
}

