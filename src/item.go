package main

import (
	"fmt"
	"time"
)

// Utiliser une potion de vie
func (player *Character) takePotS() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de vie" && player.Inventaire[index].Quantite > 0 {
			player.Pv += 50
			if player.Pv > player.PvMax {
				player.Pv = player.PvMax
			}
			fmt.Println("Potion de vie utilisée (quantité -1)")
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)

			player.Inventaire[index].Quantite -= 1
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println("Utilisation impossible : potion de vie manquante")
}

// Utiliser une potion de poison
func (player *Character) takePotP() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de poison" && player.Inventaire[index].Quantite > 0 {
			player.Pv -= 10
			time.Sleep(3 * time.Second)
			if player.Pv <= 0 {
				return ((func(player *Character) IsDead.))
			}
			fmt.Println("Potion de vie utilisée (quantité -1)")
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)

			player.Inventaire[index].Quantite -= 1
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println("Utilisation impossible : potion de poison manquante")
}
