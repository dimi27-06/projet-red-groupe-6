package main

import (
	"fmt"
	"math/rand"
	"time"
)

func (player *Character) choisirAttaque() {

	player.TakePotM()
	player.TakePotS()
	var enemyHp int
	player.TakePotP(&enemyHp)
	var choix int

	fmt.Println(Cyan + "\n=== Choisissez une attaque ===" + Reset)
	fmt.Println("1 - Attaque de base")
	fmt.Println("2 - Attaque spéciale (100 mana)")
	fmt.Println("3 - Potion de vie")
	fmt.Println("4 - Potion de poison")
	fmt.Println("5 - Potion de mana")
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
				fmt.Printf("\n⚔️ %s utilise %s et inflige %d dégâts !\n",
					player.Nom, skill.Nom, skill.Degats)
			}
			return
		}
	}
	fmt.Println(Red+"❌ Compétence introuvable :"+Reset, skillName)
}

func combatOrcs(player *Character) {
	stopSound()
	playSoundAsyncCombat()
	fmt.Println(Bold, Yellow+"\n Début de l’histoire "+Reset)

	histoire := `Aragorn, Gimli et Legolas, après avoir parcouru les collines d’Emyn Muil, 
arrivent en lisière de la Forêt de Fangorn. 
Ils savent que deux Hobbits de la Communauté – Merry et Pippin – y ont été capturés par des Orques d’Isengard. 
Leur mission : les retrouver avant que le maléfique Saroumane ou d’autres créatures ne mettent la main sur eux.`
	typeWriter(histoire, 7*time.Millisecond)

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
		fmt.Println(Cyan + "\n=== Tour du joueur ===" + Reset)
		fmt.Printf("1 - %s (dégâts : %d)\n", player.BaseAttackName, player.BaseAttackDmg)
		fmt.Printf("2 - %s (dégâts : %d, coût : %d mana)\n", player.SkillName, player.SkillDmg, player.SkillManaCost)
		fmt.Printf("3 - Potion de vie (vie régénérée : 20, -1 potion de vie)\n")
		fmt.Printf("4 - Potion de poison (inflige : 10 dégâts, -1 potion de poison)\n")
		fmt.Printf("5 - Potion de mana (mana régénéré : 30, -1 potion de mana)\n")
		fmt.Print("Votre choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			// attaque de baseeeeee
		case 2:
			// attaque spécialeeeee
		case 3:
			player.TakePotS()
		case 4:
			player.TakePotP(&orcPv)
		case 5:
			player.TakePotM()
		default:
			fmt.Println(Red + "❌ Choix invalide !" + Reset)
		}

		var degats int
		switch choix {
		case 1:
			degats = player.BaseAttackDmg
			fmt.Printf("\n💥 Vous utilisez %s et infligez %d dégâts !\n", player.BaseAttackName, degats)
		case 2:
			if player.Mana >= player.SkillManaCost {
				degats = player.SkillDmg
				player.Mana -= player.SkillManaCost
				fmt.Printf("🔥 Vous utilisez %s et infligez %d dégâts !\n", player.SkillName, degats)
			} else {
				fmt.Println(Red + "⚠️ Pas assez de mana ! Vous attaquez normalement." + Reset)
				degats = player.BaseAttackDmg
			}
		default:
			fmt.Println(Red + "❌ Mauvais choix, attaque de base par défaut." + Reset)
			degats = player.BaseAttackDmg
		}

		orcPv -= degats
		if orcPv < 0 {
			orcPv = 0
		}
		fmt.Printf("👹 PV des orcs restants : %d\n", orcPv)

		if orcPv <= 0 {
			stopSound()
			fmt.Println(Green + "🎉 Vous avez vaincu les orcs, bien joué ! Le shop Maison de Sylvebarbe est maintenant disponible. " + Reset)
			player.ShopUnlocked = true
			player.Gold += 100
			fmt.Println(Yellow + "Vous avez gagné 100 pièces d’or !" + Reset)
			player.TrollUnlocked = true
			playSoundAsyncDebut()
			return

		}

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
				stopSound()
				fmt.Println(Red + "☠️ Vous êtes tombé définitivement... Les hobbits sont perdus." + Reset)

				gameOver()

				return
			}
		}
	}
}

