package board_test

import (
	"bytes"
	"fmt"
	"github.com/piojablonski/quizzgame/board"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var questions = []board.Question{
	{"2+2", "4"},
	{"2+3", "5"},
	{"2+4", "6"},
}

const csvContent = `2+2,4
2+3,5
2+4,6
`

func TestInitialization(t *testing.T) {
	file := strings.NewReader(csvContent)
	b := board.New(file)
	assert.EqualValues(t, questions, b.Questions)
}

func assertOutputContains(t *testing.T, buffer fmt.Stringer, msgs ...string) {
	t.Helper()
	got := buffer.String()

	for _, m := range msgs {
		assert.Contains(t, got, m)
	}
}

func TestDisplayingQuestions(t *testing.T) {
	t.Run("contains start game prompt and shows the first question after user presses enter", func(t *testing.T) {
		b := board.Board{Questions: questions[:1]}

		var out bytes.Buffer

		in := bytes.NewBufferString("\n")
		err := b.DisplayQuestion(in, &out)
		assert.NoError(t, err)

		wantQuestion := "what is 2+2?"
		assertOutputContains(t, &out, board.WelcomePrompt, wantQuestion)

	})

	t.Run("asks one question and reads an answer and displays new question if it is correct", func(t *testing.T) {
		b := board.Board{Questions: questions[:1]}

		var out bytes.Buffer

		in := bytes.NewBufferString("\n4")
		err := b.DisplayQuestion(in, &out)
		assert.NoError(t, err)

		got := out.String()
		want := "what is 2+2?4"
		assert.Contains(t, got, want, "want a question with answer %q, got %q", want, got)
	})

	t.Run("asks many questions and reads an answer for each", func(t *testing.T) {
		b := board.Board{questions}

		var out bytes.Buffer
		questions := b.Questions

		in := bytes.Buffer{}
		fmt.Fprintln(&in, "")
		for _, q := range questions {
			fmt.Fprintln(&in, q.Answer)
		}
		err := b.DisplayQuestion(&in, &out)
		assert.NoError(t, err)

		for i, q := range questions {
			t.Run(fmt.Sprintf("question %d should contain q&a", i+1), func(t *testing.T) {
				assertOutputContains(t, &out, q.Question, q.Answer)
			})
		}
	})

	t.Run("asks many questions and display summary", func(t *testing.T) {
		b := board.Board{questions}

		var out bytes.Buffer

		in := bytes.Buffer{}
		fmt.Fprintln(&in, "")
		fmt.Fprintln(&in, questions[0].Answer)
		fmt.Fprintln(&in, "wrong or incorrect")
		fmt.Fprintln(&in, questions[2].Answer)
		err := b.DisplayQuestion(&in, &out)
		assert.NoError(t, err)

		want := "total questions: 3, correct answers: 2. Bravo!"

		assertOutputContains(t, &out, want)

	})
}
