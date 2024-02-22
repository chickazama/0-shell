package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

var (
	LS_CMD = flag.NewFlagSet("ls", flag.ContinueOnError)
	LS_A   = LS_CMD.Bool("a", false, "include hidden files and directories")
)

func ls(args []string) {
	if len(args) > 1 {
		err := LS_CMD.Parse(args[1:])
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// fmt.Println(*LS_A)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	dirent, err := fs.ReadDir(os.DirFS(dir), ".")
	if err != nil {
		log.Fatal(err.Error())
	}
	if *LS_A {
		fmt.Print(".  ..  ")
	}
	for _, ent := range dirent {
		name := ent.Name()
		if !(strings.HasPrefix(name, ".") && !*LS_A) {
			fmt.Printf("%s  ", ent.Name())
		}
	}
	*LS_A = false
	fmt.Println()
}