func combatTrollEtOrcs(player *Character) {
	stopSound()
	playSoundAsyncCombat()
	histoire2 := Yellow + "\n apres avoir vaincu les orcs notre trio continue leurs route mais La forêt devient plus sombre. Après une courte accalmie, deux Orques, accompagnés d’un Troll, attaquent. " + Reset
	typeWriter(histoire2, 7*time.Millisecond)

	ennemiPv := 200
	ennemiDegats := 25

	for {
		if ennemiPv <= 0 {
			fmt.Println(Green + "🎉 Victoire ! Vous avez vaincu le Troll et ses orcs !" + Reset)
			player.Gold += 200
			fmt.Println(Yellow + "💰 Vous gagnez 200 pièces d’or !" + Reset)
			return
		}

		if player.Pv <= 0 {
			player.IsDead()
			if player.Pv <= 0 {
				fmt.Println(Red + "☠️ Le Troll et ses orcs vous ont écrasé... C’est la fin." + Reset)
				gameOver()
				return
			}
		}

		fmt.Println(Cyan + "\n=== Tour du joueur ===" + Reset)
		fmt.Printf("1 - %s (dégâts : %d)\n", player.BaseAttackName, player.BaseAttackDmg)
		fmt.Printf("2 - %s (dégâts : %d, coût : %d mana)\n", player.SkillName, player.SkillDmg, player.SkillManaCost)
		fmt.Printf("3 - Potion de vie (vie régénérée : 20, -1 potion de vie)\n")
		fmt.Printf("4 - Potion de poison (inflige : 10 dégâts, -1 potion de poison)\n")
		fmt.Printf("5 - Potion de mana (mana régénéré : 30, -1 potion de mana)\n")
		fmt.Print("Votre choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			// attaque de baseeeeee
		case 2:
			// attaque spécialeeeee
		case 3:
			player.TakePotS()
		case 4:
			player.TakePotP(&ennemiPv)
		case 5:
			player.TakePotM()
		default:
			fmt.Println(Red + "❌ Choix invalide !" + Reset)
		}
		var degats int
		switch choix {
		case 1:
			degats = player.BaseAttackDmg
			fmt.Printf("\n💥 Vous utilisez %s et infligez %d dégâts !\n", player.BaseAttackName, degats)
		case 2:
			if player.Mana >= player.SkillManaCost {
				degats = player.SkillDmg
				player.Mana -= player.SkillManaCost
				fmt.Printf("\n🔥 Vous utilisez %s et infligez %d dégâts !\n", player.SkillName, degats)
			} else {
				fmt.Println(Red + "⚠️ Pas assez de mana ! Attaque normale." + Reset)
				degats = player.BaseAttackDmg
			}
		}

		ennemiPv -= degats
		if ennemiPv < 0 {
			ennemiPv = 0
		}
		fmt.Printf("👹 PV du Troll + Orcs : %d\n", ennemiPv)

		if ennemiPv <= 0 {
			fmt.Println(Green + "🎉 Victoire ! Vous avez vaincu le Troll et ses orcs !" + Reset)
			player.Gold += 200
			fmt.Println(Yellow + "💰 Vous gagnez 200 pièces d’or !" + Reset)
			player.SauronUnlocked = true
			fmt.Println(Cyan + "\n🔔 Le combat final contre SAURON est désormais disponible dans le menu principal !" + Reset)
			stopSound()
			playSoundAsyncDebut()

			return
		}

		fmt.Println(Cyan + "\n=== Tour des ennemis ===" + Reset)
		player.Pv -= ennemiDegats
		if player.Pv < 0 {
			player.Pv = 0
		}
		fmt.Printf("👹 Le Troll et les orcs frappent %s pour %d dégâts ! (PV : %d/%d)\n",
			player.Nom, ennemiDegats, player.Pv, player.PvMax)

		if player.Pv <= 0 {
			player.IsDead()
			if player.Pv <= 0 {
				fmt.Println(Red + "☠️ Vous êtes tombé face au Troll et ses orcs..." + Reset)
				gameOver()
				playSoundAsyncCombat()
				playSoundAsyncDebut()
				return
			}
		}
	}
}

