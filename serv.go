// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// serv is a very small example of an http server connected to a mysql database
package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
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
	body, mime, _ := loadPage(filename)
	w.Header().Set("Content-Type", mime)
	fmt.Fprintf(w, "%s", body)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var (
		u_id       int
		u_login    string
		u_nom      string
		u_prenom   string
		u_role     int
		pdv_label  string
		u_supprime bool
	)
	rows, err := db.Query(
		`SELECT u_id,u_login,u_nom,u_prenom,u_role,concat(u_pdv,' - ',pv_nom, ' - ',a_ville),u_supprime 
		 FROM utilisateur LEFT JOIN pdv ON pv_id=u_pdv LEFT JOIN adresse ON a_id=pv_adr_pdv 
		 WHERE u_nom LIKE ? or u_login LIKE ?`,
		"hamdouni",
		"hamdouni")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&u_id, &u_login, &u_nom, &u_prenom, &u_role, &pdv_label, &u_supprime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%d : %s %s %s %d %s %d \n", u_id, u_login, u_nom, u_prenom, u_role, pdv_label, u_supprime)
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
