package addfreeparagraph

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	//	"domains"
	"log"

//	"strconv"

)

func AddPr(db *sql.DB, siteid int64, locale string, themes string) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("update paragraphs set Siteid=? where rowid IN (select rowid from paragraphs where Siteid is null limit 1)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var rs sql.Result
	if rs, err = stmt.Exec(siteid); err != nil {

		log.Fatal(err)

	} else {

		log.Println("Update ok rs ", rs.LastInsertId)
	}
	tx.Commit()

}
