package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Print("\033[H\033[2J")
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func TypeWriter(text string, delay time.Duration) {
	for _, r := range text {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
	fmt.Println()
}

func ShowCredits() {
	credits := []string{
		"🌟THE LORD OF RINGS GAMES 🌟",
		"",
		"Développement : " + Cyan + "\n \t - MANFREDONIA DIMITRI\n \t - TOLISANO MATTHIEU\n \t - ADIB HOUSSINE\n" + Reset,
		"Scénario : " + Green + "inspiration de l'univers de J.R.R. Tolkien " + Reset,
		"Code: " + Yellow + "Go Language " + Reset,

		"",
		"Merci d'avoir joué ! 🙏",
		"La Terre du Milieu est sauvée grâce à VOUS ! 💍🔥",
		"À bientôt pour de nouvelles aventures ! 🚀",
		"Une pensée pour nos mentors, sans qui la réalisation de ce projet n’aurait pas été possible.",
	}
	ClearTerminal()

	for _, line := range credits {
		typeWriter(line, 50*time.Millisecond)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println()
	fmt.Scanln()
}
