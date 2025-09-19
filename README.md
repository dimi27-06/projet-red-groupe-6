# projet-red-groupe-6
# projet-red-groupe-6
#  The Lord of Rings Game 

Un jeu d’aventure inspiré de l’univers de **J.R.R. Tolkien**, développé en **Go**
Les joueurs choisissent un personnage, gerent leur inventaire, combattent des ennemis et progressent dans lhistoire pour sauver les hobbits et affronter Sauron

---

##  Fonctionnaliter

- Choix du personnage : **legolas , gimli , aragorn**
- Menu principal interactif
- Gestion de l’inventaire avec ajout suppression et limites
- Shop pour acheter potions et équipement
- Combats :
  - Orcs
  - Trolls
  - Sauron (combat final)
  - Nazgûls (niveau bonus)
- Gestion des points de vie et résurrection par Gandalf
- Effets visuels dans le terminal (ASCII art, couleurs, typewriter)
- Musiques d’ambiance pour le début le combat, Sauron et les crédits

---

##  Installation et exécution

1 Cloner le dépôt :

```bash
git clone <lien_du_dépôt>
cd projet-red

    Installer Go (version 1.25 ou supérieure) et les dépendances :

go mod tidy

    Exécuter le jeu :

go run main.go

2   Assurez-vous que les fichiers MP3 pour les musiques sont dans le dossier ./assets/.


 Musiques

Début : hobbit.mp3

Combat : Combat.mp3

Sauron : Sauron.mp3

Crédits : CREDITS.mp3

Nazgûls : nazgul.mp3



Couleurs et effets dans le terminal

Rouge : \033[31m

Vert : \033[32m

Jaune : \033[33m

Bleu : \033[34m

Cyan : \033[36m

Gras : \033[1m

Souligné : \033[4m

Typewriter pour un effet de texte animé


