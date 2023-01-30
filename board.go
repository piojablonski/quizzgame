package board

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

var Questions = []struct {
	Question string
	Answer   string
}{
	{"2+2", "4"},
	{"2+3", "5"},
	{"2+4", "6"},
}

func DisplayQuestion(in io.Reader, out io.Writer) {
	var answer string
	var scanner = bufio.NewScanner(in)
	fmt.Fprint(out, "what is 2+2?")
	answer = waitForMessage(scanner, answer)
	log.Printf("read scanned text: %q\n", answer)
	fmt.Fprint(out, answer)
}

func waitForMessage(scanner *bufio.Scanner, answer string) string {
	if scanner.Scan() {
		answer = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
	return answer
}
