package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// Структура токена
type Piece struct {
	Value string
	Class int
	Prior int
}

// Блок для error-handling
type ErrInvalidElement struct {
	Element string
	Pos     int
}

func (x ErrInvalidElement) Error() string {
	return fmt.Sprintf("неадекватный элемент или его положение: %s Позиция: %v", string(x.Element), x.Pos)
}

//

// Первичное разбивание строки на токены типа Piece
func Tear(xas string) (result []Piece, err error) {
	if len(xas) < 1 {
		result = append(result, Piece{"0", 0, 0})
		return result, errors.New("пустая входная строка")
	}

	x := strings.Split(xas, "")
	phase := true

	openers := 0
	closers := 0

	for i, t := 0, len(x); i < t; {
		c := x[i]
		var pr *Piece
		if len(result) != 0 {
			pr = &result[(len(result))-1]
		}

		// Пробел не несёт математического значения
		if c == " " {
			i++
			continue
		}

		// Поиск скобочек
		if c == "(" && phase {
			openers++
			// Добавлено умножение, чтобы строки вроде 1(2+3) могли корретно работать
			if !(pr == nil) && (pr.Class == 0 || pr.Class == -2) {
				result = append(result, Piece{"*", 3, 0})
			}
			result = append(result, Piece{"(", -1, 0})
			i++
			continue
		}
		if c == ")" && phase {
			closers++
			result = append(result, Piece{")", -2, 0})
			i++
			continue
		}
		//

		// Унарный минус
		if c == "-" && (pr == nil || !automlp(pr.Class)) && phase {
			result = append(result, Piece{"-", 10, 0})
			i++
			continue
		}
		// Нет смысла создавать токен для унарного плюса т.к. это равносильно умножению на 1
		if c == "+" && (pr == nil || !automlp(pr.Class)) && phase {
			i++
			continue
		}
		//

		// Базовая арифметика
		if c == "+" && phase {
			result = append(result, Piece{"+", 1, 0})
			i++
			continue
		}
		if c == "-" && phase {
			result = append(result, Piece{"-", 2, 0})
			i++
			continue
		}
		if c == "*" && phase {
			result = append(result, Piece{"*", 3, 0})
			i++
			continue
		}
		if c == "/" && phase {
			result = append(result, Piece{"/", 4, 0})
			i++
			continue
		}
		if c == "%" && phase {
			result = append(result, Piece{"%%", 5, 0})
			i++
			continue
		}
		if c == "^" && phase {
			result = append(result, Piece{"^", 6, 0})
			i++
			continue
		}
		// Конец базовой арифметики

		// Поиск числа
		if ispon(c) || (c == "-" && !phase) {
			// проверка на необходимость добавления умножения
			if !(pr == nil) && automlp(pr.Class) && phase {
				result = append(result, Piece{"*", 3, 0})
			}

			warn := 0
			var numStr string
			if c == "." {
				numStr = "0"
				warn++
			}
			numStr += string(c)

			// собираем оставшиеся цифры и точки
			g := 1
			for i+g < t && ispon(x[i+g]) {
				if x[i+g] == "." {
					warn++
				}
				numStr += string(x[i+g])
				g++
			}

			// Если было получено число с больше чем одной точкой, например, 100.01.02,
			// его нельзя явно преобразовать во float
			if warn > 1 {
				return result, ErrInvalidElement(ErrInvalidElement{numStr, i + 1})
			} else {
				result = append(result, Piece{numStr, 0, 0})
			}

			i += g
			continue
		}
		// конец чисел

		// Слова и константы
		if islett(c) && phase {
			// once again multiply
			if !(pr == nil) && automlp(pr.Class) {
				result = append(result, Piece{"*", 3, 0})
			}

			var letStr string

			letStr += string(c)

			// собираем оставшиеся буквы
			g := 1
			for i+g < t && islett(x[i+g]) {
				letStr += string(x[i+g])
				g++
			}
			wordslist := words(letStr)
			result = append(result, wordslist...)

			i += g
			continue
		}
		// конец букв

		// Разделитель, после которого следуют значения переменных
		if c == "|" {

			if openers != closers {
				for openers > closers {
					result = append(result, Piece{")", -2, 0})
					closers++
				}
				var xleb []Piece
				for openers < closers {
					xleb = append(xleb, Piece{"(", -1, 0})
					openers++
				}
				if len(xleb) > 0 {
					result = append(xleb, result...)
				}
			}

			result = append(result, Piece{"|", -10, 0})
			phase = false
			i++
			continue
		}
		// конец блок с разделителем

		return result, ErrInvalidElement(ErrInvalidElement{c, i + 1})
	}

	return result, nil
}

// "Is Part Of Number" (часть числа)
// Проверяет явялется ли входная строка точкой или цифрой
func ispon(x string) bool {
	return x == "." || (x >= "0" && x <= "9")
}

