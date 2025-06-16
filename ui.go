package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Главное меню
func MainMenu() {
	for {
		menuItems := []string{"Ввести строку для решения",
			"Закончить работу"}
		printMenu(menuItems)

		choice := getUserInput("Выберите нужное действие: ")
		index, err := strconv.Atoi(choice)

		if err != nil || index < 1 || index > len(menuItems) {
			fmt.Println("Невозможный выбор, попробуйте еще раз")
			continue
		}

		switch index {
		case 1:
			question := getUserInput("Введите математическую строку: ")
			result, err := Linesolver(question)
			if err == nil {
				fmt.Println("Ответ: ", result)
				fmt.Println("Возвращение в главное меню...")
			} else {
				fmt.Println("Встречена ошибка: ", err)
				fmt.Println("Возвращение в главное меню...")
			}

		case 2:
			fmt.Println("Завершение работы...")
			return
		}
	}
}

// Выводит функции меню на экран
func printMenu(items []string) {
	fmt.Println("\nМеню:")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

// Получение ввода от пользователя
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
