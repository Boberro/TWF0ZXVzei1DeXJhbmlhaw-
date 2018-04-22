package main

import (
	"github.com/boltdb/bolt"
	"log"
)

var db *bolt.DB

func dbConnect() {
	log.Println("Opening DB connection...")
	var err error
	db, err = bolt.Open("TWF0ZXVzei1DeXJhbmlhaw==.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dbClose() {
	log.Println("Closing DB connection...")
	db.Close()
}
