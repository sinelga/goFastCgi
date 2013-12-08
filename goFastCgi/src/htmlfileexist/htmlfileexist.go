package htmlfileexist

import (
	"os"
	"time"
	"addfreeparagraph"
	"checkdbexist"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createpage"
	"database/sql"
	"domains"
	"getalldbparagraphs"
	"log"
	"log/syslog"
	"sentencesforpr"
	"strconv"
)

func StartCheck(golog syslog.Writer, htmlfile string, host string, pathinfo string) {

	var paragraphsarr []domains.Paragraph

	//	db, err := sql.Open("sqlite3", "gofast.db")
	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	webcontents := checkdbexist.Checkdb(golog, db, host, pathinfo)

//	log.Println("StartCheck:webcontents.Hits", webcontents.Hits, "webcontents.Rowid", webcontents.Rowid)

	if webcontents.Rowid > 0 {
		
		deltamin := int(time.Since(webcontents.Updated).Minutes())

		if webcontents.Hits < 10 && deltamin > 10 {

			paragraphsarr = getalldbparagraphs.GetAllPr(db, webcontents.Rowid, host)

			for i, paragraph := range paragraphsarr {

				paragraphsarr[i].Sentences = sentencesforpr.GetSents(db, paragraph.Rowid)
			}

			addfreeparagraph.AddPr(db, webcontents.Rowid, webcontents.Locale, webcontents.Themes)

			webcontents.Paragraphs = paragraphsarr
			createpage.CreatePg(htmlfile, webcontents)

		} else {

			golog.Info("Dont Update page hits > 10 --> "+strconv.Itoa(webcontents.Hits )+" or delatamin to shot "+strconv.Itoa(deltamin)+" "+htmlfile)

		}

	} else {

		golog.Warning("recod don't exit must delete file " + htmlfile)
		os.Remove(htmlfile)

	}

	if db.Close(); err != nil {

		log.Fatal(err)
	}

}
