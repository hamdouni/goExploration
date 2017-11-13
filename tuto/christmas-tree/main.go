package main

func main() {

	largeur := 10
	hauteur := 3

	for h := 1; h <= hauteur; h++ {
		for space := largeur - 1; space > 0; space-- {
			for s := 1; s <= space; s++ {
				print(" ")
			}
			for x := 1; x < 2*(largeur-space); x++ {
				print("*")
			}
			print("\n")
		}
	}
}
