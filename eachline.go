package main

import (
	"bufio"
	"errors"
	"io"
)

func eachLine(r io.Reader, yield func(string) error) error {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if e := yield(line); e != nil {
			return e
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
