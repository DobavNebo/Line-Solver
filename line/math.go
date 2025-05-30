package main

import (
	"fmt"
	"strconv"
	"math"
	"slices"
)

func Calculate(x []Piece) float64{
	x = escapeParents(x)
	x, y := priorities(x)
	return solve(x, y)
}

func escapeParents(x []Piece) []Piece{
	var openers []int 
	bump := len(x)
	for i, v := range x {
		c := &v
		if c.Class == -1{
			openers = append(openers, i)
			i++
			continue
		}
	}

	for len(openers) > 0 {
		A := openers[len(openers)-1]
		for i := A; i < len(x); i++ {
			c := &x[i]
			if c.Class == -2 {
				x = append(x[:i], x[(i+1):]...)
				x = append(x[:A], x[(A+1):]...)
				break
			}
			if c.Class > 0 {
				c.Prior += 4*bump
			}
		}
		openers = openers[:(len(openers)-1)]
	}

	return x
}

func priorities(x []Piece) ([]Piece, []int){
	leng := len(x)
	var priors []int

	g := 0
	for i := 0; i < leng; i++ {
		c := &x[i]
		if isunar(c.Class) {
			c.Prior += g+1+(3*leng)
			g++
			priors = append(priors, c.Prior)
		}
	}
	g = 0
	for i := 0; i < leng; i++ {
		c := &x[i]
		if c.Class == 6 {
			c.Prior += g+1+(2*leng)
			g++
			priors = append(priors, c.Prior)
		}
	}
	g = 0
	for i := leng-1; i >= 0; i-- {
		c := &x[i]
		if istwo(c.Class) {
			c.Prior += g+1+leng
			g++
			priors = append(priors, c.Prior)
		}
	}
	g = 0
	for i := leng-1; i >= 0; i-- {
		c := &x[i]
		if isone(c.Class) {
			c.Prior += g+1
			g++
			priors = append(priors, c.Prior)
		}
	}
	slices.Sort(priors)
	slices.Reverse(priors)
	return x, priors
}

func isone(a int) bool{
	return a == 1 || a == 2
}

func istwo(a int) bool{
	return a == 3 || a == 4 || a == 5
}

func isunar(a int) bool{
	return a >= 10 && a <= 20
}

func solve(x []Piece, priors []int) float64 {
	for _, g := range priors {
		//for debug purposes
		//fmt.Println(x)
		for i := 0; i < len(x) ; i++ {
			c := &x[i]
			if c.Prior == g {
				switch c.Class {
				case 1:
					//+
					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value,64)
					b, _ := strconv.ParseFloat(nxt.Value,64)
					pr.Value = fmt.Sprintf("%f", a+b)
					x = append(x[:i], x[(i+2):]...)
				case 2:
					//-
					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value,64)
					b, _ := strconv.ParseFloat(nxt.Value,64)
					pr.Value = fmt.Sprintf("%f", a-b)
					x = append(x[:i], x[(i+2):]...)
				case 3:
					//*
					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value,64)
					b, _ := strconv.ParseFloat(nxt.Value,64)
					pr.Value = fmt.Sprintf("%f", a*b)
					x = append(x[:i], x[(i+2):]...)
				case 4:
					// /
					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value,64)
					b, _ := strconv.ParseFloat(nxt.Value,64)
					pr.Value = fmt.Sprintf("%f", a/b)
					x = append(x[:i], x[(i+2):]...)
				case 5:
					// %
					pr := &x[i-1]
					nxt := &x[i+1]
					pro, _ := strconv.ParseFloat(pr.Value,64)
					a := int(pro)
					nxto, _ := strconv.ParseFloat(nxt.Value,64)
					b := int(nxto)
					pr.Value = fmt.Sprintf("%f", float64(a%b))
					x = append(x[:i], x[(i+2):]...)
				case 6:
					// ^
					pr := &x[i-1]
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(pr.Value,64)
					b, _ := strconv.ParseFloat(nxt.Value,64)
					pr.Value = fmt.Sprintf("%f", math.Pow(a, b))
					x = append(x[:i], x[(i+2):]...)
				case 10:
					// - unar
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", (-1.0)*a)
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 11:
					// acosh
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Acosh(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 12:
					// asinh
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Asinh(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 13:
					// atanh
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Atanh(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 14:
					// acos
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Acos(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 15:
					// asin
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Asin(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 16:
					// atan
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Atan(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 17:
					// cos
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Cos(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 18:
					// sin
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Sin(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 19:
					// tan
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Tan(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				case 20:
					// log on base e
					nxt := &x[i+1]
					a, _ := strconv.ParseFloat(nxt.Value,64)
					c.Value = fmt.Sprintf("%f", math.Log(a))
					c.Prior = 0
					c.Class = 0
					x = append(x[:i+1], x[(i+2):]...)
				}
			}
		}
	}
	sh := &x[0]
	end, _ := strconv.ParseFloat(sh.Value,64)
	return end
}