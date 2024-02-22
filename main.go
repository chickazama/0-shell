package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

func main() {
	clear()
	for {
		prompt()
		str, err := reader.ReadString(newline)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Print(str)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func prompt() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	if strings.HasPrefix(dir, HOME) {
		dir = strings.TrimPrefix(dir, HOME)
		dir = fmt.Sprintf("~%s", dir)
	}
	fmt.Printf("%s$ ", dir)
}
