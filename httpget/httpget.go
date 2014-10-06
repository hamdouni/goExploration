package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://apple.com", nil)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		res, err := http.DefaultTransport.RoundTrip(req)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		} else {
			defer res.Body.Close()
			contents, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			} else {
				fmt.Printf("Content: %s\n", string(contents))
				fmt.Printf("Finale url: %s\n", res.Request.URL.String())
				fmt.Printf("Status: %s\n", res.Status)
			}
		}
	}
}
