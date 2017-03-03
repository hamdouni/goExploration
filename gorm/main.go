package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	dbserver := "localhost:3306"
	gdb := new(gorm.DB)
	gdb, err := gorm.Open("mysql", "admin:admin@("+dbserver+")/rgc?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	err = gdb.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	gdb.LogMode(true)

}
