package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var rxEmoji = regexp.MustCompile("[" +
	"\u2300-\u23FF" + // Miscellaneous Technical
	"\u2600-\u27BF" +
	"\uFE00-\uFE0F" + // Variation Selectors
	"\U0001F300-\U0001F9FF" +
	"\U0001F1E6-\U0001F1FF" +
	"]+")

func unemojiString(s string) string {
	var b strings.Builder
	for _, c := range s {
		fmt.Fprintf(&b, "&#x%X;", c)
	}
	return b.String()
}

func unemoji(r io.Reader, name string) error {
	return eachLine(r, func(line string) error {
		_, err := io.WriteString(os.Stdout,
			rxEmoji.ReplaceAllStringFunc(line, unemojiString))
		return err
	})
}

func main() {
	runArgf(os.Args[1:], unemoji)
}
