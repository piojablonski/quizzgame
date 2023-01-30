package board

import (
	"fmt"
	"io"
)

func DisplayQuestion(out io.Writer) {
	fmt.Fprint(out, "what is 2+2?")
}
