package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

var rxEmoji = regexp.MustCompile(
	"[\u2600-\u27BF\U0001F300-\U0001F9FF\U0001F1E6-\U0001F1FF]")

func unemoji(r io.Reader, name string) error {
	return eachLine(r, func(line string) error {
		_, err := io.WriteString(os.Stdout,
			rxEmoji.ReplaceAllStringFunc(
				line,
				func(s string) string {
					u, _ := utf8.DecodeRuneInString(s)
					return fmt.Sprintf("&#x%x;", u)
				}))
		return err
	})
}

func main() {
	runArgf(os.Args[1:], unemoji)
}
