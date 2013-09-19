package createparagraphs

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"time"

//	"log"

)

var keywordsarr_fi_FI_finance []string

func CreatePr(db *sql.DB, locale string, themes string, keywords []string, phrases []string, quant int) {

	log.Println("start CreatePr")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into paragraphs(Created,Locale,Themes,Ptitle,Pphrase) values(?,'fi_FI','finance','lslsls','lslslslsls')")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var rs sql.Result
	for i := 0; i < 3; i++ {

		now := time.Now().Unix()
		if rs, err = stmt.Exec(now); err != nil {

			log.Fatal(err)

		} else {

			id, _ := rs.LastInsertId()
			log.Println(id)

		}

	}
	tx.Commit()

	rows, err := db.Query("select Created from paragraphs where locale='fi_FI' and themes='finance'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var date int64
		rows.Scan(&date)
		log.Println(date)

		//		 t := time.Unix(date, 0)

		//		 log.Println(t.Local())

		//		keywordsarr_fi_FI_finance = append(keywordsarr_fi_FI_finance, keyword)

	}
	rows.Close()
	//	log.Println(len(keywordsarr_fi_FI_finance))

}
