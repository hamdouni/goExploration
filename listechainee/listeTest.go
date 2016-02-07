package main

type Cellule struct {
	Suivant *Cellule
	Valeur  interface{}
}

func (c *Cellule) Ajouter(v int) {
	if c.Valeur == nil {
		c.Valeur = v
	} else {
		for c.Suivant != nil {
			c = c.Suivant
		}
		c.Suivant = &Cellule{nil, v}
	}
}

func (c *Cellule) Print() {
	println(c.Valeur.(int))
}

func (c *Cellule) PrintAll() {
	c.Print()
	for c.Suivant != nil {
		c = c.Suivant
		c.Print()
	}
}

func main() {

	var liste Cellule

	for i := 1; i < 5; i++ {
		liste.Ajouter(99 + i)
	}

	liste.PrintAll()
}
