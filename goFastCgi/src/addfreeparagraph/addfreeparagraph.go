package addfreeparagraph

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"queue/findfreeparagraph"
	"time"
)

func AddPr(db *sql.DB, siteid int64, locale string, themes string) {

	now := time.Now().Unix()

	paragraph := findfreeparagraph.FindFromQ(locale, themes)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	var rs sql.Result
	stmt, err := tx.Prepare("insert into paragraphs(Created,Siteid,Locale,Themes,Ptitle,Pphrase,Host,Locallink) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if rs, err = stmt.Exec(now, siteid, locale, themes, paragraph.Ptitle, paragraph.Pphrase, paragraph.Phost, paragraph.Plocallink); err != nil {
		log.Fatal(err)

	} else {

		paragrphid, _ := rs.LastInsertId()

		for _, sentence := range paragraph.Sentences {
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

	stmt.Close()
	tx.Commit()
	

}
