package database

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("url not found")

type DB struct{
	sqlDB *sql.DB
}

func New(sqlDb *sql.DB) *DB {
	return &DB {
		sqlDb,
	}
}

func (db *DB) GetOriginal(tiny string) (string, error) {
	row := db.sqlDB.QueryRow("SELECT original FROM urls WHERE tiny=?", tiny)
	var originalUrl string
	err := row.Scan(&originalUrl)
	if err == sql.ErrNoRows {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func (db *DB) Set(originalUrl string, tiny string) error {
	_, err := db.sqlDB.Exec("INSERT INTO urls (original, tiny) VALUES (?, ?)", originalUrl, tiny)
	return err
}