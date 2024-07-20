package internal

import (
	"io"
	"os"
)

// default os.Stdout
var writer io.Writer = io.Writer(os.Stdout)

func SetWriter(w io.Writer) {
	writer = w
}

func Show(ch chan string) {
	for s := range ch {
		_, _ = writer.Write([]byte(s))
	}
}
