package main

import (
	"fmt"
	"os"
)

// Menu principal
func (player *Character) MainMenu() {
	for {
		fmt.Println(Green + "=== Menu principal ===" + Reset)
		fmt.Printf("\t1 - Afficher les informations du personnage\n")
		fmt.Printf("\t2 - Accéder au contenu de l’inventaire\n")
		fmt.Printf("\t3 - Shop\n")
		fmt.Printf("\t0 - Quitter\n")
		fmt.Print("Sélectionner un choix (1,2 ou 0) : ")

		var userChose int
		_, err := fmt.Scan(&userChose)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("⚠️ Entrée invalide, tape un nombre.")
			continue
		}

		switch userChose {
		case 1:
			player.displayInfo()
		case 2:
			player.MenuInventory()
		case 3:
			player.shop()
		case 0:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}

}
