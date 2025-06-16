package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
)

// Блок для error-handling
type ErrUnpairedBracket struct {
	Pos int
	Bul bool
}

func (x ErrUnpairedBracket) Error() string {
	if x.Bul {
		return fmt.Sprintf("скобка ) не имеет парной открывающей Позиция: %v", x.Pos)
	} else {
		return fmt.Sprintf("скобка ( не имеет парной закрывающей Позиция: %v", x.Pos)
	}
}

type ErrImposibleOperand string

func (x ErrImposibleOperand) Error() string {
	return fmt.Sprintf("невозможное действие: %s", string(x))
}

//

// Головная функция. Сводит все промежуточные функции в этом файле
func Calculate(x []Piece) (ans float64, err error) {
	x, err = escapeParents(x)
	if err != nil {
		return 1, err
	}
	x, y := priorities(x)
	ans, err = solve(x, y)
	if err != nil {
		return 1, err
	}
	return ans, nil
}

// Избавление от скобок
func escapeParents(x []Piece) (res []Piece, err error) {
	bump := len(x)

	for i := 0; i < len(x); i++ {
		cur := x[i]
		if cur.Class == -1 {
			for g := i; g < len(x); g++ {
				c := &x[g]
				if c.Class == -2 {
					x = append(x[:g], x[(g+1):]...)
					x = append(x[:i], x[(i+1):]...)
					break
				}
				if c.Class > 0 {
					c.Prior += 4 * bump
				}
				if i == len(x) {
					return x, ErrUnpairedBracket(ErrUnpairedBracket{i, false})
				}
			}
			i--
		}
		if cur.Class == -2 {
			return x, ErrUnpairedBracket(ErrUnpairedBracket{i + 1, true})
		}
	}

	return x, nil
}

// Проставляет операторам приоритеты
func priorities(x []Piece) ([]Piece, []int) {
	leng := len(x)
	var priors []int

	g := 0
	for i := 0; i < leng; i++ {
		c := &x[i]
		if isunar(c.Class) {
			c.Prior += g + 1 + (3 * leng)
			g++
			priors = append(priors, c.Prior)
		}
	}
	g = 0
	for i := 0; i < leng; i++ {
		c := &x[i]
		if c.Class == 6 {
			c.Prior += g + 1 + (2 * leng)
			g++
			priors = append(priors, c.Prior)
		}
	}
	g = 0
	for i := leng - 1; i >= 0; i-- {
		c := &x[i]
		if istwo(c.Class) {
			c.Prior += g + 1 + leng
			g++
			priors = append(priors, c.Prior)
		}
	}
	g = 0
	for i := leng - 1; i >= 0; i-- {
		c := &x[i]
		if isone(c.Class) {
			c.Prior += g + 1
			g++
			priors = append(priors, c.Prior)
		}
	}
	slices.Sort(priors)
	slices.Reverse(priors)
	return x, priors
}

// Принадлежность к функциям 1-го (снизу) приоритета
func isone(a int) bool {
	return a == 1 || a == 2
}

// Принадлежность к функциям 2-го приоритета
func istwo(a int) bool {
	return a == 3 || a == 4 || a == 5
}

// Принадлежность к унарным функциям (они же функции 3-го и наибольшего приоритета)
func isunar(a int) bool {
	return a >= 10 && a <= 20
}

