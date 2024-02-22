package main

import (
	"bufio"
	"os"
)

const (
	newline = '\n'
)

var (
	HOME   string
	reader *bufio.Reader
)

func init() {
	HOME = os.Getenv("HOME")
	reader = bufio.NewReader(os.Stdin)
}
