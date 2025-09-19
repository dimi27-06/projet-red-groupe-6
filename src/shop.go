package main

import "fmt"

func (player *Character) shop() {
	for {
		var choix int
		player.LimitInventory()

		fmt.Println(Cyan + "\n=== Shop ===" + Reset)
		fmt.Println(Cyan + "Bienvenue chez Sylvebarbe l'Ent" + Reset)
		if player.FreePotion {
			fmt.Println(Pink + "1 - Potion de vie (gratuite)" + Reset)
		} else {
			fmt.Println(Pink + "1 - Potion de vie (10 écus)" + Reset)
		}
		fmt.Println(Green + "2 - Potion de poison (15 écus)" + Reset)
		fmt.Println(Blue + "3 - Potion de mana (20 écus)" + Reset)
		fmt.Println("4 - Bottes de garde du Gondor (+20 PV) [20 écus]")
		fmt.Println("5 - Jambières de garde du Gondor (+35 PV) [35 écus]")
		fmt.Println("6 - Plastron de garde du Gondor (+40 PV) [45 écus]")
		fmt.Println("7 - Casque de garde du Gondor (+25 PV) [25 écus]")
		fmt.Println(Yellow + "8 - Changer de pseudo (10 écus)" + Reset)
		fmt.Println(Red + "0 - Quitter" + Reset)
		fmt.Printf("💰 Vous avez %d écus\n", player.Gold)
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if player.FreePotion {
				player.addInventory("Potion de vie", 1)
				fmt.Println(Pink + "🎉 Vous avez obtenu : Potion de vie (gratuite) !" + Reset)
				player.FreePotion = false
			} else {
				if player.Gold >= 35 {
					player.addInventory("Potion de vie", 1)
					player.Gold -= 35
					fmt.Println("💰 Vous avez acheté une Potion de vie pour 35 écus.")
				} else {
					fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter cette potion.")
				}
			}
		case 2:
			if player.Gold >= 15 {
				player.addInventory("Potion de poison", 1)
				player.Gold -= 15
				fmt.Println("💰 Vous avez acheté une Potion de poison pour 15 écus.")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter cette potion.")
			}
		case 3:
			if player.Gold >= 20 {
				player.addInventory("Potion de mana", 1)
				player.Gold -= 20
				fmt.Println("💰 Vous avez acheté une Potion de mana pour 20 écus.")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter cette potion.")
			}
		case 4:
			if player.Gold >= 20 {
				player.addInventory("Bottes de garde du Gondor", 1)
				player.Gold -= 20
				player.PvMax += 20
				fmt.Println("🛡️ Vous avez équipé les Bottes de garde du Gondor ! (+20 PV)")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter ces bottes.")
			}
		case 5:
			if player.Gold >= 35 {
				player.addInventory("Jambières de garde du Gondor", 1)
				player.Gold -= 35
				player.PvMax += 35
				fmt.Println("🛡️ Vous avez équipé les Jambières de garde du Gondor ! (+35 PV)")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter ces jambières.")
			}
		case 6:
			if player.Gold >= 45 {
				player.addInventory("Plastron de garde du Gondor", 1)
				player.Gold -= 45
				player.PvMax += 40
				fmt.Println("🛡️ Vous avez équipé le Plastron de garde du Gondor ! (+40 PV)")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter ce plastron.")
			}
		case 7:
			if player.Gold >= 25 {
				player.addInventory("Casque de garde du Gondor", 1)
				player.Gold -= 25
				player.PvMax += 25
				fmt.Println("🛡️ Vous avez équipé le Casque de garde du Gondor ! (+25 PV)")
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour acheter ce casque.")
			}
		case 8:
			if player.Gold >= 10 {
				var newName string
				fmt.Print(Yellow + "Entrez votre nouveau pseudo : " + Reset)
				fmt.Scan(&newName)
				player.Nom = newName
				player.Gold -= 10
				fmt.Println(Yellow + "🎭 Votre pseudo est maintenant : " + player.Nom + Reset)
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'écus pour changer de pseudo.")
			}
		case 0:
			fmt.Println(Red + "Vous quittez le shop." + Reset)
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}
