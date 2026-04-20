package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var flagInplace = flag.Bool("i", false, "edit files in place")

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

func (inplace *Inplace) unemoji(r io.Reader, name string) (_err error) {
	fmt.Fprintln(os.Stderr, name)
	var w io.Writer = os.Stdout
	if inplace.Flag && name != "" && name != "-" {
		fd, err := os.CreateTemp(filepath.Dir(name), filepath.Base(name))
		if err != nil {
			return err
		}
		w = fd
		defer func() {
			tmpName := fd.Name()
			err := fd.Close()
			if _err != nil {
				return
			}
			if err != nil {
				_err = err
			}
			inplace.Add(name, tmpName)
		}()
	}
	return eachLine(r, func(line string) error {
		_, err := io.WriteString(w,
			rxEmoji.ReplaceAllStringFunc(line, unemojiString))
		return err
	})
}

func mains() error {
	flag.Parse()
	inplace := &Inplace{Flag: *flagInplace}
	if err := parseArgf(flag.Args(), inplace.unemoji); err != nil {
		return err
	}
	return inplace.Commit()
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
