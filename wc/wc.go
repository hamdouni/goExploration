package main

import (
	"bufio"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(input.Text(), "-----")
	}
}
