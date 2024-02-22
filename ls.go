package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func ls(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	dirent, err := fs.ReadDir(os.DirFS(dir), ".")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, ent := range dirent {
		fmt.Printf("%s  ", ent.Name())
	}
	fmt.Println()
}
