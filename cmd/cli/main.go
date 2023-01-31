package main

import (
	"flag"
	"github.com/piojablonski/quizzgame/board"
	"log"
	"os"
)

var filename = flag.String("f", "problem.csv", `gives a path to a file with questions, defaults to "problem.csv"`)

func main() {
	flag.Parse()
	file, err := os.OpenFile(*filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("problem opening file %s, %v", *filename, err)
	}
	b := board.New(file)
	err = b.DisplayQuestion(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatalf("problem asking questions, %v", err)
	}
}
