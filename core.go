package main

import (
	"fmt"
)

func Linesolver(x string) float64 {
	ans := Tear(x)
	ans = Zandatsu(ans)
	return Calculate(ans)
}

func main() {
	fmt.Println("LineSolver загружен")
	fmt.Println("Запуск автоматических тестов...")
	Boottest()

	MainMenu()
}
