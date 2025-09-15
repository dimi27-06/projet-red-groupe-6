package main

import "fmt"

func (player *Character) IsDead() {
	if player.Pv <= 0 {
		fmt.Println("vous êtes mort")
		player.Pv = player.PvMax / 2
		fmt.Println("Gandalf vous a ramené à la vie")
	}
}
