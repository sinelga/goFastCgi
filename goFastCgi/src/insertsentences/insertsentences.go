package insertsentences

import (
	"cleansentence"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
	"markgen"
	"strings"
	ml "marklib"
)

func Insert(db *sql.DB, tx *sql.Tx,c *ml.Chain, locale string, themes string, paragraphid int64) {

	sentences := markgen.Generate(c,locale, themes)

	stmt, err := tx.Prepare("insert into sentences(Prid,Sentence) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < len(sentences); i++ {

		words := strings.Split(sentences[i], " ")

		if len(words) < 2 {

			log.Println(sentences[i])

		} else {

			clsentence := cleansentence.Clean(sentences[i])

			if _, err = stmt.Exec(paragraphid, clsentence+"."); err != nil {

				log.Fatal(err)
			}
		}
	}


}
