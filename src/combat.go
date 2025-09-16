package main

import "fmt"

func (player *Character) choisirAttaque() {
	var choix int

	fmt.Println(Cyan + "\n=== Choisissez une attaque ===" + Reset)
	fmt.Println("1 - Attaque de base")
	fmt.Println("2 - Attaque spéciale (100 mana)")
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		for _, skill := range player.Skills {
			if !skill.Special {
				player.AttackSkill(skill.Nom)
				return
			}
		}
	case 2:
		for _, skill := range player.Skills {
			if skill.Special {
				player.AttackSkill(skill.Nom)
				return
			}
		}
	default:
		fmt.Println(Red + "❌ Choix invalide" + Reset)
	}
}

func (player *Character) AttackSkill(skillName string) {
	for _, skill := range player.Skills {
		if skill.Nom == skillName {
			if skill.Special && player.Mana < skill.CoutMana {
				fmt.Println(Red+"❌ Pas assez de mana pour lancer"+Reset, skill.Nom)
				return
			}
			if skill.Special {
				player.Mana -= skill.CoutMana
				fmt.Printf("🔥 %s utilise %s et inflige %d dégâts ! (Mana restant : %d/%d)\n",
					player.Nom, skill.Nom, skill.Degats, player.Mana, player.ManaMax)
			} else {
				fmt.Printf("⚔️ %s utilise %s et inflige %d dégâts !\n",
					player.Nom, skill.Nom, skill.Degats)
			}
			return
		}
	}
	fmt.Println(Red+"❌ Compétence introuvable :"+Reset, skillName)
}

func combatOrcs(player *Character) {
	fmt.Println(Yellow + "\n Début de l’histoire " + Reset)
	fmt.Println("Aragorn, Gimli et Legolas, après avoir parcouru les collines d’Emyn Muil, arrivent en lisière de la Forêt de Fangorn. Ils savent que deux Hobbits de la Communauté – Merry et Pippin – y ont été capturés par des Orques d’Isengard. Leur mission : les retrouver avant que le maléfique Saroumane ou d’autres créatures ne mettent la main sur eux.")

	orcPv := 90
	orcDegats := 8

	for {
		if orcPv <= 0 {
			fmt.Println(Green + "🎉 Victoire ! Vous avez éliminé le groupe d'orcs et sauvé les hobbits !" + Reset)
			player.ShopUnlocked = true
			return
		}

		if player.Pv <= 0 {
			player.IsDead()
			if player.Pv <= 0 {
				fmt.Println(Red + "☠️ La Communauté a échoué..." + Reset)
				return
			}
		}

		// --- Tour du joueur ---
		fmt.Println(Cyan + "\n=== Tour du joueur ===" + Reset)
		fmt.Printf("1 - %s (dégâts : %d)\n", player.BaseAttackName, player.BaseAttackDmg)
		fmt.Printf("2 - %s (dégâts : %d, coût : %d mana)\n", player.SkillName, player.SkillDmg, player.SkillManaCost)
		fmt.Print("Votre choix : ")
		var choix int
		fmt.Scan(&choix)

		var degats int
		if choix == 1 {
			degats = player.BaseAttackDmg
			fmt.Printf("💥 Vous utilisez %s et infligez %d dégâts !\n", player.BaseAttackName, degats)
		} else if choix == 2 {
			if player.Mana >= player.SkillManaCost {
				degats = player.SkillDmg
				player.Mana -= player.SkillManaCost
				fmt.Printf("🔥 Vous utilisez %s et infligez %d dégâts !\n", player.SkillName, degats)
			} else {
				fmt.Println(Red + "⚠️ Pas assez de mana ! Vous attaquez normalement." + Reset)
				degats = player.BaseAttackDmg
			}
		} else {
			fmt.Println(Red + "❌ Mauvais choix, attaque de base par défaut." + Reset)
			degats = player.BaseAttackDmg
		}

		orcPv -= degats
		if orcPv < 0 {
			orcPv = 0
		}
		fmt.Printf("👹 PV des orcs restants : %d\n", orcPv)

		if orcPv <= 0 {
			fmt.Println(Green + "🎉 Vous avez vaincu les orcs, bien joué ! Le shop Maison de Sylvebarbe est maintenant disponible. " + Reset)
			player.ShopUnlocked = true
			player.Gold += 100
			fmt.Println(Yellow + "Vous avez gagné 100 pièces d’or !" + Reset)
			return

		}

		// --- Tour des orcs ---
		fmt.Println(Cyan + "\n=== Tour des orcs ===" + Reset)
		player.Pv -= orcDegats
		if player.Pv < 0 {
			player.Pv = 0
		}
		fmt.Printf("👹 Les orcs frappent %s pour %d dégâts ! (PV : %d/%d)\n",
			player.Nom, orcDegats, player.Pv, player.PvMax)

		if player.Pv <= 0 {
			player.IsDead()
			if player.Pv <= 0 {
				fmt.Println(Red + "☠️ Vous êtes tombé définitivement... Les hobbits sont perdus." + Reset)

				gameOver()

				return
			}
		}
	}
}
