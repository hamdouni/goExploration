package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var arg string
	flag.StringVar(&arg, "date", "", "format YYYY-MM-DD")
	flag.Parse()
	if arg == "" {
		panic("argument date required")
	}
	d, e := time.Parse("2006-01-02", arg)
	if e != nil {
		panic(e)
	}

	fmt.Println(d)
}
