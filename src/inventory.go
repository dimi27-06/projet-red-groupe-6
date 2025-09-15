package main

import "fmt"

// Affiche l’inventaire
func (player Character) accessInventory() {
	fmt.Println(Cyan + "=== Inventaire du personnage ===" + Reset)
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
		fmt.Printf("\t2 - Utiliser une potion de poison\n")
		fmt.Printf("\t0 - Retour\n")
		fmt.Print("Sélectionner un choix (1, 2, 0) : ")

		var userChose int
		_, err := fmt.Scan(&userChose)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(Red + "⚠️ Entrée invalide, tape un nombre." + Reset)
			continue
		}
		switch userChose {
		case 2:
			player.takePotP()
		case 1:
			player.takePotS()
		case 0:
			return
		default:
			fmt.Println(Red + "Erreur : choix non valide" + Reset)
		}
	}
}

func (player *Character) addInventory(nom string, quantite int) {
	for i, item := range player.Inventaire {
		if item.Nom == nom {

			player.Inventaire[i].Quantite += quantite
			return
		}
	}
	player.Inventaire = append(player.Inventaire, Item{Nom: nom, Quantite: quantite})
}

func (player *Character) removeInventory(nom string, quantite int) {
	for i, item := range player.Inventaire {
		if item.Nom == nom {
			if player.Inventaire[i].Quantite > quantite {
				player.Inventaire[i].Quantite -= quantite
			} else {
				player.Inventaire = append(player.Inventaire[:i], player.Inventaire[i+1:]...)
			}
			return
		}
	}
}
