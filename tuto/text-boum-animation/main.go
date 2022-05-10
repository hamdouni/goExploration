package main

import (
	"fmt"
	"time"
)

func main() {
	var ch []string
	var s string
	var cls = "\033[2J\x0c" // \033[2J clear screan on terminal and \x0c on playground

	c := ")"
	for z := 20; z >= 0; z-- {
		ch = append(ch, s)
		s += c
	}
	for i := 2; i >= 0; i-- {
		for j := 0; j < len(ch); j++ {
			println(cls, ch[j])
			time.Sleep(time.Second / 15)
		}
		for j := len(ch) - 1; j >= 0; j-- {
			println(cls, ch[j])
			time.Sleep(time.Second / 15)
		}
	}
	println(cls, "BOUM")
	fmt.Printf("Hello le %v\n", "Rendu")
}
