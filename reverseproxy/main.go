package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	rpURL, err := url.Parse("http://localhost:8002")
	if err != nil {
		log.Fatal(err)
	}
	proxyGO := httputil.NewSingleHostReverseProxy(rpURL)
	http.HandleFunc("/pdv/ext-users/api/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/pdv/ext-users")
		proxyGO.ServeHTTP(w, r)
	})

	rpURL, err = url.Parse("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	proxyPDV := httputil.NewSingleHostReverseProxy(rpURL)
	http.HandleFunc("/pdv/", func(w http.ResponseWriter, r *http.Request) {
		proxyPDV.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":9001", nil))
}
