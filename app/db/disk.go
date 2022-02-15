package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"larssonoliver.com/lnkshrt/app/config"
)

func (db *Database) loadFromDisk() {
	log.Println("Loading database from file:", config.DBFile())

	buf, err := ioutil.ReadFile(config.DBFile())

	if err != nil {
		log.Println("Error reading from file:", err)
		log.Println("Creating new database...")
		return
	}

	db.lock.Lock()
	err = json.Unmarshal(buf, &db.links)
	db.lock.Unlock()

	if err != nil {
		log.Println("Error deserializing data:", err)
		return
	}

	log.Println("Database loaded")
}

func (db *Database) writeToDisk() {
	db.lock.Lock()
	buf, err := json.Marshal(db.links)
	db.lock.Unlock()

	if err != nil {
		log.Println("Error serializing data:", err)
		return
	}

	err = ioutil.WriteFile(config.DBFile(), buf, os.ModePerm)

	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}
