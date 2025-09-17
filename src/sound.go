package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var ctrl *beep.Ctrl // contrôle global de la musique en cours

// Initialise le speaker UNE SEULE FOIS
func initSpeaker() {
	// on initialise avec une valeur par défaut
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Millisecond*50))
}

// Lance une musique en boucle
func playMusic(path string) {
	stopSound() // coupe la musique précédente

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur ouverture fichier :", err)
		return
	}

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Erreur décodage mp3 :", err)
		return
	}

	// Crée un contrôleur pour pouvoir stopper la musique plus tard
	ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	// Joue la musique
	speaker.Play(ctrl)
}

// Stoppe la musique en cours
func stopSound() {
	if ctrl != nil {
		speaker.Lock()
		ctrl.Streamer = nil
		ctrl.Paused = true
		speaker.Unlock()
	}
}

// Fonctions spécifiques
func playSoundAsyncDebut() {
	playMusic("./assets/hobbit.mp3")
}

func playSoundAsyncCombat() {
	playMusic("./assets/Combat.mp3")
}
