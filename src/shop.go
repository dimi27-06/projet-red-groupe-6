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
	fmt.Println(Green + "2 - Potion de poison (15 écus)" + Reset)
	fmt.Println(Green + "3 - Fourrure de loup (20 écus)" + Reset)
	fmt.Println(Green + "4 - Peau de troll (35 écus)" + Reset)
	fmt.Println(Green + "5 - Cuir de sanglier (25 écus)" + Reset)
	fmt.Println(Green + "6 - Plume de corbeau (15 écus)" + Reset)
	fmt.Println(Red + "0 - Quitter" + Reset)
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
	case 3:
		if player.Gold >= 20 {
			player.addInventory("Fourrure de loup", 1)
			player.Gold -= 20
			fmt.Println("💰Vous avez acheté une Fourrure de loup pour 20 écus.")
		} else {
			fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter une fourrure.")
		}
	case 4:
		if player.Gold >= 35 {
			player.addInventory("Peau de troll", 1)
			player.Gold -= 35
			fmt.Println("💰Vous avez acheté une Peau de troll pour 35 écus.")
		} else {
			fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter une Peau de troll.")
		}
	case 5:
		if player.Gold >= 25 {
			player.addInventory("Cuir de sanglier", 1)
			player.Gold -= 25
			fmt.Println("💰Vous avez acheté du Cuir de sanglier pour 25 écus.")
		} else {
			fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter ce Cuir.")
		}
	case 6:
		if player.Gold >= 15 {
			player.addInventory("Plume de Corbeau", 1)
			player.Gold -= 15
			fmt.Println("💰Vous avez acheté une Plume de Corbeau pour 15 écus.")
		} else {
			fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter cette Plume.")
		}
	case 0:
		fmt.Println(Red + "Vous quittez le shop." + Reset)
	default:
		fmt.Println(Red + "Choix invalide." + Reset)
	}
}
