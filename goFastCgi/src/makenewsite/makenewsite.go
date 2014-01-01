package makenewsite

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"log"
	"time"
)

func Makenew(firstpage domains.FirstPage) {

	now := time.Now().Unix()

	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
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
	//	defer stmt.Close()

	var rs sql.Result
	hit := int(0)
	if rs, err = stmt.Exec(now, now, hit, hit, firstpage.Locale, firstpage.Themes, firstpage.Domain, firstpage.Pathinfo, firstpage.Title); err != nil {

		log.Fatal(err)

	} else {

		siteid, _ := rs.LastInsertId()
		stmt, err := tx.Prepare("insert into paragraphs(Created,Siteid,Locale,Themes,Ptitle,Pphrase,Host,Locallink) values(?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		//		defer stmt.Close()
		if rs, err = stmt.Exec(now, siteid, firstpage.Locale, firstpage.Themes, firstpage.Ptitle, firstpage.Pphrase, firstpage.Phost, firstpage.Plocallink); err != nil {

			log.Fatal(err)

		} else {

			paragrphid, _ := rs.LastInsertId()

			for _, sentence := range firstpage.Sentences {
				stmt, err := tx.Prepare("insert into sentences(Prid,Sentence) values(?,?)")
				if err != nil {
					log.Fatal(err)
				} else {

					if _, err = stmt.Exec(paragrphid, sentence); err != nil {
						log.Fatal(err)

					}

				}
			}
		}

	}
	stmt.Close()
	tx.Commit()

	if db.Close(); err != nil {

		log.Fatal(err)
	}
}
