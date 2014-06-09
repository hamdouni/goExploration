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

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	var (
		pv_id  int
		pv_nom string
	)
	rows, err := db.Query("select pv_id,pv_nom from pdv limit 10")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&pv_id, &pv_nom)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%d : %s \n", pv_id, pv_nom)
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
	http.HandleFunc("/service", serviceHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