// "Is letter" (буква)
// Проверяет является ли входная строка буквой
func islett(x string) bool {
	return (x >= "a" && x <= "z") || (x >= "A" && x <= "Z") || x == "π"
}

// Проверяет нужен ли автоматический знак умножения
func automlp(x int) bool {
	return x == 0 || x == -2 || x == -3
}

// Проверяет нужен ли автоматический знак умножения (для других случаев)
func betamlp(x int) bool {
	return x == 0 || x == -1 || x == -3
}

// Разбивает на соответствюущие токены строку из букв
func words(x string) (wordslist []Piece) {
	runes := []rune(x)
	for len(runes) > 0 {
		n := len(runes)

		// тригонометрия
		if n >= 5 && string(runes[n-5:n]) == "acosh" {
			wordslist = append(wordslist, Piece{"acosh", 11, 0})
			runes = runes[:len(runes)-5]
			continue
		}
		if n >= 5 && string(runes[n-5:n]) == "asinh" {
			wordslist = append(wordslist, Piece{"asinh", 12, 0})
			runes = runes[:len(runes)-5]
			continue
		}
		if n >= 5 && string(runes[n-5:n]) == "atanh" {
			wordslist = append(wordslist, Piece{"atanh", 13, 0})
			runes = runes[:len(runes)-5]
			continue
		}

		if n >= 4 && string(runes[n-4:n]) == "acos" {
			wordslist = append(wordslist, Piece{"acos", 14, 0})
			runes = runes[:len(runes)-4]
			continue
		}
		if n >= 4 && string(runes[n-4:n]) == "asin" {
			wordslist = append(wordslist, Piece{"asin", 15, 0})
			runes = runes[:len(runes)-4]
			continue
		}
		if n >= 4 && string(runes[n-4:n]) == "atan" {
			wordslist = append(wordslist, Piece{"atan", 16, 0})
			runes = runes[:len(runes)-4]
			continue
		}

		if n >= 3 && string(runes[n-3:n]) == "cos" {
			wordslist = append(wordslist, Piece{"cos", 17, 0})
			runes = runes[:len(runes)-3]
			continue
		}
		if n >= 3 && string(runes[n-3:n]) == "sin" {
			wordslist = append(wordslist, Piece{"sin", 18, 0})
			runes = runes[:len(runes)-3]
			continue
		}
		if n >= 3 && string(runes[n-3:n]) == "tan" {
			wordslist = append(wordslist, Piece{"tan", 19, 0})
			runes = runes[:len(runes)-3]
			continue
		}
		// конец тригонометрии

		if n >= 2 && string(runes[n-2:n]) == "ln" {
			wordslist = append(wordslist, Piece{"ln", 20, 0})
			runes = runes[:len(runes)-2]
			continue
		}

		// константы
		if n >= 2 && string(runes[n-2:n]) == "pi" {
			wordslist = append(wordslist, Piece{"3.141593", 0, 0})
			runes = runes[:len(runes)-2]
			continue
		}
		if n >= 1 && string(runes[n-1:n]) == "π" {
			wordslist = append(wordslist, Piece{"3.141593", 0, 0})
			runes = runes[:len(runes)-1]
			continue
		}
		if n >= 1 && string(runes[n-1:n]) == "e" {
			wordslist = append(wordslist, Piece{"2.7182818", 0, 0})
			runes = runes[:len(runes)-1]
			continue
		}

		// конец констант

		wordslist = append(wordslist, Piece{string(runes[n-1]), -3, 0})
		runes = runes[:len(runes)-1]
	}
	slices.Reverse(wordslist)

	// Добавляем умножения если обнаружили идущие подряд переменные
	if len(wordslist) > 1 {
		i := 0
		for i+1 < len(wordslist) {
			a := wordslist[i]
			b := wordslist[i+1]
			if automlp(a.Class) && betamlp(b.Class) {
				wordslist = slices.Insert(wordslist, i+1, Piece{"*", 3, 0})
				i += 2
			} else {
				i++
			}
		}
	}
	return wordslist
}

// Получает на вход массив токенов (структур типа Piece), предположительно содержащий разделитель  " | "
// При необходимости обрабатывает массив, избавляясь от разделителя и заменяя
// переменные на необходимые значения.
// Если переменные не используются, то просто возвращает исходный массив
func Zandatsu(x []Piece) []Piece {
	var rights []Piece
	var result []Piece
	var lefts []string

	for i := range x {
		current := &x[i]
		if current.Class == -3 {
			if !slices.Contains(lefts, current.Value) {
				lefts = append(lefts, current.Value)
			}
		}
		if current.Class == -10 {
			rights = append(rights, x[(i+1):]...)
			result = append(result, x[:i]...)
			break
		}
	}
	if len(rights) == 0 {
		return x
	}

	for i := range result {
		current := &result[i]
		if current.Class == -3 {
			if g := slices.Index(lefts, current.Value); g != -1 {
				current.Value = rights[g].Value
				current.Class = 0
			}
		}
	}
	return result
}
