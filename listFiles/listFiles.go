// listFiles print only files in the specified directory or the current directory if no parameter is passed
package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var dirs string

	flag.Parse()

	switch l := len(flag.Args()); {
	case l == 0:
		dirs = "."
	case l == 1:
		dirs = flag.Arg(0)
	default:
		log.Fatal("usage: listFiles [dir]")
	}
	files, err := ioutil.ReadDir(dirs)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			println(file.Name())
		}
	}
}
