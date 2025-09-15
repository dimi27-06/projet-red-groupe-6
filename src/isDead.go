package main

import "fmt"

func (player *Character) IsDead() {
	if player.Pv <= 0 {
		fmt.Println(Red + "vous êtes mort" + Reset)
		player.Pv = player.PvMax / 2
		fmt.Println("Gandalf vous a ramené à la vie")
	}
}
