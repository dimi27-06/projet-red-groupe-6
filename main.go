package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	Pvmax      int
	Inventaire []Item
}

type Item struct {
	Nom      string
	Quantite int
}

func (player *Character) initCharacter() {
	*player = Character{"aragorn", "humain", 1, 100, 100, []Item{}}
}

func (player Character) displayInfo() {
	fmt.Println("=== info perso ===")
	fmt.Printf("\t- Nom : %s")
}

func main() {
	p1 := Character{}
	p1.initCharacter()
	fmt.Println(p1)
}
