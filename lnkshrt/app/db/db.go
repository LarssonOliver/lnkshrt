package db

import (
	"sync"

	"larssonoliver.com/lnkshrt/app/config"
)

type Database struct {
	lock  *sync.Mutex
	links map[string]string
}

func New() *Database {
	db := &Database{
		lock:  &sync.Mutex{},
		links: make(map[string]string),
	}

	if config.Persistent() {
		db.loadFromDisk()
	}

	return db
}

func (db *Database) Set(id string, url string) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.links[id] = url

	if config.Persistent() {
		go db.writeToDisk()
	}
}

func (db *Database) Delete(id string) bool {
	db.lock.Lock()
	defer db.lock.Unlock()

	_, found := db.links[id]
	if found {
		delete(db.links, id)

		if config.Persistent() {
			go db.writeToDisk()
		}
	}

	return found
}

func (db *Database) Get(id string) (string, bool) {
	db.lock.Lock()
	defer db.lock.Unlock()
	url, found := db.links[id]
	return url, found
}
