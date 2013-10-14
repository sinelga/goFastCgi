package createfirstpagedb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"strconv"
)

func FindFreeSentences(locale string, themes string) (int64, []string) {

	var paragraphid int64
	var sentencesarr []string

	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	sqlstr := "select rowid from paragraphs where locale='" + locale + "' and themes='" + themes + "' and Siteid is null limit 1"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&paragraphid)

	}
	rows.Close()

	sqlstr = "select Sentence from sentences where Prid=" + strconv.FormatInt(paragraphid, 10)
//	log.Println(sqlstr)

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
	rows.Close()
	
	db.Close()
	
	return paragraphid, sentencesarr

}
