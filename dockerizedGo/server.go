// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// serv is a very small example of an http server connected to a mysql database
package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)

var (
	port  int
)

func init() {
	flag.IntVar(&port, "port", 8080, "port to run the server")
	flag.Parse()
}

func loadPage(filename string) ([]byte, string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}
	ext := filepath.Ext(filename)
	mime := mime.TypeByExtension(ext)
	return body, mime, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/"):]
	body, mime, _ := loadPage(filename)
	w.Header().Set("Content-Type", mime)
	fmt.Fprintf(w, "%s", body)
}

func main() {
	println("Server start on port ", port)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
