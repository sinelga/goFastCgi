package htmlfileexist

import (
	"addfreeparagraph"
	"checkdbexist"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createpage"
	"database/sql"
	"domains"
	"getalldbparagraphs"
//	"log"
	"log/syslog"
	"os"
	"sentencesforpr"
	"strconv"
//	"strings"
	"time"
	"checkpathinfo"
)

func StartCheck(golog syslog.Writer, htmlfile string, host string, pathinfo string) {

	var paragraphsarr []domains.Paragraph
	
	thispathinfo := checkpathinfo.Check(pathinfo)
	
	golog.Info("StartCheck: thispathinfo "+thispathinfo) 

	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	if err != nil {
//		log.Fatal(err)
		golog.Err(err.Error())
	}

	webcontents := checkdbexist.Checkdb(golog, db, host, thispathinfo)

	if webcontents.Rowid > 0 {

		deltamin := int(time.Since(webcontents.Updated).Minutes())

		if webcontents.Hits < 10 && deltamin >= 0 {

			paragraphsarr = getalldbparagraphs.GetAllPr(golog,db, webcontents.Rowid, host)

			for i, paragraph := range paragraphsarr {

				paragraphsarr[i].Sentences = sentencesforpr.GetSents(golog,db, paragraph.Rowid)
			}

			addfreeparagraph.AddPr(golog, db, webcontents.Rowid, webcontents.Locale, webcontents.Themes)

			webcontents.Paragraphs = paragraphsarr
			createpage.CreatePg(golog,htmlfile, webcontents)

		} else {

			golog.Info("Dont Update page hits > (set 5) --> " + strconv.Itoa(webcontents.Hits) + " or delatamin (set 30) to shot " + strconv.Itoa(deltamin) + " " + htmlfile)

		}

	} else {

		if f, err := os.Open(htmlfile); err != nil {

			golog.Err(err.Error())
			
		} else {

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

						golog.Warning("Don't exit " + htmlfile + "/index.html not think to delete")
					}
				} else {
					golog.Warning("OK delete " + htmlfile + "/index.html")
					os.Remove(htmlfile + "/index.html")

				}

			case mode.IsRegular():

				golog.Warning("recod don't exit must delete file " + htmlfile)
				os.Remove(htmlfile)

			}

		}

	}

	if db.Close(); err != nil {

		golog.Err(err.Error())
	}

}
