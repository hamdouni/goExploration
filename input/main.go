package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := false
	rand.Seed(time.Now().UnixNano())
	guess := rand.Intn(100)
	count := 0
	println("Guess a number between 1 and 100. q to quit.")
	for !out {
		count++
		scanner.Scan()
		t := scanner.Text()
		if t == "q" {
			out = true
		}
		r, e := strconv.Atoi(t)
		if e != nil {
			println("not a number - try again")
		}
		switch {
		case r < guess:
			println("not enough")
		case r > guess:
			println("too much")
		default:
			fmt.Printf("you win after %v attempts.\n", count)
			out = true
		}
	}
}
