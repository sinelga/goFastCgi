package htmlfileexist

import (
	"addfreeparagraph"
	"checkdbexist"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createpage"
	"database/sql"
	"domains"
	"getalldbparagraphs"
	"log"
	"log/syslog"
	"os"
	"sentencesforpr"
	"strconv"
	"time"
)

func StartCheck(golog syslog.Writer, htmlfile string, host string, pathinfo string) {

	var paragraphsarr []domains.Paragraph

	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	webcontents := checkdbexist.Checkdb(golog, db, host, pathinfo)

	if webcontents.Rowid > 0 {

		deltamin := int(time.Since(webcontents.Updated).Minutes())

		if webcontents.Hits < 5 && deltamin > 30 {

			paragraphsarr = getalldbparagraphs.GetAllPr(db, webcontents.Rowid, host)

			for i, paragraph := range paragraphsarr {

				paragraphsarr[i].Sentences = sentencesforpr.GetSents(db, paragraph.Rowid)
			}

			addfreeparagraph.AddPr(golog, db, webcontents.Rowid, webcontents.Locale, webcontents.Themes)

			webcontents.Paragraphs = paragraphsarr
			createpage.CreatePg(htmlfile, webcontents)

		} else {

			golog.Info("Dont Update page hits > (set 5) --> " + strconv.Itoa(webcontents.Hits) + " or delatamin (set 30) to shot " + strconv.Itoa(deltamin) + " " + htmlfile)

		}

	} else {

		f, err := os.Open(htmlfile)
		if err != nil {
			//			fmt.Println(err)
			golog.Err(err.Error())
			//			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			//			fmt.Println(err)
			golog.Err(err.Error())
//			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():

			golog.Info("directory " + htmlfile)
			golog.Warning("try delete index.html file " + htmlfile + "/index.html")

			if _, err := os.Stat(htmlfile + "/index.html"); err != nil {
				if os.IsNotExist(err) {
					//                        return false
					golog.Warning("Don't exit " + htmlfile + "/index.html")
				}
			} else {
				golog.Warning("OK delete " + htmlfile + "/index.html")
				os.Remove(htmlfile + "/index.html")

			}
			//			os.Remove(htmlfile+/index.html)

		case mode.IsRegular():
			// do file stuff
			//			fmt.Println("file")
			golog.Info("file " + htmlfile)
			golog.Warning("recod don't exit must delete file " + htmlfile)
			os.Remove(htmlfile)

		}

	}

	if db.Close(); err != nil {

		log.Fatal(err)
	}

}
