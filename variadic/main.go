package main

import "fmt"

func main() {
	println("hello")
	v := []string{"Jean", "Paul", "Deux"}
	fmt.Println(explose(v...))
	fmt.Println("Jean", "Paul", "Deux")
}

func explose(s ...string) (r string) {
	for i, c := range s {
		r = r + c
		if i != len(s)-1 {
			r = r + " "
		}
	}
	return r
}
