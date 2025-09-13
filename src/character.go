package main

import "fmt"

// Structure Personnage
type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	PvMax      int
	Inventaire []Item
}

// Structure Item
type Item struct {
	Nom      string
	Quantite int
}

// Initialisation du personnage
func (player *Character) initCharacter() {
	*player = Character{
		Nom:    "cyril",
		Classe: "Elfe",
		Niveau: 1,
		Pv:     40,
		PvMax:  100,
		Inventaire: []Item{
			{"Épée", 1},
			{"Potion de vie", 2},
		},
	}
}

// Affichage infos personnage
func (player Character) displayInfo() {
	fmt.Println("=== Information du personnage ===")
	fmt.Printf("\t - Nom : %s\n", player.Nom)
	fmt.Printf("\t - Classe : %s\n", player.Classe)
	fmt.Printf("\t - Niveau : %d\n", player.Niveau)
	fmt.Printf("\t - Pv : %d\n", player.Pv)
	fmt.Printf("\t - PvMax : %d\n", player.PvMax)
}
