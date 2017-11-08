package main

import (
	"bufio"
	"os"
)

func input() (string, error) {
	r := bufio.NewReader(os.Stdin)
	print("Enter text: ")
	return r.ReadString('\n')
}

func main() {
	t1, _ := input()
	t2, _ := input()
	print(t1, t2)
}
