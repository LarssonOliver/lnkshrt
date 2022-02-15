package db

import (
	"sync"
)

type Database struct {
	lock  *sync.Mutex
	links map[string]string
}

func New() *Database {
	return &Database{
		lock:  &sync.Mutex{},
		links: make(map[string]string),
	}
}

func (db *Database) Set(id string, url string) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.links[id] = url
}

func (db *Database) Delete(id string) bool {
	db.lock.Lock()
	defer db.lock.Unlock()
	if _, found := db.links[id]; found {
		delete(db.links, id)
		return true
	}
	return false
}

func (db *Database) Get(id string) (string, bool) {
	db.lock.Lock()
	defer db.lock.Unlock()
	url, found := db.links[id]
	return url, found
}
