package main

import (
	"fmt"
	"strings"

	"github.com/mywrap/textproc"
)

func main() {
	// see `text_test.go` and `html_test.go` for example

	for i := 0; i < 10; i++ {
		fmt.Println(strings.ToLower(textproc.GenRandomWord(
			8, 12, textproc.AlphaEnList)))
	}
}
