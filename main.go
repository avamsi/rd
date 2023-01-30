package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/avamsi/ergo"
)

func rd(p string) string {
	// If the input is an absolute path, don't check for relative paths up the
	// CWD -- convert the path to a directory (if needed) for cd though.
	if filepath.IsAbs(p) {
		if s, err := os.Stat(p); err != nil && !s.IsDir() {
			return filepath.Dir(p)
		}
		return p
	}
	// For a relative path, check every directory starting with the CWD to the
	// root directory till a path is found (and covert said path to a directory
	// if needed, just like we do above).
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
		fmt.Fprintln(os.Stderr, "rd: expected at most 1 argument, got", os.Args[1:])
	}
}
