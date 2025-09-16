package main

import "fmt"

func (player *Character) shop() {
	var choix int

	fmt.Println(Cyan + "=== Shop ===" + Reset)
	fmt.Println(Cyan + "Bienvenue chez Sylvebarbe l'Ent" + Reset)
	if player.FreePotion {
		fmt.Println(Pink + "1 - Potion de vie (gratuite)" + Reset)
	} else {
		fmt.Println(Pink + "1 - Potion de vie (10 écus)" + Reset)
	}
	fmt.Println("2 - Potion de poison (15 écus)")
	fmt.Println("0 - Quitter")
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if player.FreePotion {
			player.addInventory("Potion de vie", 1)
			fmt.Println(Pink + "🎉 Vous avez obtenu : Potion de vie (gratuite) !" + Reset)
			player.FreePotion = false
		} else {
			if player.Gold >= 10 {
				player.addInventory("Potion de vie", 1)
				player.Gold -= 10
				fmt.Println("💰 Vous avez acheté une Potion de vie pour 10 écus.")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter cette potion.")
			}
		}
	case 2:
		if player.Gold >= 15 {
			player.addInventory("Potion de poison", 1)
			player.Gold -= 15
			fmt.Println("💰Vous avez acheté une Potion de poison pour 10 écus.")
		} else {
			fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter cette potion.")
		}
	case 0:
		fmt.Println(Red + "Vous quittez le shop." + Reset)
	default:
		fmt.Println(Red + "Choix invalide." + Reset)
	}
}
