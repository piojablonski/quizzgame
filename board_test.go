package board_test

import (
	"bytes"
	board "github.com/piojablonski/quizzgame"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisplayQuestions(t *testing.T) {
	t.Run("asks one question", func(t *testing.T) {
		out := bytes.Buffer{}
		board.DisplayQuestion(&out)

		got := out.String()
		want := "what is 2+2?"

		assert.Equal(t, want, got, "want a question %q, got %q", want, got)
	})
}
