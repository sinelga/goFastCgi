package checkdbexist

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"time"
)

func Checkdb(db *sql.DB, host string, pathinfo string) (int64, string, string, string) {
	var locale string
	var themes string
	var title string
	var rowid int64
	var hits int

	now := time.Now().Unix()

	sqlstr := "select rowid,Hits,Locale,Themes,Title from sites where Site='" + host + "' and Pathinfo='" + pathinfo + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&rowid, &hits, &locale, &themes, &title)

	}
	rows.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	newhits := hits + 1

	sqlstr = "update sites set Hits = ?,Updated=? where rowid = ?"
	log.Println(sqlstr)

	stmt, err := tx.Prepare(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(newhits, now,rowid); err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	return rowid, locale, themes, title

}
