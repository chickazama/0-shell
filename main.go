package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	clear()
	loop := true
	for loop {
		prompt()
		str, err := reader.ReadString(newline)
		if err != nil {
			log.Fatal(err.Error())
		}
		loop = parse(str)
	}
}

func parse(str string) bool {

	fields := strings.Fields(str)
	switch fields[0] {
	case "cd":
		if len(fields) > 1 {
			err := os.Chdir(fields[1])
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			err := os.Chdir(HOME)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	case "exit":
		return false
	default:
		cmd := exec.Command(fields[0], fields[1:]...)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println(err.Error())
		}
	}
	return true
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
		dir = fmt.Sprintf("\033[0;32m~%s", dir)
	}
	fmt.Printf("%s$ ", dir)
}
