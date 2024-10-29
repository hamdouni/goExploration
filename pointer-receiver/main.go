package main

import "fmt"

func main() {
	t := toto{
		name: "bertrand",
	}
	pabo(&t)
}

type toto struct {
	name string
}

func (t *toto) bonjour() {
	t.name = "aopzieopazei"
	fmt.Printf("bonjour %s \n", t.name)
}

func (t *toto) aurevoir() {
	println("au revoir")
}

type Joueur interface {
	bonjour()
	aurevoir()
}

func pabo(i Joueur) {
}
