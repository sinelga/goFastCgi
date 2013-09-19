package main

import (
	"log"
	"database/sql"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createparagraphs"
)

var keywordsarr_fi_FI_finance []string
var phrasesarr_fi_FI_finance []string


func main() {

	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.Query("select keyword from keywords where locale='fi_FI' and themes='finance'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var keyword string
		rows.Scan(&keyword)
		keywordsarr_fi_FI_finance = append(keywordsarr_fi_FI_finance, keyword)

	}
	rows.Close()
	log.Println(len(keywordsarr_fi_FI_finance))

	rows, err = db.Query("select phrase from phrases where locale='fi_FI' and themes='finance'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var phrase string
		rows.Scan(&phrase)
		phrasesarr_fi_FI_finance = append(phrasesarr_fi_FI_finance, phrase)

	}
	rows.Close()
	log.Println(len(phrasesarr_fi_FI_finance))
	
	createparagraphs.CreatePr(db,"fi_FI","finance",keywordsarr_fi_FI_finance,phrasesarr_fi_FI_finance,3) 


}