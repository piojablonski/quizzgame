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

func TestDisplayingQuestions(t *testing.T) {
	t.Run("asks one question and reads an answer and displays new question if it is correct", func(t *testing.T) {
		b := board.Board{Questions: questions[:1]}

		var out bytes.Buffer

		in := bytes.NewBufferString("4")
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
		for _, q := range questions {
			fmt.Fprintln(&in, q.Answer)
		}
		err := b.DisplayQuestion(&in, &out)
		assert.NoError(t, err)

		for i, q := range questions {
			t.Run(fmt.Sprintf("question %d should contain q&a", i+1), func(t *testing.T) {
				got := out.String()

				wantQuestion := q.Question
				assert.Contains(t, got, wantQuestion, "output %q should contain %q", got, wantQuestion)

				wantAnswer := q.Answer
				assert.Contains(t, got, wantAnswer, "output %q should contain %q", got, wantQuestion)

			})
		}
	})

	t.Run("asks many questions and display summary", func(t *testing.T) {
		b := board.Board{questions}

		var out bytes.Buffer

		in := bytes.Buffer{}
		fmt.Fprintln(&in, questions[0].Answer)
		fmt.Fprintln(&in, "wrong or incorrect")
		fmt.Fprintln(&in, questions[2].Answer)
		err := b.DisplayQuestion(&in, &out)
		assert.NoError(t, err)

		got := out.String()
		want := "total questions: 3, correct answers: 2. Bravo!"

		assert.Contains(t, got, want, "it should contain summary with propoer counts")

	})
}
