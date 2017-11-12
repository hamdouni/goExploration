package main

import "math"

func display(screen [50][100]bool) {
	for l := 0; l < 50; l++ {
		for c := 0; c < 100; c++ {
			if screen[l][c] {
				print("*")
			} else {
				print(" ")
			}
		}
		println()
	}
}

func main() {
	screen := [50][100]bool{}
	for degree := 0.0; degree < 2*math.Pi; degree += 0.01 {
		y := 24 + int(24*math.Sin(degree))
		x := 49 + int(49*math.Cos(degree))
		screen[y][x] = true
	}
	display(screen)
}
