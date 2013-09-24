package sentencesforpr

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"strconv"
)

func GetSents(db *sql.DB, prid int64) []string {

	var sentencesarr []string
	sqlstr := "select Sentence from Sentences where Prid =" + strconv.FormatInt(prid, 10)
	log.Println(sqlstr)
	rows, err := db.Query(sqlstr)
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
