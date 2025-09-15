package main

import (
	"fmt"
	"time"
)

const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Pink      = "\033[35m"
)

func typeWriter(str string, delay time.Duration) {
	for _, r := range str {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
	fmt.Println()
}
