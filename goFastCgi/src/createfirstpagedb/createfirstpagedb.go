package createfirstpagedb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"strconv"
)

var sentencesid int64
var sentencesarr []string

func FindFreeSentences(locale string, themes string) []string{

	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}
	sqlstr := "select rowid from paragraphs where locale='" + locale + "' and themes='" + themes + "' and Siteid is null limit 1"
	
	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&sentencesid)

	}
	rows.Close()

	log.Println(sentencesid)

	sqlstr = "select Sentence from sentences where Prid=" + strconv.FormatInt(sentencesid, 10)
	log.Println(sqlstr)

	rows, err = db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
		for rows.Next() {
		var sentence string
		rows.Scan(&sentence)
		sentencesarr = append(sentencesarr, sentence)

	}
	
	return sentencesarr

}
