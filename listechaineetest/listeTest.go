package main

import "github.com/hamdouni/goExploration/listechainee"

func main() {

	var liste listechainee.Cellule

	for i := 1; i < 5; i++ {
		liste.Ajouter(99 + i)
	}

	liste.PrintAll()
}
