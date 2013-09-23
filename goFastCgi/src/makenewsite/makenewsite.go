package makenewsite

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"time"
)

func Makenew(locale string, themes string, host string, pathinfo string,title string, paragraphid int64) {

	now := time.Now().Unix()

	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into sites(Created,Updated,AllHits,Hits,Locale,Themes,Site,Pathinfo,Title) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var rs sql.Result
	hit := int(0)
	if rs, err = stmt.Exec(now, now, hit, hit, locale, themes, host, pathinfo, title); err != nil {

		log.Fatal(err)

	} else {

		siteid, _ := rs.LastInsertId()
		log.Println(siteid)

		stmt, err := tx.Prepare("update paragraphs set Siteid=? where rowid=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		if rs, err = stmt.Exec(siteid,paragraphid); err != nil {

			log.Fatal(err)

		} else {

			log.Println("Update ok")
		}

	}
	tx.Commit()
}
