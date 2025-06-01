package main

func Zandatsu(x []Piece) []Piece {
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

	for i, _ := range result {
		current := &result[i]
		if current.Class == -3 {
			g := 0
			for g = 0; g < len(lefts); g++ {
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
