package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func parseArgf(args []string, routine func(io.Reader, string) error) error {
	if len(args) <= 0 {
		return routine(os.Stdin, "")
	}
	for _, arg := range args {
		filenames, err := filepath.Glob(arg)
		if err != nil || len(filenames) <= 0 {
			filenames = []string{arg}
		}
		for _, fname1 := range filenames {
			if fname1 == "-" {
				if err := routine(os.Stdin, "-"); err != nil {
					return err
				}
				continue
			}
			fd, err := os.Open(fname1)
			if err != nil {
				return err
			}
			err = routine(fd, fname1)
			err1 := fd.Close()
			if err != nil {
				return err
			}
			if err1 != nil {
				return err1
			}
		}
	}
	return nil
}

func runArgf(args []string, routine func(io.Reader, string) error) {
	if err := parseArgf(os.Args[1:], routine); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