// Получение конечного ответа из массивая токенов и списка приоритетов операторов
func solve(x []Piece, priors []int) (end float64, err error) {
	for _, g := range priors {
		for i := 0; i < len(x); i++ {
			c := &x[i]
			if c.Prior == g {
				switch c.Class {
				case 1:
					//+
					err = safecheck(i, x, "+", 2)
					if err != nil {
						return 1, err
					}

					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value, 64)
					b, _ := strconv.ParseFloat(nxt.Value, 64)
					pr.Value = fmt.Sprintf("%f", a+b)
					x = append(x[:i], x[(i+2):]...)
				case 2:
					//-
					err = safecheck(i, x, "-", 2)
					if err != nil {
						return 1, err
					}

					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value, 64)
					b, _ := strconv.ParseFloat(nxt.Value, 64)
					pr.Value = fmt.Sprintf("%f", a-b)
					x = append(x[:i], x[(i+2):]...)
				case 3:
					//*
					err = safecheck(i, x, "*", 2)
					if err != nil {
						return 1, err
					}

					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value, 64)
					b, _ := strconv.ParseFloat(nxt.Value, 64)
					pr.Value = fmt.Sprintf("%f", a*b)
					x = append(x[:i], x[(i+2):]...)
				case 4:
					// /
					err = safecheck(i, x, "/", 2)
					if err != nil {
						return 1, err
					}

					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value, 64)
					b, _ := strconv.ParseFloat(nxt.Value, 64)
					if b == 0 {
						return 1, errors.New("деление на 0")
					}
					pr.Value = fmt.Sprintf("%f", a/b)
					x = append(x[:i], x[(i+2):]...)
				case 5:
					// %
					err = safecheck(i, x, "%", 2)
					if err != nil {
						return 1, err
					}

					pr := &x[i-1]
					nxt := &x[i+1]
					pro, _ := strconv.ParseFloat(pr.Value, 64)
					a := int(pro)
					nxto, _ := strconv.ParseFloat(nxt.Value, 64)
					b := int(nxto)
					if b == 0 {
						return 1, errors.New("остаток от деления на 0")
					}
					pr.Value = fmt.Sprintf("%f", float64(a%b))
					x = append(x[:i], x[(i+2):]...)
				case 6:
					// ^
					err = safecheck(i, x, "^", 2)
					if err != nil {
						return 1, err
					}

					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value, 64)
					b, _ := strconv.ParseFloat(nxt.Value, 64)
					if a == 0 {
						if b <= 0 {
							erstr := pr.Value + " ^ " + nxt.Value
							return 1, ErrImposibleOperand(ErrImposibleOperand(erstr))
						}
						if b == 0 {
							pr.Value = "1"
							x = append(x[:i], x[(i+2):]...)
							continue
						}
						if b > 0 {
							pr.Value = fmt.Sprintf("%f", math.Pow(a, b))
							x = append(x[:i], x[(i+2):]...)
							continue
						}
					}
					pr.Value = fmt.Sprintf("%f", math.Pow(a, b))
					x = append(x[:i], x[(i+2):]...)
				case 10:
					// - unar
					err = safecheck(i, x, "-", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", (-1.0)*a)
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 11:
					// acosh
					err = safecheck(i, x, "acosh", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Acosh(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 12:
					// asinh
					err = safecheck(i, x, "asinh", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Asinh(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 13:
					// atanh
					err = safecheck(i, x, "atanh", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Atanh(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 14:
					// acos
					err = safecheck(i, x, "acos", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Acos(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 15:
					// asin
					err = safecheck(i, x, "asin", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Asin(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 16:
					// atan
					err = safecheck(i, x, "atan", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Atan(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 17:
					// cos
					err = safecheck(i, x, "cos", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Cos(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 18:
					// sin
					err = safecheck(i, x, "sin", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Sin(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 19:
					// tan
					err = safecheck(i, x, "tan", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Tan(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 20:
					// натуральный логарифм ln (он же log с базой e)
					err = safecheck(i, x, "ln", 1)
					if err != nil {
						return 1, err
					}

					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value, 64)
					c.Value = fmt.Sprintf("%f", math.Log(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				}
			}
		}
	}
	end, _ = strconv.ParseFloat(x[0].Value, 64)
	return end, nil
}

func safecheck(i int, x []Piece, operand string, class int) error {
	if class == 2 {
		if i-1 < 0 && i+1 == len(x) {
			return ErrImposibleOperand("нет элемента" + operand + " нет элемента")
		}
		if i-1 >= 0 && i+1 == len(x) {
			return ErrImposibleOperand(x[i-1].Value + operand + " нет элемента")
		}
		if i-1 < 0 && i+1 < len(x) {
			return ErrImposibleOperand("нет элемента" + operand + " " + x[i+1].Value)
		}
	}
	if class == 1 {
		if i-1 < 0 && i+1 == len(x) {
			return ErrImposibleOperand(operand + "(нет элемента)")
		}
	}
	return nil
}
