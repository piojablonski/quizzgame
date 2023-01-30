package board_test

import (
	"bytes"
	"fmt"
	"github.com/piojablonski/quizzgame/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisplayQuestions(t *testing.T) {
	t.Run("asks one question and reads an answer and displays new question if it is correct", func(t *testing.T) {
		var out bytes.Buffer

		in := bytes.NewBufferString("4")
		board.DisplayQuestion(in, &out)

		got2 := out.String()
		want2 := "what is 2+2?4"
		assert.Equal(t, want2, got2, "want a question with answer %q, got %q", want2, got2)
	})

	t.Run("asks many questions and reads an answer for each", func(t *testing.T) {
		var out bytes.Buffer
		questions := board.Questions

		in := bytes.Buffer{}
		for _, q := range questions {
			fmt.Fprintln(&in, q.Answer)
		}
		board.DisplayQuestion(&in, &out)

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
}
