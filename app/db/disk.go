package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"larssonoliver.com/lnkshrt/app/config"
)

func (db *Database) loadFromDisk() {
	dbfile := config.DBFile()

	log.Println("Loading database from file:", dbfile)

	buf, err := ioutil.ReadFile(dbfile)

	if err != nil {
		log.Println("Error reading from file:", err)
		log.Println("Creating new database...")

		dir := filepath.Dir(dbfile)
		err = os.MkdirAll(dir, os.ModePerm)

		if err != nil {
			log.Println("Error creating directory:", err)
		}

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
