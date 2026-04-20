package main

import (
	"os"
)

type pair struct {
	tmp string
	org string
}

type Inplace struct {
	Flag bool
	list []pair
}

func (inplace *Inplace) Add(name, tmp string) {
	inplace.list = append(inplace.list, pair{
		org: name,
		tmp: tmp,
	})
}

func (inplace *Inplace) Commit() error {
	for _, p := range inplace.list {
		if err := os.Rename(p.org, p.org+"~"); err != nil {
			return err
		}
		if err := os.Rename(p.tmp, p.org); err != nil {
			return err
		}
	}
	return nil
}
