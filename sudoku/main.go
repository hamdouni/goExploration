package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// sdk := [][]int{
	// 	{0, 0, 7, 9, 0, 3, 4, 0, 0},
	// 	{0, 2, 0, 7, 0, 8, 0, 1, 0},
	// 	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	// 	{5, 0, 0, 0, 6, 0, 0, 0, 7},
	// 	{0, 0, 4, 1, 0, 9, 6, 0, 0},
	// 	{8, 0, 0, 0, 7, 0, 0, 0, 1},
	// 	{0, 0, 0, 0, 2, 0, 0, 0, 0},
	// 	{0, 4, 0, 8, 0, 7, 0, 6, 0},
	// 	{0, 0, 3, 5, 0, 1, 8, 0, 0},
	// }
	// sdk := [][]int{
	// 	{0, 0, 0, 9, 8, 0, 0, 7, 0},
	// 	{6, 0, 0, 2, 0, 0, 9, 0, 0},
	// 	{0, 0, 9, 0, 0, 0, 3, 0, 4},
	// 	{0, 0, 0, 0, 3, 0, 1, 0, 0},
	// 	{0, 8, 0, 0, 5, 0, 0, 9, 0},
	// 	{0, 0, 6, 0, 2, 0, 0, 0, 0},
	// 	{4, 0, 8, 0, 0, 0, 2, 0, 0},
	// 	{0, 0, 5, 0, 0, 2, 0, 0, 8},
	// 	{0, 3, 0, 0, 7, 1, 0, 0, 0},
	// }
	sdk := [][]int{
		//		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 7, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 0, 8, 0, 4},
		{0, 6, 0, 0, 0, 2, 0, 0, 7},
		{0, 2, 0, 0, 0, 3, 0, 0, 6},
		{0, 0, 0, 2, 0, 5, 0, 0, 0},
		{3, 0, 0, 4, 0, 0, 0, 8, 0},
		{6, 0, 0, 8, 0, 0, 0, 9, 0},
		{5, 0, 1, 0, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 5, 0},
	}

	var recommence, trouve bool = true, true
	for recommence {
		recommence = false
		trouve = false
		for l := 0; l < 9; l++ {
			for c := 0; c < 9; c++ {
				if sdk[l][c] != 0 {
					continue
				}
				recommence = true
				possibilities := solveByElimination(sdk, l, c)
				if len(possibilities) == 1 {
					trouve = true
					sdk[l][c] = possibilities[0]
					// fmt.Printf("Finded by elimination %v %v => %v\n", l+1, c+1, possibilities[0])
					continue
				}
				sol := solveByPosition(sdk, l, c, possibilities)
				if sol != -1 {
					trouve = true
					sdk[l][c] = sol
					// fmt.Printf("Finded by position %v %v => %v\n", l+1, c+1, sol)
					continue
				}
				res := solveByIntersection(sdk, l, c, possibilities)
				if len(res) == 1 {
					trouve = true
					sdk[l][c] = res[0]
					// fmt.Printf("Finded by intersection %v %v => %v\n", l+1, c+1, possibilities[0])
				}

				if sdk[l][c] == 0 {
					view(sdk, l, c)
					duration := 20 * time.Millisecond
					time.Sleep(duration)
					// 	input()
				}
			}
		}
		if !trouve {
			break
		}
	}
	view(sdk, -1, -1)
}

func view(s [][]int, y, x int) {
	fmt.Print("\033[2J")
	for j, l := range s {
		if j%3 == 0 {
			println("-------------------------")
		}
		for i, c := range l {
			if i%3 == 0 {
				print("| ")
			}
			if c != 0 {
				fmt.Printf("%v ", c)
				continue
			}
			if y == j && x == i {
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("-------------------------")
}

func getLine(s [][]int, l int) []int {
	return s[l]
}

func getCol(s [][]int, c int) (col []int) {
	for l := 0; l < 9; l++ {
		col = append(col, s[l][c])
	}
	return col
}

func getCube(s [][]int, l, c int) (cub []int) {
	lx := c - (c % 3)
	ly := l - (l % 3)
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			cub = append(cub, s[ly+j][lx+i])
		}
	}
	return cub
}

func solveByElimination(s [][]int, l, c int) []int {
	line := getLine(s, l)
	col := getCol(s, c)
	cub := getCube(s, l, c)
	possibilities := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, elem := range line {
		if elem != 0 {
			possibilities[elem-1] = -1
		}
	}
	for _, elem := range col {
		if elem != 0 {
			possibilities[elem-1] = -1
		}
	}
	for _, elem := range cub {
		if elem != 0 {
			possibilities[elem-1] = -1
		}
	}
	resp := []int{}
	for _, v := range possibilities {
		if v != -1 {
			resp = append(resp, v)
		}
	}
	return resp
}
func solveByPosition(s [][]int, l, c int, p []int) (sol int) {
	lx := c - (c % 3)
	ly := l - (l % 3)
	var common []int
	var pos int
	for _, potentiel := range p {
		pos = 0
		for j := 0; j < 3; j++ {
			for i := 0; i < 3; i++ {
				if s[ly+j][lx+i] == 0 {
					common = solveByElimination(s, ly+j, lx+i)
					for _, el := range common {
						if el == potentiel {
							pos++
						}
					}
				}
			}
		}
		if pos == 1 {
			return potentiel
		}
	}
	return -1
}

func solveByIntersection(s [][]int, l, c int, p []int) (possibilities []int) {
	var ol1, ol2, oc1, oc2 []int // other lines and columns
	if l%3 == 0 {
		ol1 = getLine(s, l+1)
		ol2 = getLine(s, l+2)
	} else if l%3 == 1 {
		ol1 = getLine(s, l-1)
		ol2 = getLine(s, l+1)
	} else {
		ol1 = getLine(s, l-1)
		ol2 = getLine(s, l-2)
	}
	if c%3 == 0 {
		oc1 = getCol(s, c+1)
		oc2 = getCol(s, c+2)
	} else if c%3 == 1 {
		oc1 = getCol(s, c-1)
		oc2 = getCol(s, c+1)
	} else {
		oc1 = getCol(s, c-1)
		oc2 = getCol(s, c-2)
	}
	possibilities = p
	if len(ol1) > 0 {
		possibilities = intersection(possibilities, ol1)
	}
	if len(ol2) > 0 {
		possibilities = intersection(possibilities, ol2)
	}
	if len(oc1) > 0 {
		possibilities = intersection(possibilities, oc1)
	}
	if len(oc2) > 0 {
		possibilities = intersection(possibilities, oc2)
	}
	resp := []int{}
	for _, v := range possibilities {
		if v != 0 {
			resp = append(resp, v)
		}
	}
	return resp
}

func intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return c
}

func input() (string, error) {
	r := bufio.NewReader(os.Stdin)
	print("Enter text: ")
	return r.ReadString('\n')
}
