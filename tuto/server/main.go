// HTTP server demo
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleroot)
	http.HandleFunc("/bar", handlebar)
	http.HandleFunc("/foo", handlefoo)
	http.ListenAndServe(":8080", nil)
}

func handleroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v", r.URL.Path)
}
func handlebar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is bar")
}
func handlefoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is foo")
}
