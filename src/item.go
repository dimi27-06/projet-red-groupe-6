package main

import (
	"fmt"
)

func (player *Character) TakePotS() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de vie" && player.Inventaire[index].Quantite > 0 {
			player.Pv += 20
			if player.Pv > player.PvMax {
				player.Pv = player.PvMax
			}
			fmt.Println(Pink + "💖 Potion de vie utilisée (quantité -1)" + Reset)
			fmt.Printf("Nouveau PV : %d/%d\n", player.Pv, player.PvMax)

			player.Inventaire[index].Quantite--
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println(Red + "❌ Pas de potion de vie disponible" + Reset)
}

func (player *Character) TakePotP(enemyPv *int) {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de poison" && player.Inventaire[index].Quantite > 0 {
			*enemyPv -= 15
			if *enemyPv < 0 {
				*enemyPv = 0
			}
			fmt.Println(Green + "☠️ Potion de poison utilisée (quantité -1)" + Reset)
			fmt.Printf("PV de l’ennemi après poison : %d\n", *enemyPv)

			player.Inventaire[index].Quantite--
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println(Red + "❌ Pas de potion de poison disponible" + Reset)
}

func (player *Character) TakePotM() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de mana" && player.Inventaire[index].Quantite > 0 {
			player.Mana += 30
			if player.Mana > player.ManaMax {
				player.Mana = player.ManaMax
			}
			fmt.Println(Blue + "🔮 Potion de mana utilisée (quantité -1)" + Reset)
			fmt.Printf("Nouveau Mana : %d/%d\n", player.Mana, player.ManaMax)

			player.Inventaire[index].Quantite--
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println(Red + "❌ Pas de potion de mana disponible" + Reset)
}
