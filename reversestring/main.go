package main

import "fmt"

func main() {
	fmt.Printf("%s\n", rev("bonjour"))
}

func rev(chaine string) string {

	s := []byte(chaine)
	l := len(s) - 1
	var j int

	for i := 0; i < len(s)/2; i++ {
		j = l - i
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
