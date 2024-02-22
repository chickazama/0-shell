package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
