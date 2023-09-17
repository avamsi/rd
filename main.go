package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/avamsi/ergo/assert"
)

// immediateDir returns the immediate directory of the given path -- i.e., the
// path itself if it's a directory, its parent directory otherwise.
func immediateDir(p string) (string, error) {
	s, err := os.Stat(p)
	if err != nil {
		return "", err
	}
	if s.IsDir() {
		return p, nil
	}
	return filepath.Dir(p), nil
}

func noSuchErr(p string) error {
	return fmt.Errorf("no such file or directory: %s", p)
}

func rd(p string) (string, error) {
	// If the input is an absolute path, don't check for relative paths up the
	// CWD -- convert the path to a directory (if needed) for cd though.
	if filepath.IsAbs(p) {
		d, err := immediateDir(p)
		if errors.Is(err, fs.ErrNotExist) {
			return "", noSuchErr(p)
		}
		return d, err
	}
	// For a relative path, check every directory starting with the CWD to the
	// root directory till a path is found (and covert said path to a directory
	// if needed, just like we do above).
	cwd := assert.Ok(os.Getwd())
	for next := filepath.Dir(cwd); cwd != next; {
		d, err := immediateDir(filepath.Join(cwd, p))
		if errors.Is(err, fs.ErrNotExist) {
			cwd, next = next, filepath.Dir(next)
			continue
		}
		return d, err
	}
	return "", noSuchErr(p)
}

func main() {
	switch args := os.Args[1:]; len(args) {
	case 0:
		// Do nothing.
	case 1:
		if d, err := rd(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "rd: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println(d)
		}
	default:
		fmt.Fprintln(os.Stderr, "rd: expected at most 1 argument, got", args)
		os.Exit(1)
	}
}
