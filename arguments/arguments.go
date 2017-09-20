package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Salam le world")

	args := os.Args[1:] // args without prog name
	l := len(args)
	if l == 0 {
		fmt.Println("Aucun argment")
		return
	}
	fmt.Println("Les arguments sont:")
	for i, v := range args {
		fmt.Printf("\t%d : \t %v\n", i, v)
	}
}
