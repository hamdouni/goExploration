package main

import (
	"flag"
	"fmt"
	"time"
)

const dateformat = "2006-01-02"

func main() {
	var date string
	flag.StringVar(&date, "date", time.Now().Format(dateformat), "format YYYY-MM-DD")
	flag.Parse()

	if len(date) == 0 {
		fmt.Println(time.Now())
		return
	}

	d, err := time.Parse(dateformat, date)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}
