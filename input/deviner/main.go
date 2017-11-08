package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	guess := rand.Intn(10)
	count := 0
	println("Guess a number between 1 and 10. q to quit. ==> ", guess)
loop:
	for {
		count++
		scanner.Scan()
		t := scanner.Text()
		if t == "q" {
			break
		}
		r, e := strconv.Atoi(t)
		if e != nil {
			println("not a number - try again")
			continue
		}
		switch {
		case r < guess:
			println("not enough")
		case r > guess:
			println("too much")
		default:
			println("you win after", count, "attempts.")
			break loop
		}
	}
}
