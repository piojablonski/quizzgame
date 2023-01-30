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
	for _, q := range Questions {
		fmt.Fprintf(out, "what is %s?", q.Question)
		answer = waitForMessage(scanner, answer)
		fmt.Fprintln(out, answer)
	}
}

func waitForMessage(scanner *bufio.Scanner, answer string) string {
	if scanner.Scan() {
		answer = scanner.Text()
		log.Printf("read scanned text: %q\n", answer)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
	return answer
}
