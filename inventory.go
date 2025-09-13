package main

import "fmt"

// Affiche l’inventaire
func (player Character) accessInventory() {
	fmt.Println("=== Inventaire du personnage ===")
	if len(player.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
	} else {
		for _, item := range player.Inventaire {
			fmt.Printf("\t- %s x %d\n", item.Nom, item.Quantite)
		}
	}
}

// Menu inventaire
func (player *Character) MenuInventory() {
	for {
		player.accessInventory()
		fmt.Println("=== Menu inventaire ===")
		fmt.Printf("\t1 - Utiliser une potion de vie\n")
		fmt.Printf("\t0 - Retour\n")
		fmt.Print("Sélectionner un choix (1 ou 0) : ")

		var userChose int
		_, err := fmt.Scanln(&userChose)
		if err != nil {
			fmt.Println("⚠️ Entrée invalide, tape un nombre.")
			continue
		}

		switch userChose {
		case 1:
			player.takePot()
		case 0:
			return
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}
