package main

import (
	"fmt"
	"time"
)

// Structure Personnage
type Character struct {
	Nom         string
	Classe      string
	Niveau      int
	Pv          int
	PvMax       int
	Inventaire  []Item
	Description string
	FreePotion  bool
	Gold        int
}

// Structure Item
type Item struct {
	Nom      string
	Quantite int
}

// Initialisation d'Aragorn
func (player *Character) initAragorn() {
	*player = Character{
		Nom:    "Aragorn",
		Classe: "Humain",
		Niveau: 1,
		Pv:     100,
		PvMax:  100,
		Inventaire: []Item{
			{"Épée Andúril", 1},
			{"Potion de vie", 0},
		},
		FreePotion: true,
		Gold:       0,

		Description: `
Aragorn, également connu sous le nom de Grands-pas,  
Héritier d'Isildur, il est le descendant direct des rois de Gondor. 
Pendant longtemps, il a erré à travers la Terre du Milieu, protégeant les peuples libres sous le nom de rôdeur. 
Il est un expert en traque, en combat et possède une grande connaissance des arts anciens.

Il manie l'épée légendaire Andúril, reforgée à partir des fragments de Narsil, l'épée d' Elendil. 
Aragorn est un leader courageux, doté d'une grande sagesse et d'une force morale exceptionnelle. 
Il est également le compagnon fidèle de Frodon et de la Communauté de l'Anneau, jouant un rôle crucial dans la défaite de Sauron.`,
	}
}

// Initialisation de Legolas
func (player *Character) initLegolas() {
	*player = Character{
		Nom:    "Legolas",
		Classe: "Elfe",
		Niveau: 1,
		Pv:     80,
		PvMax:  80,
		Inventaire: []Item{
			{"Arc elfique", 1},
			{"Potion de vie", 0},
		},
		FreePotion: true,
		Gold:       0,
		Description: `Legolas - Le Prince des Elfes Sylvains

Legolas est un Elfe sylvain du royaume de la Forêt Noire, fils du roi Thranduil. 
Membre de la Communauté de l’Anneau, il est reconnu pour son agilité, sa rapidité et sa maîtrise exceptionnelle de l’arc. 
Doté d’une vue perçante et d’une ouïe fine, il peut repérer ses ennemis à de grandes distances et ne manque presque jamais sa cible.

Toujours calme et déterminé, Legolas se distingue aussi par son amitié indéfectible avec Gimli, le nain, démontrant ainsi que les vieilles rivalités entre leurs peuples peuvent être surmontées. 
Il est un combattant hors pair, agile aussi bien sur terre que dans les environnements difficiles, et joue un rôle crucial dans les batailles de la guerre de l’Anneau.`,
	}
}

// Initialisation de Gimli
func (player *Character) initGimli() {
	*player = Character{
		Nom:    "Gimli",
		Classe: "Nain",
		Niveau: 1,
		Pv:     120,
		PvMax:  120,
		Inventaire: []Item{
			{"Hache de guerre", 1},

			{"Potion de vie", 0},
		},
		FreePotion: true,
		Gold:       0,
		Description: `Gimli - Le Guerrier Nain de la Montagne

Gimli est un nain du peuple de Durin, fils de Glóin, et l’un des membres de la Communauté de l’Anneau. 
Courageux, robuste et déterminé, il manie la hache avec une résistance redoutable, et son endurance dépasse celle de nombreux guerriers. 
Bien que méfiant envers les elfes au départ, il développe une amitié indestructible avec Legolas, prouvant que les vieilles querelles entre leurs peuples peuvent être surmontées.

Sa bravoure, son humour franc et sa ténacité font de lui un allié indispensable dans les batailles contre les forces de Sauron. 
Gimli représente la fierté et la force du peuple nain, toujours prêt à défendre ses compagnons et à plonger au cœur du danger.

`,
	}
}

// Affichage infos personnage
func (player Character) displayInfo() {
	fmt.Println("=== Information du personnage ===")
	fmt.Printf("\t - Nom : %s\n", player.Nom)
	fmt.Printf("\t - Classe : %s\n", player.Classe)
	fmt.Printf("\t - Niveau : %d\n", player.Niveau)
	fmt.Printf("\t - Pv : %d/%d\n", player.Pv, player.PvMax)
	fmt.Println("\t - Inventaire :")
	for _, item := range player.Inventaire {
		fmt.Printf("\t   * %s (x%d)\n", item.Nom, item.Quantite)
	}

	fmt.Print("\t - Description : ")
	TypeWriter(player.Description, 4*time.Millisecond)
}

// Fonction de sélection du personnage
func choisirPersonnage() Character {
	var choix int
	var perso Character

	fmt.Println("Choisissez votre personnage :")
	fmt.Println("1 - Aragorn")
	fmt.Println("2 - Legolas")
	fmt.Println("3 - Gimli")
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		perso.initAragorn()
	case 2:
		perso.initLegolas()
	case 3:
		perso.initGimli()
	default:
		fmt.Println("Choix invalide, Aragorn est sélectionné par défaut.")
		perso.initAragorn()
	}
	return perso
}
