package main

import (
	"fmt"
)

func Linesolver(x string) (answer float64, err error) {
	buff, err := Tear(x)
	if err != nil {
		return 1, err
	}
	buff = Zandatsu(buff)
	answer, err = Calculate(buff)
	if err != nil {
		return 1, err
	}
	return answer, err
}

func main() {
	fmt.Println("LineSolver загружен")
	fmt.Println("Запуск автоматических тестов...")
	Boottest()

	MainMenu()
}
