package htmlfileexist

import (
	"addfreeparagraph"
	"checkdbexist"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createpage"
	"database/sql"
	"domains"
	"getalldbparagraphs"

	"log/syslog"

	"sentencesforpr"
	"strconv"

	"time"
	"checkpathinfo"
)

func StartCheck(golog syslog.Writer, htmlfile string, host string, pathinfo string) {

	var paragraphsarr []domains.Paragraph
	
	thispathinfo := checkpathinfo.Check(pathinfo)
	
//	golog.Info("StartCheck: thispathinfo "+thispathinfo) 

	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	if err != nil {
		golog.Err(err.Error())
	}

	webcontents := checkdbexist.Checkdb(golog, db, host, thispathinfo)

	if webcontents.Rowid > 0 {

		deltamin := int(time.Since(webcontents.Updated).Minutes())

		if webcontents.Hits < 5 && deltamin >= 0 {

			paragraphsarr = getalldbparagraphs.GetAllPr(golog,db, webcontents.Rowid, host)

			for i, paragraph := range paragraphsarr {

				paragraphsarr[i].Sentences = sentencesforpr.GetSents(golog,db, paragraph.Rowid)
			}

			addfreeparagraph.AddPr(golog, db, webcontents.Rowid, webcontents.Locale, webcontents.Themes)

			webcontents.Paragraphs = paragraphsarr
			createpage.CreatePg(golog,htmlfile, webcontents)

		} else {

			golog.Info("Dont Update page hits > (set 5) --> " + strconv.Itoa(webcontents.Hits) + " or delatamin (set 0) to shot " + strconv.Itoa(deltamin) + " " + htmlfile)

		}

	} else {


		golog.Alert("!!!StartCheck: no record for "+htmlfile+" thispathinfo "+thispathinfo)


	}

	if db.Close(); err != nil {

		golog.Err(err.Error())
	}

}
