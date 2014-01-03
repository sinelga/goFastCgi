package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"createfirstpage"
	"database/sql"
	//	"encoding/json"
	"clean_pathinfo"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"pushinqueue"
	"strconv"
	"sync"
	"createfirstgz"
)

var startOnce sync.Once
var keywordsarr_fi_FI_finance []string
var phrasesarr_fi_FI_finance []string
var keywordsarr_fi_FI_porno []string
var phrasesarr_fi_FI_porno []string
var keywordsarr_it_IT_finance []string
var phrasesarr_it_IT_finance []string
var keywordsarr_fi_FI_fortune []string
var phrasesarr_fi_FI_fortune []string

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	themes := req.Header.Get("X-THEMES")
	locale := req.Header.Get("X-LOCALE")
	host := req.Header.Get("X-DOMAIN")
	pathinfo := req.Header.Get("X-PATHINFO")

	checkfirstpage(resp, req, locale, themes, host, pathinfo)

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

func checkfirstpage(resp http.ResponseWriter, req *http.Request, locale string, themes string, host string, pathinfo string) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	startOnce.Do(func() {
		startones(*golog)
	})
	pathinfostr := clean_pathinfo.CleanPath(*golog,pathinfo)

	htmlfile := string("www/" + locale + "/" + themes + "/" + host + pathinfostr)

	if fi, err := os.Stat(htmlfile); err != nil {

		if os.IsNotExist(err) {

			fileNotExistCreate(*golog, resp, req, locale, themes, host, pathinfostr, htmlfile)

		} else {

			golog.Err("something wrong??? "+pathinfostr+" "+ htmlfile)

		}

	} else {

		switch mode := fi.Mode(); {

		case mode.IsDir():

			golog.Info("directory " + htmlfile)
			//				golog.Warning("try delete index.html file " + htmlfile + "/index.html")

			if _, err := os.Stat(htmlfile + "/index.html"); err != nil {
				if os.IsNotExist(err) {

					golog.Warning("Don't exit " + htmlfile + "/index.html so create new")
					fileNotExistCreate(*golog, resp, req, locale, themes, host, pathinfostr, htmlfile+"/index.html")
//					http.ServeFile(resp, req, htmlfile+"/index.html")
				}
			} else {
//									golog.Warning("OK for " + htmlfile + "/index.html exist serve normaly")
				//					os.Remove(htmlfile + "/index.html")
//				http.ServeFile(resp, req, htmlfile+"/index.html")

			}

		case mode.IsRegular():

			golog.Warning("app: IsRegular file" + htmlfile +" OK static serveFile")
//			http.ServeFile(resp, req, htmlfile)

			//				os.Remove(htmlfile)

		}

		//		http.ServeFile(resp, req, htmlfile)
		go pushinqueue.PushInQueue(*golog,"redis", locale, themes, host, pathinfostr)

	}
}

func fileNotExistCreate(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, host string, pathinfostr string, htmlfile string) {

	var bytepage []byte
	if locale == "fi_FI" && themes == "finance" {
		bytepage = createfirstpage.CreatePage(golog, locale, themes, host, pathinfostr, keywordsarr_fi_FI_finance, phrasesarr_fi_FI_finance)
	} else if locale == "fi_FI" && themes == "porno" {
		bytepage = createfirstpage.CreatePage(golog, locale, themes, host, pathinfostr, keywordsarr_fi_FI_porno, phrasesarr_fi_FI_porno)

	} else if locale == "it_IT" && themes == "finance" {
		bytepage = createfirstpage.CreatePage(golog, locale, themes, host, pathinfostr, keywordsarr_it_IT_finance, phrasesarr_it_IT_finance)
	} else if locale == "fi_FI" && themes == "fortune" {
		bytepage = createfirstpage.CreatePage(golog, locale, themes, host, pathinfostr, keywordsarr_fi_FI_fortune, phrasesarr_fi_FI_fortune)
	}

//	http.ServeFile(resp, req, htmlfile)
	
	resp.Write(bytepage)
	go createfirstgz.Creategzhtml(htmlfile,bytepage)
	

}

func startones(golog syslog.Writer) {

	golog.Info("startones")
	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select Locale,Themes,Keyword from keywords order by Hits desc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var locale string
		var themes string
		var keyword string

		rows.Scan(&locale, &themes, &keyword)
		if locale == "fi_FI" && themes == "finance" {

			if len(keywordsarr_fi_FI_finance) < 1000 {
				keywordsarr_fi_FI_finance = append(keywordsarr_fi_FI_finance, keyword)
			}
		} else if locale == "fi_FI" && themes == "porno" {
			if len(keywordsarr_fi_FI_porno) < 1000 {
				keywordsarr_fi_FI_porno = append(keywordsarr_fi_FI_porno, keyword)
			}

		} else if locale == "it_IT" && themes == "finance" {
			if len(keywordsarr_it_IT_finance) < 1000 {
				keywordsarr_it_IT_finance = append(keywordsarr_it_IT_finance, keyword)
			}

		} else if locale == "fi_FI" && themes == "fortune" {
			if len(keywordsarr_fi_FI_fortune) < 1000 {
				keywordsarr_fi_FI_fortune = append(keywordsarr_fi_FI_fortune, keyword)
			}
		}

	}
	rows.Close()
	golog.Info("keywordsarr_fi_FI_finance " + strconv.Itoa(len(keywordsarr_fi_FI_finance)))
	golog.Info("keywordsarr_fi_FI_porno " + strconv.Itoa(len(keywordsarr_fi_FI_porno)))
	golog.Info("keywordsarr_it_IT_finance " + strconv.Itoa(len(keywordsarr_it_IT_finance)))
	golog.Info("keywordsarr_fi_FI_fortune " + strconv.Itoa(len(keywordsarr_fi_FI_fortune)))

	rows, err = db.Query("select Locale,Themes,Phrase from phrases")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var locale string
		var themes string
		var phrase string
		rows.Scan(&locale, &themes, &phrase)

		if locale == "fi_FI" && themes == "finance" {
			phrasesarr_fi_FI_finance = append(phrasesarr_fi_FI_finance, phrase)
		} else if locale == "fi_FI" && themes == "porno" {
			phrasesarr_fi_FI_porno = append(phrasesarr_fi_FI_porno, phrase)
		} else if locale == "it_IT" && themes == "finance" {
			phrasesarr_it_IT_finance = append(phrasesarr_it_IT_finance, phrase)
		} else if locale == "fi_FI" && themes == "fortune" {
			phrasesarr_fi_FI_fortune = append(phrasesarr_fi_FI_fortune, phrase)
		}

	}
	rows.Close()

	golog.Info("phrasesarr_fi_FI_finance " + strconv.Itoa(len(phrasesarr_fi_FI_finance)))
	golog.Info("phrasesarr_fi_FI_porno " + strconv.Itoa(len(phrasesarr_fi_FI_porno)))
	golog.Info("phrasesarr_it_IT_finance " + strconv.Itoa(len(phrasesarr_it_IT_finance)))
	golog.Info("phrasesarr_fi_FI_fortune " + strconv.Itoa(len(phrasesarr_fi_FI_fortune)))
	db.Close()

}
