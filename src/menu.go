package main

import (
	"fmt"
	"os"
)

// Menu principal
func (player *Character) MainMenu() {
	for {
		fmt.Println(Cyan + "=== Menu principal ===" + Reset)
		fmt.Printf("\t1 - Afficher les informations du personnage\n")
		fmt.Printf("\t2 - Accéder au contenu de l’inventaire\n")
		fmt.Printf("\t3 - Sauver les hobbits (Combat)\n")
		if player.ShopUnlocked {
			fmt.Println("\t4 - Shop\n")
		}
		fmt.Printf(Red + "\t0 - Quitter\n" + Reset)
		fmt.Print("Sélectionner un choix : ")

		var userChose int
		_, err := fmt.Scan(&userChose)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(Red + "⚠️ Entrée invalide, tape un nombre." + Reset)
			continue
		}

		switch userChose {
		case 1:
			player.displayInfo()
		case 2:
			player.MenuInventory()
		case 3:
			combatOrcs(player)
			player.ShopUnlocked = true
		case 4:
			if player.ShopUnlocked {
				player.shop()
			} else {
				fmt.Println(Red + "❌ Le shop n'est pas encore disponible." + Reset)
			}
		case 0:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)
		default:
			fmt.Println(Red + "Erreur : choix non valide" + Reset)
		}
	}

}
