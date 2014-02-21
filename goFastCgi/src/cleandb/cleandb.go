package cleandb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"strconv"
	"os"
)

func Makeclean(db *sql.DB, siteid int) {

	var paragraphsidtodelete []int

	sqlstr := "SELECT rowid FROM paragraphs where Siteid=" + strconv.Itoa(siteid)

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	defer rows.Close()

	for rows.Next() {
		var paragrphid int

		rows.Scan(&paragrphid)

		paragraphsidtodelete = append(paragraphsidtodelete, paragrphid)
	}

	deleteall(db, siteid, paragraphsidtodelete)

}

func deleteall(db *sql.DB, siteid int, paragraphsidtodelete []int) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	for _, parid := range paragraphsidtodelete {

		stmt, err := tx.Prepare("delete from sentences where Prid=?")

		if _, err = stmt.Exec(parid); err != nil {
			log.Fatal(err)
			os.Exit(-1)

		}

		stmt.Close()

		stmt, err = tx.Prepare("delete from paragraphs where Siteid=?")

		if _, err = stmt.Exec(siteid); err != nil {
			log.Fatal(err)
			os.Exit(-1)

		}

		stmt.Close()

		stmt, err = tx.Prepare("delete from sites where id=?")

		if _, err = stmt.Exec(siteid); err != nil {
			log.Fatal(err)
			os.Exit(-1)

		}
		stmt.Close()
	}
	tx.Commit()

}
