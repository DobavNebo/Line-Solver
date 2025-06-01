package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Piece struct {
	Value string
	Class int
	Prior int
}

func Linesolver(x string) float64 {
	ans := Tear(x)
	ans = Zandatsu(ans)
	return Calculate(ans)
}

func main() {
	fmt.Println("LineSolver is loaded")
	fmt.Println("Tests are getting started...")
	Boottest()

	menuItems := []string{"Help. How to use programm properly",
		"Insert math line",
		"Exit"}

	for {
		printMenu(menuItems)

		choice := getUserInput("Enter needed option: ")
		index, err := strconv.Atoi(choice)

		if err != nil || index < 1 || index > len(menuItems) {
			fmt.Println("Invalid option. Please try again.")
			continue
		}

		switch index {
		case 1:
			fmt.Println(" ")
			fmt.Println("===HELP===")
			fmt.Println("Program will the math line you insert.")
			fmt.Println("To insert the line choose option 2.")
			fmt.Println("If you want to solve something like x*x*x with x = 3 and y = 4")
			fmt.Println("You should write it like x*x*y | 3 4")
			fmt.Println("Variables are treated in the same order they appear left-to-right")
			fmt.Println("==========")
		case 2:
			question := getUserInput("Enter math line: ")
			fmt.Println("Answer: ", Linesolver(question))
		case 3:
			fmt.Println("Exiting...")
			return
		}
	}
}

func printMenu(items []string) {
	fmt.Println("\nMenu:")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
