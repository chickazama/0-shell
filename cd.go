package main

import (
	"log"
	"os"
)

func cd(args []string) {
	if len(args) > 1 {
		err := os.Chdir(args[1])
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		err := os.Chdir(HOME)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
