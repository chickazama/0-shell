package main

import (
	"bufio"
	"flag"
	"os"
)

const (
	newline = '\n'
)

var (
	HOME   string
	LS_CMD *flag.FlagSet
	LS_A   *bool
	LS_r   *bool
	reader *bufio.Reader
)

func init() {
	setLS_CMD()
	HOME = os.Getenv("HOME")
	reader = bufio.NewReader(os.Stdin)
}

func setLS_CMD() {
	LS_CMD = flag.NewFlagSet("ls", flag.ContinueOnError)
	LS_A = LS_CMD.Bool("a", false, "list all")
	LS_r = LS_CMD.Bool("r", false, "reverse")
}
