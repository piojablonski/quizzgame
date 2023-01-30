package board_test

import (
	"bytes"
	board "github.com/piojablonski/quizzgame"
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
}
