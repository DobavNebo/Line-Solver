package main

import (
	"fmt"
	"slices"
	"strings"
)

func Tear(xas string) []Piece {
	var result []Piece
	if len(xas) < 1 {
		result = append(result, Piece{"0", 0, 0})
		return result
	}

	x := strings.Split(xas, "")
	phase := true

	for i, t := 0, len(x); i < t; {
		c := x[i]
		var pr *Piece
		if len(result) != 0 {
			pr = &result[(len(result))-1]
		}

		// space is not an element of equation
		if c == " " {
			i++
			continue
		}

		// finding these pals ( )
		if c == "(" && phase {
			//added multiply there so 1(2+3) is valid
			if !(pr == nil) && (pr.Class == 0 || pr.Class == -2) {
				result = append(result, Piece{"*", 3, 0})
			}
			result = append(result, Piece{"(", -1, 0})
			i++
			continue
		}
		if c == ")" && phase {
			result = append(result, Piece{")", -2, 0})
			i++
			continue
		}
		// end of ( ) part

		// is minus guy unar?
		if c == "-" && (pr == nil || !automlp(pr.Class)) && phase {
			result = append(result, Piece{"-", 10, 0})
			i++
			continue
		}
		// we surely need no unar plus
		if c == "+" && (pr == nil || !automlp(pr.Class)) && phase {
			i++
			continue
		}
		// end of unars minus

		// basic arifmethics
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
		// end of basic arifmethics part

		// is a number?
		if ispon(c) || (c == "-" && !phase) {
			// once again multiply
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

			// gather the rest of digits and periods
			g := 1
			for i+g < t && ispon(x[i+g]) {
				if x[i+g] == "." {
					warn++
				}
				numStr += string(x[i+g])
				g++
			}
			result = append(result, Piece{numStr, 0, 0})
			if warn > 1 {
				fmt.Println("Warning! Incorrect number: " + numStr + "\n It will be perceived as 0 (zero)")
			}

			i += g
			continue
		}
		// end of number part

		// consts, letters and other
		if islett(c) && phase {
			// once again multiply
			if !(pr == nil) && automlp(pr.Class) {
				result = append(result, Piece{"*", 3, 0})
			}

			var letStr string

			letStr += string(c)

			// gather the rest of letters
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
		// end of letters

		// dark magic territory (messing with mathematical variables)
		if c == "|" {
			result = append(result, Piece{"|", -10, 0})
			phase = false
			i++
			continue
		}
		// end of dark magic

		fmt.Println("WARNING! Inapropriate element or its positioning:", c)
		fmt.Println("Position: ", i)
		fmt.Println("It will be ignored")

		i += 1
	}
	return result
}

func ispon(x string) bool {
	// Is Part Of Number
	// Checks if input byte-typed rune is either digit or period
	return x == "." || (x >= "0" && x <= "9")
}

func islett(x string) bool {
	// Is letter
	return (x >= "a" && x <= "z") || (x >= "A" && x <= "Z") || x == "π"
}

func automlp(x int) bool {
	// Checks if additional multiplier needed
	return x == 0 || x == -2 || x == -3
}

func betamlp(x int) bool {
	// Checks if additional multiplier needed
	return x == 0 || x == -1 || x == -3
}

func words(x string) (wordslist []Piece) {
	runes := []rune(x)
	for len(runes) > 0 {
		n := len(runes)

		// trigonometry
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
		// end of trigonometry

		if n >= 2 && string(runes[n-2:n]) == "ln" {
			wordslist = append(wordslist, Piece{"ln", 20, 0})
			runes = runes[:len(runes)-2]
			continue
		}

		//constants
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

		//end of constants

		wordslist = append(wordslist, Piece{string(runes[n-1]), -3, 0})
		runes = runes[:len(runes)-1]
	}
	slices.Reverse(wordslist)
	// safety net for messing with multiplying
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
