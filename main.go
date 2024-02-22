package main

import (
	"log"
)

func main() {
	clear()
	loop := true
	for loop {
		prompt()
		cmd, err := reader.ReadString(newline)
		if err != nil {
			log.Fatal(err.Error())
		}
		loop = execute(cmd)
	}
}
