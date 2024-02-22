package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func execute(cmd string) bool {
	fields := strings.Fields(cmd)
	switch fields[0] {
	case "ls":
		ls(fields[1:])
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
	case "clear":
		clear()
	default:
		fmt.Printf("'%s' command not implemented.\n", fields[0])
		// cmd := exec.Command(fields[0], fields[1:]...)
		// cmd.Stdout = os.Stdout
		// err := cmd.Run()
		// if err != nil {
		// 	log.Println(err.Error())
		// }
	}
	return true
}
