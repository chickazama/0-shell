package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	HOME string
)

func init() {
	HOME = os.Getenv("HOME")
}

func main() {
	prompt()
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
