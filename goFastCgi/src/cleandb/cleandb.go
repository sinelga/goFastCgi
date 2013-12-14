package cleandb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"strconv"
)

func Makeclean(db *sql.DB, siteid int) {

	var paragraphsidtodelete []int

	sqlstr := "SELECT rowid FROM paragraphs where Siteid=" + strconv.Itoa(siteid)

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var paragrphid int

		rows.Scan(&paragrphid)
		
		paragraphsidtodelete = append(paragraphsidtodelete, paragrphid)
	}

	if len(paragraphsidtodelete) > 5 {

		deleteall(db, siteid, paragraphsidtodelete)

	}

}

func deleteall(db *sql.DB, siteid int, paragraphsidtodelete []int) {

//	log.Println("start delete")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	for _, parid := range paragraphsidtodelete {

		stmt, err := tx.Prepare("delete from sentences where Prid=?")

		if _, err = stmt.Exec(parid); err != nil {
			log.Fatal(err)

		}

		stmt.Close()

		stmt, err = tx.Prepare("delete from paragraphs where Siteid=?")

		if _, err = stmt.Exec(siteid); err != nil {
			log.Fatal(err)

		}

		stmt.Close()

		stmt, err = tx.Prepare("delete from sites where id=?")
		log.Println("delete site",siteid)

		if _, err = stmt.Exec(siteid); err != nil {
			log.Fatal(err)

		}
		stmt.Close()
	}
	tx.Commit()

}
