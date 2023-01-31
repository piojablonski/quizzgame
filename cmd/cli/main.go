package main

import (
	"github.com/piojablonski/quizzgame/board"
	"log"
	"os"
)

func main() {
	const filename = "problems.csv"
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("problem opening file %s, %v", filename, err)
	}
	b := board.New(file)
	err = b.DisplayQuestion(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatalf("problem asking questions, %v", err)
	}
}
