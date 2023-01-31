package board

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

type Question struct {
	Question string
	Answer   string
}

type Board struct {
	Questions []Question
}

func New(file io.Reader) *Board {
	b := Board{}
	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("problem parsing csv, %v", err)
		}
		b.Questions = append(b.Questions, Question{record[0], record[1]})
	}
	return &b
}

const (
	WelcomePrompt = "Click enter to start a quizz:\n\n"
)

func (b *Board) DisplayQuestion(in io.Reader, out io.Writer) error {
	var answer string
	var scanner = bufio.NewScanner(in)
	correctAnswers := 0
	fmt.Fprint(out, WelcomePrompt)
	waitForMessage(scanner, answer)
	for _, q := range b.Questions {
		if _, err := fmt.Fprintf(out, "what is %s?", q.Question); err != nil {
			return err
		}
		answer = waitForMessage(scanner, answer)

		if answer == q.Answer {
			correctAnswers++
		}

		if _, err := fmt.Fprintln(out, answer); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(out, "total questions: %d, correct answers: %d. Bravo!", len(b.Questions), correctAnswers); err != nil {
		return err
	}
	return nil
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
