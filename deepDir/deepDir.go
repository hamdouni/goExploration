// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// deepDir print all files and directories and sub-directories and sub-sub...
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	var dirs []string

	all := flag.Bool("a", false, "do not ignore entries starting with .")

	flag.Parse()

	log.Println(flag.Args())

	switch l := len(flag.Args()); {
	case l == 1:
		dirs = append(dirs, flag.Arg(0))
	case l == 0:
		dirs = append(dirs, ".")
	default:
		println("usage: deepDir [directory]")
		os.Exit(1)
	}

	for index := 0; index < len(dirs); index++ {
		var dir = dirs[index]
		fmt.Print("\n [" + dir + "] ")

		var entries, _ = ioutil.ReadDir(dir)
		for _, entry := range entries {
			var name = entry.Name()
			if *all || !strings.HasPrefix(name, ".") {
				if entry.IsDir() {
					dirs = append(dirs, dir+"/"+name)
					fmt.Print(name + "/ ")
				} else {
					fmt.Print(name + " ")
				}
			}
		}
	}
	fmt.Print("\n")
}
