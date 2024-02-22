package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	HOME string
)

func init() {
	HOME = os.Getenv("HOME")
}

func main() {
	clear()
	prompt()
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
