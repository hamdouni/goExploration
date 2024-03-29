package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := make(handler)
	go compte(h)
	if err := http.ListenAndServe(":8000", h); err != nil {
		log.Print(err)
	}
}

func compte(ch chan<- int) {
	for n := 0; ; n++ {
		ch <- n
	}
}

type handler chan int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "%s: you are visitor #%d", req.URL, <-h)
}
