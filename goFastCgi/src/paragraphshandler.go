package main

import (
	"log"
	"database/sql"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createparagraphs"
)

var keywordsarr []string
var phrasesarr []string


func main() {
	
	locale :="fi_FI"
	themes :="finance"
	
	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlstr := "select keyword from keywords where locale='"+locale +"' and themes='"+themes+"'"
	
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
	log.Println(len(keywordsarr))
	
	sqlstr = "select phrase from phrases where locale='"+locale +"' and themes='"+themes+"'"
	
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
	log.Println(len(phrasesarr))
	
	createparagraphs.CreatePr(db,locale,themes,keywordsarr,phrasesarr,2) 

}