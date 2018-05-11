package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	exitonerror(err)

	res, err := http.DefaultTransport.RoundTrip(req)
	exitonerror(err)

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	exitonerror(err)

	fmt.Printf("Content: %s\nFinale url: %s\nStatus: %s\n", string(contents), res.Request.URL.String(), res.Status)
}
func exitonerror(e error) {
	if e != nil {
		log.Fatalf("%s\n", e)
	}
}
