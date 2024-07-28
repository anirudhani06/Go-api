package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func PostgresStorage() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=secret sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Db connected succesfully")
	return &PostgresStore{db: db}, nil
}

func (p *PostgresStore) Init() error {
	return p.createUserTable()
}

func (p *PostgresStore) createUserTable() error {
	query := `create table if not exists "user" (
        id serial primary key,
        first_name varchar(40),
        last_name varchar(40),
        email varchar(255) unique,
		password varchar(255),
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp
    )`

	_, err := p.db.Exec(query)

	return err
}
