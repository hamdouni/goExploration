package main

type uneStructure struct {
	numero int
	nom    string
}

func afficheStruct(s uneStructure) {
	println(s.numero, " -> ", s.nom)
}

func doubleStruct(s *uneStructure) {
	(*s).numero = 2 * (*s).numero
}

func afficheTableau(t []int) {
	for i, c := range t {
		println(i, " -> ", c)
	}
}

func doubleTableau(t []int) {
	for i, c := range t {
		t[i] = c * 2
	}
}

func afficheNombre(n int) {
	println(n)
}

func doubleNombre(n *int) {
	*n = 2 * *n
}

func main() {
	println("bonjour le monde")

	//tableau := []int{0, 2, 4, 6, 8}
	//nombre := 100

	s := uneStructure{123, "jules"}
	afficheStruct(s)
	doubleStruct(&s)
	afficheStruct(s)
}
