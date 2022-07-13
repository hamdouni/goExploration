package main

import (
	"database/sql"
	"fmt"
	"time"

	"log"

	_ "modernc.org/sqlite"
)

const dsnURI = "./toto.db"

func main() {
	db, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		log.Fatalf("error opening database %s got %s", dsnURI, err)
	}
	defer db.Close()

	log.Println("starting...")

	query := "INSERT INTO userinfo(username, departname, created) values(?,?,?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("error insert into database %s got %s", dsnURI, err)
	}

	res, err := stmt.Exec("barim", "si", "1972-06-02")
	if err != nil {
		log.Fatalf("error executing query %s got %s", query, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("error retrieving last id got %s", err)
	}
	log.Printf("last id %d", id)

	query = "SELECT * FROM userinfo"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("error executing query %s got %s", query, err)
	}

	var (
		uid        int
		username   string
		department string
		created    time.Time
	)

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			log.Printf("warning scanning results got %s", err)
			continue
		}
		fmt.Printf("uid: %d username: %s department: %s date: %v\n", uid, username, department, created)
	}

}
