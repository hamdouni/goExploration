package main

import "fmt"

func main() {
	// Déclaration de constante avec iota
	// iota augmente de un pour chaque membre de la liste "const" et on peut
	// la composer dans un calcul, ici décalage de 1 bit à gauche
	const (
		PuissanceZero = 1 << iota
		PuissanceUn
		PuissanceDeux
		PuissanceTrois
		PuissanceQuatre
	)
	fmt.Printf(`
	Deux Puissance Zéro   %d 
	               Un     %d 
	               Deux   %d
	               Trois  %d
	               Quatre %d
	`, PuissanceZero, PuissanceUn, PuissanceDeux, PuissanceTrois, PuissanceQuatre)
}
