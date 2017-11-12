package main

import "math/rand"

func main() {
	motif := []string{"\\", "/"}
	for n := 0; n < 10000; n++ {
		print(motif[rand.Intn(len(motif))])
	}
}
