package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb(name string) (*sql.DB, error) {

	cwd, err := os.Getwd()
	path := filepath.Dir(cwd)

	addr := filepath.Join(path, name)
	if _, err := os.Stat(addr); os.IsNotExist(err) {
		file, err := os.Create(addr)
		if err != nil {
			fmt.Println("Error creating database file:", err)
		}

		file.Close()
	}

	db, err := sql.Open("sqlite3", addr)

	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
