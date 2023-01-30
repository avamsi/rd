package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/avamsi/ergo"
)

func rd(p string) string {
	if filepath.IsAbs(p) {
		if s, err := os.Stat(p); err != nil && !s.IsDir() {
			return filepath.Dir(p)
		}
		return p
	}
	cwd := ergo.Must1(os.Getwd())
	for next := filepath.Dir(cwd); cwd != next; cwd, next = next, filepath.Dir(next) {
		q := filepath.Join(cwd, p)
		if s, err := os.Stat(q); errors.Is(err, os.ErrNotExist) {
			continue
		} else if s.IsDir() {
			return q
		} else {
			return filepath.Dir(q)
		}
	}
	return p
}

func main() {
	switch len(os.Args) {
	case 1:
		// Do nothing.
	case 2:
		fmt.Println(rd(os.Args[1]))
	default:
		fmt.Println("rd: expected at most 1 argument, got", os.Args[1:])
	}
}
