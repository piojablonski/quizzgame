package main

import (
	"github.com/piojablonski/quizzgame/board"
	"os"
)

func main() {
	done := make(chan bool)
	go func() {
		board.DisplayQuestion(os.Stdin, os.Stdout)
		done <- true
	}()

	<-done
}
