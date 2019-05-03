// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// serv is a very small example of an http server connected to a mysql database
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	port  int
	db, _ = sql.Open("mysql", "root:@/prod")
)

func browser(url string) error {
	var commands = map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	cmd := exec.Command(run, url)
	return cmd.Start()
}

func init() {
	flag.IntVar(&port, "port", 8080, "port to run the server")
	flag.Parse()
	browser("http://0.0.0.0:" + strconv.Itoa(port) + "/index.html")
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
	if filename == "" {
		filename = "index.html"
	}
	body, mime, err := loadPage(filename)
	if err != nil {
		log.Printf("Get error : %v \n", err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", mime)
	fmt.Fprintf(w, "%s", body)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var (
		uId       int
		uLogin    string
		uNom      string
		uPrenom   string
		uRole     int
		pdvLabel  string
		uSupprime bool
	)
	rows, err := db.Query(
		`SELECT u_id,u_login,u_nom,u_prenom,u_role,concat(u_pdv,' - ',pv_nom, ' - ',a_ville),u_supprime 
		 FROM utilisateur LEFT JOIN pdv ON pv_id=u_pdv LEFT JOIN adresse ON a_id=pv_adr_pdv 
		 WHERE u_nom LIKE ? or u_login LIKE ?`,
		"test",
		"test")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&uId, &uLogin, &uNom, &uPrenom, &uRole, &pdvLabel, &uSupprime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%d : %s %s %s %d %s %v \n", uId, uLogin, uNom, uPrenom, uRole, pdvLabel, uSupprime)
	}
}

func main() {
	print("Connecting to database? ")

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	println("OK")

	println("Server start on port ", port)
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/get/users", getUsers)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
