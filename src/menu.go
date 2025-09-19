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
		fmt.Printf("\t3 - Sauver les hobbits 1/3 (Combat)\n")
		if player.TrollUnlocked {
			fmt.Printf("\t4 - Sauver les hobbits 2/3 (Combat)\n")
		}
		if player.SauronUnlocked {
			fmt.Printf("\t5 - Combattre SAURON (combat final)\n")
		}

		fmt.Printf("\t6 - Niveau bonus : Combattre les 9 Nazgûls (déconseillé si histoire non fini)\n")

		if player.ShopUnlocked {
			fmt.Println("\t7 - Shop")
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
			if player.TrollUnlocked {
				stopSound()
				playSoundAsyncCombat()
				combatTrollEtOrcs(player)
			} else {
				fmt.Println(Red + "🚫 Combat non disponible !" + Reset)

			}
		case 5:
			if player.SauronUnlocked {
				combatSauron(player)
				playSoundAsyncCredits()
				ClearTerminal()
				ShowCredits()
			} else {
				fmt.Println(Red + "🚫 Combat final non débloqué." + Reset)
			}
		case 6:
			playSoundAsyncNazgul()
			combatNazguls(player)

		case 7:
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
