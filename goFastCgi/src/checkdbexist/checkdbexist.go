package checkdbexist

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
)

func Checkdb(db *sql.DB, host string, pathinfo string) (int64, string, string, string) {
	var locale string
	var themes string
	var title string
	var rowid int64
	var hits int

	//	tx, err := db.Begin()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	stmt, err := tx.Prepare("update sites set Hits(Created,Locale,Themes,Ptitle,Pphrase,Host,Locallink) values(?,?,?,?,?,?,?)")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer stmt.Close()

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

	//	log.Println("Checkdb:Start", locale, themes)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	newhits := hits + 1

	sqlstr = "update sites set Hits = ? where rowid = ?"
	log.Println(sqlstr)

	stmt, err := tx.Prepare(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(newhits, rowid); err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	return rowid, locale, themes, title

}
