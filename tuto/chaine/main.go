package main

import "fmt"

func main() {
	// Déclaration d'une chaine de caractères
	ch := "Bonjour le monde"
	fmt.Println(ch)

	// Une chaine littérale se déclare avec une back quote et permet de spécifier
	// tous les caractères qu'on veut même les sauts de ligne
	sch := `Bonjour
	le monde`
	fmt.Println(sch) // affiche la chaine avec le saut de ligne

	// Les chaines sont immutables = on ne peut pas les modifier
	//ch[2] = 'X' // va générer une erreur de compilation "cannot assign to ch[2]"

	// Sous-chaines
	sch = ch[1:3]    // ch[d:f] d l'indice du début inclu, f l'indice de fin exclu
	fmt.Println(sch) // "on"

	sch = ch[:4]     // on peut omêtre les indices, ici on part de 0
	fmt.Println(sch) // "Bonj"

	sch = ch[8:]     // la on va jusqu'à la fin = len(ch)
	fmt.Println(sch) // "le monde !"

	sch = ch[:] // ça sert à rien mais c'est fun
	fmt.Println(sch)

	// On peut parcourir une chaine grâce à "range" qui tient compte des runes
	// Un rune c'est un caractère encodé en UTF-8
	for i, e := range ch {
		fmt.Printf("%d\t%q\n", i, e)
	}

	// On concatène 2 chaines avec l'addition
	sch = ch + " et aurevoir ...."
	fmt.Println(sch)
}
