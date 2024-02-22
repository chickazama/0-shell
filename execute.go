package main

import (
	"fmt"
	"strings"
)

func execute(cmd string) bool {
	fields := strings.Fields(cmd)
	switch fields[0] {
	case "ls":
		ls(fields)
	case "cd":
		cd(fields)
	case "exit":
		return false
	case "clear":
		clear()
	default:
		fmt.Printf("'%s' command not implemented.\n", fields[0])
	}
	return true
}
