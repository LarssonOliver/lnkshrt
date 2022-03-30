package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	file   string
	sqlite *sql.DB
}

type rwMode string

const (
	// read   rwMode = "ro"
	// write  rwMode = "rw"
	create rwMode = "rwc"
)

const initStmt = `
CREATE TABLE IF NOT EXISTS links (
	id TEXT PRIMARY KEY,
	url TEXT NOT NULL UNIQUE,
	created_at TEXT NOT NULL
);
`

const insertStmt = `
INSERT INTO links (id, url, created_at) VALUES (?, ?, datetime('now'));
`

const deleteStmt = `
DELETE FROM links WHERE id = ?;
`

const getStmt = `
SELECT url FROM links WHERE id = ?;
`
const getIdStmt = `
SELECT id FROM links WHERE url = ?;
`

func (db *Database) genDSN(mode rwMode) string {
	return "file:" + db.file + "?cache=shared&mode=" + string(mode)
}

func New(file string) (*Database, error) {
	db := &Database{
		file: file,
	}

	// Create db file directory if it doesn't exist as sqlite doesn't
	dir := filepath.Dir(db.file)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("could not create directory %s: %w", dir, err)
	}

	conn, err := sql.Open("sqlite3", db.genDSN(create))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, db.genDSN(create))
	}

	db.sqlite = conn

	_, err = db.sqlite.Exec(initStmt)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, initStmt)
	}

	return db, nil
}

func (db *Database) Close() error {
	return db.sqlite.Close()
}

func (db *Database) Set(id string, url string) error {
	_, err := db.sqlite.Exec(insertStmt, id, url)
	return err
}

func (db *Database) Delete(id string) error {
	_, err := db.sqlite.Exec(deleteStmt, id)
	return err
}

func (db *Database) Get(id string) (string, error) {
	var url string
	err := db.sqlite.QueryRow(getStmt, id).Scan(&url)
	return url, err
}

func (db *Database) GetId(url string) (string, error) {
	var id string
	err := db.sqlite.QueryRow(getIdStmt, url).Scan(&id)
	return id, err
}
