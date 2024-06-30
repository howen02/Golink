package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitialiseDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./golink.db")
	if err != nil {
		log.Fatal(err)
	}

	createUrlsTableStmnt := `
		CREATE TABLE IF NOT EXISTS Urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			longUrl TEXT NOT NULL,
			shortUrl TEXT NOT NULL
		)
	`

	_, err = db.Exec(createUrlsTableStmnt)

	if err != nil {
		log.Fatal("Error creating Urls table: ", err)
	}

	log.Println("Urls table created")

	return db
}

func InsertLongUrl(db *sql.DB, longUrl string) string {
	shortUrl := generateShortUrl(longUrl)
	insertLongUrlStmnt := `
		INSERT INTO Urls (longUrl, shortUrl) VALUES (?, ?)
	`
	
	_, err := db.Exec(insertLongUrlStmnt, longUrl, shortUrl)
	
	if err != nil {
		log.Print("Error inserting longUrl into Urls table: ", err)
		return ""
	}

	log.Println("Inserted longUrl into Urls table")

	return shortUrl
}

func GetLongUrl(db *sql.DB, shortUrl string) string {
	longUrl := ""
	fetchLongUrlStmnt := `
		SELECT longUrl FROM URLS WHERE shortUrl = ? LIMIT 1
	`

	err := db.QueryRow(fetchLongUrlStmnt, shortUrl).Scan(&longUrl)

	if err != nil {
		log.Print("Error fetching longUrl from Urls table: ", err)
		return ""
	}

	log.Println("Fetched longUrl from Urls table")

	return longUrl
}

func generateShortUrl(longUrl string) string {
	hasher := sha256.New()
	hasher.Write([]byte(longUrl))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return hash[:8]
}