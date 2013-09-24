package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createparagraphs"
	"database/sql"
	"log"
)

var keywordsarr []string
var phrasesarr []string
var hostsarr []string

func main() {

	locale := "fi_FI"
	themes := "finance"

	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlstr := "select keyword from keywords where locale='" + locale + "' and themes='" + themes + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var keyword string
		rows.Scan(&keyword)
		keywordsarr = append(keywordsarr, keyword)

	}
	rows.Close()
	log.Println("keywords", len(keywordsarr))

	sqlstr = "select phrase from phrases where locale='" + locale + "' and themes='" + themes + "'"

	rows, err = db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var phrase string
		rows.Scan(&phrase)
		phrasesarr = append(phrasesarr, phrase)

	}
	rows.Close()
	log.Println("Phrases", len(phrasesarr))

	sqlstr = "select host from hosts where locale='" + locale + "' and themes='" + themes + "'"

	rows, err = db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var host string
		rows.Scan(&host)
		hostsarr = append(hostsarr, host)

	}
	rows.Close()
	log.Println("hosts", len(hostsarr))

	if db.Close(); err != nil {

		log.Fatal(err)
		
	} else {

		createparagraphs.CreatePr(locale, themes, keywordsarr, phrasesarr, hostsarr, 50)

	}

}
