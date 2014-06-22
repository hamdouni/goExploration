package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var allDirectories []string
	var index = 0

	switch l := len(os.Args); {
	case l == 2:
		allDirectories = append(allDirectories, os.Args[1])
	case l == 1:
		allDirectories = append(allDirectories, ".")
	default:
		println("usage: deepDir [directory]")
		os.Exit(1)
	}

	for {
		var dir = allDirectories[index]
		fmt.Println("\n\n" + dir)

		var allEntries, _ = ioutil.ReadDir(dir)
		for i := 0; i < len(allEntries); i++ {
			if allEntries[i].IsDir() {
				allDirectories = append(allDirectories, dir+"/"+allEntries[i].Name())
				fmt.Print(allEntries[i].Name() + "/ ")
			} else {
				fmt.Print(allEntries[i].Name() + " ")
			}
		}
		index++
		if index >= len(allDirectories) {
			break
		}
	}
}
