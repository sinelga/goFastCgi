package checkdbexist

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"time"
	"domains"
)

func Checkdb(db *sql.DB, host string, pathinfo string) domains.WebContents {
	var locale string
	var themes string
	var title string
	var site string
	var rowid int64
	var created int64
	var updated int64
	var hits int
	var allhits int
//	var deltamin int

	now := time.Now().Unix()

	sqlstr := "select rowid,Created,Updated,Hits,Locale,Themes,Title,Site,Allhits from sites where Site='" + host + "' and Pathinfo='" + pathinfo + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&rowid,&created,&updated, &hits, &locale, &themes, &title,&site,&allhits)

	}
	rows.Close()
	
	webcontents := domains.WebContents{
	
		Rowid: rowid,
		Locale: locale,
		Themes: themes,
		Title: title,
		Site: site,
		PathInfo: pathinfo,
		Created: time.Unix(created, 0),
		Updated: time.Unix(updated, 0),
		
	
	}
	
	
	lastHitwas := time.Unix(updated, 0)
	
	log.Println("Since last click ",int(time.Since(lastHitwas).Minutes()))
//	deltamin = int(time.Since(lastHitwas).Minutes())

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	newhits := hits + 1

	sqlstr = "update sites set Hits = ?,Updated=? where rowid = ?"
//	log.Println(sqlstr)

	stmt, err := tx.Prepare(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(newhits,now,rowid); err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	return webcontents

}
