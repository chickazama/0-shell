package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"slices"
	"strings"
)

func ls(fields []string) {
	cmd := flag.NewFlagSet("ls", flag.ContinueOnError)
	all := cmd.Bool("a", false, "include hidden files and directories")
	reverse := cmd.Bool("r", false, "reverse listing")
	recursive := cmd.Bool("R", false, "recursive listing")
	if len(fields) > 1 {
		err := cmd.Parse(fields[1:])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	argv := cmd.Args()
	var fsys fs.FS
	switch len(argv) {
	case 0:
		fsys = os.DirFS("./")
	case 1:
		fsys = os.DirFS(argv[0])
	}
	if *recursive {
		var sb strings.Builder
		start := "."
		var output []string
		err := fs.WalkDir(fsys, start, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				if !(strings.HasPrefix(path, ".") && !*all && path != ".") {
					add := fmt.Sprintf("%s:\n", path)
					sb.WriteString(add)
					entries, err := fs.ReadDir(fsys, path)
					if err != nil {
						log.Fatal(err.Error())
					}
					if *reverse {
						slices.Reverse(entries)
					}
					for _, e := range entries {
						name := e.Name()
						if !(strings.HasPrefix(name, ".") && !*all) {
							add = fmt.Sprintf("%s  ", name)
							sb.WriteString(add)
						}
					}
					output = append(output, sb.String())
					sb.Reset()
				}
			}
			return nil
		})
		if *reverse {
			slices.Reverse(output)
		}
		for _, s := range output {
			fmt.Println(s)
		}
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		entries, err := fs.ReadDir(fsys, ".")
		if err != nil {
			fmt.Printf("ls: unable to list contents of '%s'\n", argv[0])
			return
		}
		if *reverse {
			slices.Reverse(entries)
		}
		for _, e := range entries {
			name := e.Name()
			if !(strings.HasPrefix(name, ".") && !*all) {
				fmt.Printf("%s  ", name)
			}
		}
		fmt.Println()
	}
	*all = false
	*reverse = false
	*recursive = false
}
