package main

import (
	"math/rand"
	"time"
)

func maze() {
	motif := []string{"\\", "/"}
	for m := 0; m < 10; m++ {
		for n := 0; n < 40; n++ {
			print(motif[rand.Intn(len(motif))])
		}
		println()
	}
}

func main() {
	for i := 0; i < 50; i++ {
		//print("\x0c") // Clear screen for playground
		print("\033[2J") // Clear screen for terminal

		maze()
		time.Sleep(time.Second / 5)
	}
}
