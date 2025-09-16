package main

import (
	"fmt"
	"time"
)

func (player *Character) takePotS() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de vie" && player.Inventaire[index].Quantite > 0 {
			player.Pv += 20
			if player.Pv > player.PvMax {
				player.Pv = player.PvMax
			}
			fmt.Println(Pink + "Potion de vie utilisée (quantité -1)" + Reset)
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)

			player.Inventaire[index].Quantite -= 1
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println(Red + "Utilisation impossible : potion de vie manquante" + Reset)
}

func (player *Character) takePotP() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de poison" && player.Inventaire[index].Quantite > 0 {
			player.Pv -= 10
			time.Sleep(3 * time.Second)
			if player.Pv <= 0 {
				player.IsDead()
				return
			}
			fmt.Println(Green + "Potion de poison utilisée (quantité -1)" + Green)
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)

			player.Inventaire[index].Quantite -= 1
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println(Red + "Utilisation impossible : potion de poison manquante" + Reset)
}

func (player *Character) TakePotM() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de mana" && player.Inventaire[index].Quantite > 0 {
			player.Mana += 30
			if player.Mana > player.ManaMax {
				player.Mana = player.ManaMax
			}
			fmt.Println(Blue + "Potion de Mana utilisée (quantité -1)" + Reset)
			fmt.Printf("Nouveau Mana : %d\n", player.Mana)
			player.Inventaire[index].Quantite -= 1
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println(Red + "Utilisation impossible : potion de Mana manquante" + Reset)
}