func combatSauron(player *Character) {
	stopSound()
	playSoundAsyncSauron()
	histoire3 := Red + Bold + "\n⚔️ Après avoir vaincu le Troll et ses orcs, nos héros marchent vers le Mordor.\nDans les flammes de la montagne du Destin, Sauron surgit avec sa masse gigantesque... Le combat final commence !" + Reset
	typeWriter(histoire3, 7*time.Millisecond)
	sauronPv := 400

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		if sauronPv <= 0 {
			stopSound()
			fmt.Println(Green + Bold + "🎉 Victoire ultime ! Vous avez terrassé Sauron et sauvé la Terre du Milieu et les hobbits!" + Reset)
			player.Gold += 500
			fmt.Println(Yellow + "💰 Récompense : 500 pièces d’or !" + Reset)
			player.SauronUnlocked = true
			playSoundAsyncDebut()
			return
		}

		if player.Pv <= 0 {
			player.IsDead()
			if player.Pv <= 0 {
				stopSound()
				fmt.Println(Red + Bold + "☠️ Sauron vous a anéanti... les hobbits sont perdu." + Reset)
				gameOver()
				return
			}
		}

		fmt.Println(Cyan + "\n=== Tour du joueur ===" + Reset)
		fmt.Printf("1 - %s (dégâts : %d)\n", player.BaseAttackName, player.BaseAttackDmg)
		fmt.Printf("2 - %s (dégâts : %d, coût : %d mana)\n", player.SkillName, player.SkillDmg, player.SkillManaCost)
		fmt.Printf("3 - Potion de vie (vie régénérée : 20, -1 potion de vie)\n")
		fmt.Printf("4 - Potion de poison (inflige : 10 dégâts, -1 potion de poison)\n")
		fmt.Printf("5 - Potion de mana (mana régénéré : 30, -1 potion de mana)\n")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			degats := player.BaseAttackDmg
			sauronPv -= degats
			fmt.Printf("\n⚔️ Vous utilisez %s et infligez %d dégâts à Sauron !\n", player.BaseAttackName, degats)
		case 2:
			if player.Mana >= player.SkillManaCost {
				degats := player.SkillDmg
				player.Mana -= player.SkillManaCost
				sauronPv -= degats
				fmt.Printf("\n🔥 Vous utilisez %s et infligez %d dégâts à Sauron ! (Mana : %d/%d)\n", player.SkillName, degats, player.Mana, player.ManaMax)
			} else {
				fmt.Println(Red + "⚠️ Pas assez de mana ! Vous attaquez normalement." + Reset)
				degats := player.BaseAttackDmg
				sauronPv -= degats
				fmt.Printf("\n⚔️ Attaque de base : %d dégâts à Sauron.\n", degats)
			}
		case 3:
			player.TakePotS()
		case 4:
			player.TakePotP(&sauronPv)
		case 5:
			player.TakePotM()
		default:
			fmt.Println(Red + "❌ Choix invalide, attaque de base par défaut !" + Reset)
			degats := player.BaseAttackDmg
			sauronPv -= degats
			fmt.Printf("\n⚔️ Attaque de base : %d dégâts à Sauron.\n", degats)
		}

		if sauronPv < 0 {
			sauronPv = 0
		}
		fmt.Printf("👹 PV restants de Sauron : %d\n", sauronPv)

		if sauronPv <= 0 {
			stopSound()
			fmt.Println(Green + Bold + "🎉 Vous avez terrassé Sauron ! La lumière triomphe des ténèbres !" + Reset)
			return
		}

		fmt.Println(Red + "\n=== Tour de Sauron ===" + Reset)

		attaque := rnd.Intn(3)
		switch attaque {
		case 0:
			degats := 40
			player.Pv -= degats
			fmt.Printf(Red+Bold+"💀 Sauron frappe avec sa masse et inflige %d dégâts ! (PV : %d/%d)\n"+Reset,
				degats, player.Pv, player.PvMax)
		case 1:
			degats := 60
			player.Pv -= degats
			fmt.Printf(Red+Bold+"💥 Sauron utilise une FRAPPE ÉCRASANTE et inflige %d dégâts dévastateurs ! (PV : %d/%d)\n"+Reset,
				degats, player.Pv, player.PvMax)
		case 2:
			degats := 50
			player.Pv -= degats
			player.Mana -= 20
			if player.Mana < 0 {
				player.Mana = 0
			}
			fmt.Printf(Red+Bold+"🔥 Sauron invoque les FLAMMES DU MORDOR et inflige %d dégâts ! Vous perdez aussi 20 mana ! (PV : %d/%d | Mana : %d/%d)\n"+Reset,
				degats, player.Pv, player.PvMax, player.Mana, player.ManaMax)
		}

		if player.Pv < 0 {
			player.Pv = 0
		}

		if sauronPv <= 0 {
			stopSound()
			fmt.Println(Green + Bold + "🎉 Vous avez terrassé Sauron ! La lumière triomphe des ténèbres !" + Reset)
			time.Sleep(2 * time.Second)
		}

	}

}

