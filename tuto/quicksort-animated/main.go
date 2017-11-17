// Quick Sort in Golang
package main

import (
	"math/rand"
	"time"
)

var tableau []int

func main() {
	for i := 0; i < 30; i++ {
		tableau = append(tableau, 10+rand.Intn(90))
	}
	qs(tableau)
}

func qs(t []int) {
	long := len(t) - 1
	if long < 1 {
		return
	}

	left, right := 0, long
	for i := 0; i <= long; i++ {
		if t[i] < t[right] {
			permut(t, left, i)
			left++
		}
	}

	permut(t, left, right)
	qs(t[:left])
	qs(t[left+1:])
}

func showTableau(start, end int) {
	print("\033[2J\x0cQUICK SORTING\n")
	for i, c := range tableau {
		if i == start || i == end {
			print(">")
		} else {
			print(" ")
		}
		print(" ", c, " ")
		for i := 0; i < c; i++ {
			print("*")
		}
		println()
	}
	time.Sleep(time.Second / 10)
}

func permut(t []int, i, j int) {
	showTableau(i, j)
	t[i], t[j] = t[j], t[i]
	showTableau(i, j)
}
