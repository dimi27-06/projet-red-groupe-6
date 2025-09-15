package main

import "fmt"

func (player *Character) shop() {
	var choix int

	fmt.Println("=== Shop ===")
	if player.FreePotion {
		fmt.Println("1 - Potion de vie (gratuite)")
	} else {
		fmt.Println("1 - Potion de vie (10 écus)")
	}
	fmt.Println("2 - Potion de poison (15 écus)")
	fmt.Println("0 - Quitter")
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if player.FreePotion {
			player.addInventory("Potion de vie", 1)
			fmt.Println("🎉 Vous avez obtenu : Potion de vie (gratuite) !")
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
	case 0:
		fmt.Println("Vous quittez le shop.")
	default:
		fmt.Println("Choix invalide.")
	}
}