func combatNazguls(player *Character) {
	stopSound()
	playSoundAsyncNazgul()
	histoire := Red + Bold + "\n💀 Niveau Bonus ! Les 9 Nazgûls apparaissent pour venger Sauron ! Préparez-vous à un combat épique..." + Reset
	typeWriter(histoire, 7*time.Millisecond)

	nazgulsPv := 500
	nazgulsDegats := 35

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		if nazgulsPv <= 0 {
			fmt.Println(Green + Bold + "🎉 Victoire ultime ! Vous avez terrassé les 9 Nazgûls !" + Reset)
			player.Gold += 1000
			fmt.Println(Yellow + "💰 Récompense : 1000 pièces d’or !" + Reset)
			stopSound()
			playSoundAsyncDebut()
			return
		}

		if player.Pv <= 0 {
			player.IsDead()
			if player.Pv <= 0 {
				stopSound()
				fmt.Println(Red + Bold + "☠️ Les Nazgûls vous ont anéanti... les hobbits sont perdus pour de bon." + Reset)
				gameOver()
				return
			}
		}

		fmt.Println(Cyan + "\n=== Tour du joueur ===" + Reset)
		fmt.Printf("1 - %s (dégâts : %d)\n", player.BaseAttackName, player.BaseAttackDmg)
		fmt.Printf("2 - %s (dégâts : %d, coût : %d mana)\n", player.SkillName, player.SkillDmg, player.SkillManaCost)
		fmt.Printf("3 - Potion de vie (vie régénérée : 20, -1 potion de vie)\n")
		fmt.Printf("4 - Potion de poison (inflige : 10 dégâts, -1 potion de poison)\n")
		fmt.Printf("5 - Potion de mana (mana régénéré : 30, -1 potion de mana)\n")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		var degats int
		switch choix {
		case 1:
			degats = player.BaseAttackDmg
			nazgulsPv -= degats
			fmt.Printf("\n⚔️ Vous utilisez %s et infligez %d dégâts aux Nazgûls !\n", player.BaseAttackName, degats)
		case 2:
			if player.Mana >= player.SkillManaCost {
				degats = player.SkillDmg
				player.Mana -= player.SkillManaCost
				nazgulsPv -= degats
				fmt.Printf("\n🔥 Vous utilisez %s et infligez %d dégâts aux Nazgûls ! (Mana : %d/%d)\n", player.SkillName, degats, player.Mana, player.ManaMax)
			} else {
				fmt.Println(Red + "⚠️ Pas assez de mana ! Attaque normale." + Reset)
				degats = player.BaseAttackDmg
				nazgulsPv -= degats
				fmt.Printf("\n⚔️ Attaque de base : %d dégâts aux Nazgûls.\n", degats)
			}
		case 3:
			player.TakePotS()
		case 4:
			player.TakePotP(&nazgulsPv)
		case 5:
			player.TakePotM()
		default:
			fmt.Println(Red + "❌ Choix invalide, attaque de base par défaut !" + Reset)
			degats = player.BaseAttackDmg
			nazgulsPv -= degats
			fmt.Printf("\n⚔️ Attaque de base : %d dégâts aux Nazgûls.\n", degats)
		}

		if nazgulsPv < 0 {
			nazgulsPv = 0
		}
		fmt.Printf("👹 PV restants des Nazgûls : %d\n", nazgulsPv)

		fmt.Println(Red + "\n=== Tour des Nazgûls ===" + Reset)
		nazgulsAttack := rnd.Intn(3)
		var nazgulsDmg int
		switch nazgulsAttack {
		case 0:
			nazgulsDmg = nazgulsDegats
			player.Pv -= nazgulsDmg
			fmt.Printf(Red+Bold+"💀 Les Nazgûls frappent %s pour %d dégâts ! (PV : %d/%d)\n"+Reset,
				player.Nom, nazgulsDmg, player.Pv, player.PvMax)
		case 1:
			nazgulsDmg = nazgulsDegats + 20
			player.Pv -= nazgulsDmg
			fmt.Printf(Red+Bold+"💥 Les Nazgûls utilisent une attaque sombre et infligent %d dégâts ! (PV : %d/%d)\n"+Reset,
				nazgulsDmg, player.Pv, player.PvMax)
		case 2:
			nazgulsDmg = nazgulsDegats + 10
			player.Pv -= nazgulsDmg
			player.Mana -= 15
			if player.Mana < 0 {
				player.Mana = 0
			}
			fmt.Printf(Red+Bold+"🔥 Les Nazgûls invoquent l’ombre du Mordor et infligent %d dégâts ! Mana perdue : 15 (PV : %d/%d | Mana : %d/%d)\n"+Reset,
				nazgulsDmg, player.Pv, player.PvMax, player.Mana, player.ManaMax)
		}

		if player.Pv < 0 {
			player.Pv = 0
		}
	}
}
