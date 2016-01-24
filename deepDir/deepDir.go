// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// deepDir print all files and directories and sub-directories and sub-sub...
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var allDir []string

	switch l := len(os.Args); {
	case l == 2:
		allDir = append(allDir, os.Args[1])
	case l == 1:
		allDir = append(allDir, ".")
	default:
		println("usage: deepDir [directory]")
		os.Exit(1)
	}

	for index := 0; index < len(allDir); index++ {
		var dir = allDir[index]
		fmt.Print("\n [" + dir + "] ")

		var allEntries, _ = ioutil.ReadDir(dir)
		for _, entry := range allEntries {
			var name = entry.Name()
			if entry.IsDir() {
				allDir = append(allDir, dir+"/"+name)
				fmt.Print(name + "/ ")
			} else {
				fmt.Print(name + " ")
			}
		}
	}
	fmt.Print("\n")
}
