package main

import (
	"fmt"
	"goExploration/sqlite/repo/db"
	"goExploration/sqlite/task"
	"log"

	_ "modernc.org/sqlite"
)

const dbpath = "./toto.db"
const pragma = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"

func main() {

	log.Println("starting...")

	store, err := db.Open(dbpath, pragma)
	if err != nil {
		log.Fatalf("could not open database %s: %s", dbpath, err)
	}
	defer store.Close()

	task.Init(&store)

	id, err := task.Create("faire un café")
	if err != nil {
		log.Fatalf("could not create task %s", err)
	}
	log.Printf("last id %d", id)

	items := store.GetAll()
	for _, item := range items {
		fmt.Printf("uid: %d description: %s state: %d\n", item.ID, item.Description, item.State)
	}

}
