package main

import (
	"fmt"
	"strconv"
)

func Zandatsu(x []Piece) []Piece{
	var rights []Piece
	var result []Piece
	var lefts []string
	
	for i, _ := range x {
		current := &x[i]
		if current.Class == -3 {
			if len(lefts) > 0 {
				check := true
				for _, val := range lefts {
					if current.Value == val {
						check = false
					}
				}
				if check {
					lefts = append(lefts, current.Value)
				}
			} else {
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

	rights = onlyvalues(rights)
	for i, _ := range result {
		current := &result[i]
		if current.Class == -3 {
			g := 0
			for g = 0; g < len(lefts); g++{
				if current.Value == lefts[g] {
					newval := &rights[g]
					current.Value = newval.Value
					current.Class = 0
					break
				}
			}
		}
	}
	return result
}

func onlyvalues(x []Piece) []Piece{
	for i := 0; i < len(x); i++{
		c := &x[i]
		if c.Class == 10 {
			nxt := &x[i+1]
			a, _ := strconv.ParseFloat(nxt.Value,64)
			c.Value = fmt.Sprintf("%f", (-1.0)*a)
			c.Prior = 0
			c.Class = 0
			x = append(x[:i+1], x[(i+2):]...)
			continue
		}
		if c.Class == 3 {
			x = append(x[:i], x[(i+1):]...)
			continue
		}
	}
	return x
}