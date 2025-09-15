package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Joue le son sans bloquer le programme
func playSoundAsync() {
	f, err := os.Open("./assets/hobbit.mp3")
	if err != nil {
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Millisecond*50))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		// Nettoyage à la fin du son
		streamer.Close()
		f.Close()
	})))
}

// Affiche un texte progressivement, un caractère à la fois
func printSlowly(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}
