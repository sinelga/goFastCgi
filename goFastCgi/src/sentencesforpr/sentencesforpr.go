package sentencesforpr

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
//	"log"
	"log/syslog"
	"strconv"
)

func GetSents(golog syslog.Writer,db *sql.DB, prid int64) []string {

	var sentencesarr []string
	sqlstr := "select Sentence from Sentences where Prid =" + strconv.FormatInt(prid, 10)
//	log.Println(sqlstr)
	rows, err := db.Query(sqlstr)
	if err != nil {
//		log.Fatal(err)
		golog.Err(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var sentence string
		rows.Scan(&sentence)
		sentencesarr = append(sentencesarr, sentence)
	}
	return sentencesarr
}
