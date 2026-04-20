package main

import (
	"bufio"
	"errors"
	"io"
)

func untilEOF[T any](fetch func() (T, error), yield func(T) error) error {
	for {
		value, err := fetch()
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if e := yield(value); e != nil {
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

func eachLine(r io.Reader, yield func(string) error) error {
	br := bufio.NewReader(r)
	return untilEOF(
		func() (string, error) { return br.ReadString('\n') },
		yield)
}
