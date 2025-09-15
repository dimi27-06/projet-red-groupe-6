package main

import (
	"fmt"
	"time"
)

func TypeWriter(text string, delay time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println()
}
