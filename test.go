package main

import "fmt"

type testCase struct {
	Answer  float64
	Problem string
}

// Набор тестов. Несут две цели:
// 1. Проверка на то, что в системе, где запущена программа,
// подсчёт будет исполняться корректно
// 2. Регрессивное тестирование, чтобы обнаружить появление ошибок
// после внесения изменений в код
func Boottest() {
	var all []testCase
	success := 0

	// целые цисла работают корректно
	all = append(all, testCase{156.0, "156"})
	// float работает корректно
	all = append(all, testCase{33.4, "10.5 + 22.9"})
	all = append(all, testCase{-36.843, ".157 - 37"})
	all = append(all, testCase{2710.92, "11890 * .228"})
	all = append(all, testCase{40.0, "40."})
	// базовая арифметика
	all = append(all, testCase{5.0, "2 + 3"})
	all = append(all, testCase{0.0, "1 - 1"})
	all = append(all, testCase{12.0, "3 * 4"})
	all = append(all, testCase{4.0, "12 / 3"})
	all = append(all, testCase{2.0, "12 % 10"})
	// степени
	all = append(all, testCase{8.0, "2 ^ 3"})
	all = append(all, testCase{1.0, "2 ^ 0"})
	all = append(all, testCase{0.5, "2 ^ -1"})
	all = append(all, testCase{0.25, "2 ^ -2"})
	// константы
	all = append(all, testCase{3.141593, "π"})
	all = append(all, testCase{3.141593, "pi"})
	all = append(all, testCase{2.7182818, "e"})
	all = append(all, testCase{9.869607, "ππ"})
	all = append(all, testCase{6.283186, "2π"})
	all = append(all, testCase{6.283186, "π2"})

	// Унарные функции
	all = append(all, testCase{-10.0, "-10"})
	// тригонометрия
	all = append(all, testCase{1.316958, "acosh(2)"})
	all = append(all, testCase{1.443635, "asinh(2)"})
	all = append(all, testCase{0.549306, "atanh(0.5)"})
	all = append(all, testCase{1.570796, "acos(0)"})
	all = append(all, testCase{1.570796, "asin(1)"})
	all = append(all, testCase{0.785398, "atan(1)"})
	all = append(all, testCase{1.00, "cos(0)"})
	all = append(all, testCase{1.00, "sin(π/2)"})
	all = append(all, testCase{1.00, "tan(π/4)"})
	all = append(all, testCase{1.00, "ln(e)"})

	// Тест на сложные выражения
	all = append(all, testCase{-2.0, "-1 + -1"})
	all = append(all, testCase{-586755.0, "(1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) + (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) + (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) - (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) - (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) - (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) * (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4))) + (1 + 2 * (3 + 4 * (5 + 6 * (7 + 8)))) / (9 - 8 * (7 - 6 * (5 - 4)))"})
	all = append(all, testCase{0.540302, "coscos(0)"})
	// Обработка очень длинных чисел
	all = append(all, testCase{1.0, "(100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000*0)^(2-2)"})

	// Тест на корректную работу переменных
	all = append(all, testCase{-9.0, "x + 1 | -10"})
	all = append(all, testCase{-20.0, "x + x | -10"})
	all = append(all, testCase{12.0, "(x + x) * y | 2 3"})

	// Тест строк приводивших к "панике" (указали на GitHub)
	all = append(all, testCase{-10.0, " -10"})
	all = append(all, testCase{-1.0, "-(+(-3 + +(+4)))"})
	all = append(all, testCase{3.0, "++3"})

	for i := 0; i < len(all); i++ {
		cur := &all[i]
		result, _ := Linesolver(cur.Problem)
		if cur.Answer == result {
			success++
		} else {
			str := fmt.Sprintf("Баг в: %s", cur.Problem)
			fmt.Println(str)
			str = fmt.Sprintf("ожидаемый ответ: %v реальный ответ: %v", cur.Answer, result)
			fmt.Println(str)
		}
	}

	str := fmt.Sprintf("%v из %v тестов пройдены успешно", success, len(all))
	fmt.Println(str)

	if success == len(all) {
		fmt.Println("Программа готова к использованию.")
	} else {
		fmt.Println("Внимание!: могут возникнуть ошибки")
	}
}
