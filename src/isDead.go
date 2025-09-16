package main

import "fmt"

func (player *Character) IsDead() {
	if player.Pv <= 0 {
		if !player.HasResurrected {
			fmt.Println(Red + "💀 Vous êtes mort !" + Reset)
			player.Pv = player.PvMax / 2
			player.HasResurrected = true
			fmt.Println(Bold, Yellow+"✨ Gandalf vous a ramené à la vie une fois !"+Reset)
		} else {
			fmt.Println(Red + "💀 Vous êtes définitivement mort... Gandalf ne viendra plus il est occupé." + Reset)
		}
	}
}
