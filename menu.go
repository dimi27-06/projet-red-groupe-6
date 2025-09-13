package main

import (
	"fmt"
	"os"
)

// Menu principal
func (player *Character) MainMenu() {
	for {
		fmt.Println("=== Menu principal ===")
		fmt.Printf("\t1 - Afficher les informations du personnage\n")
		fmt.Printf("\t2 - Accéder au contenu de l’inventaire\n")
		fmt.Printf("\t0 - Quitter\n")
		fmt.Print("Sélectionner un choix (1,2 ou 0) : ")

		var userChose int
		_, err := fmt.Scanln(&userChose)
		if err != nil {
			fmt.Println("⚠️ Entrée invalide, tape un nombre.")
			continue
		}

		switch userChose {
		case 1:
			player.displayInfo()
		case 2:
			player.MenuInventory()
		case 0:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}
