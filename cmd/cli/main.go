package main

import (
	board "github.com/piojablonski/quizzgame"
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
