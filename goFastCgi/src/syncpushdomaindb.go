package main 

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
    "flag"
    "fmt"
    "log"
    "log/syslog"
    "domains"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
                
    }
    
    golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
    
    var extdomainsarr []domains.Domaincsv
    var domainstopush []string
                
   	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	if err != nil {

		golog.Err(err.Error())
	}
    	sqlstr := "select Locale,Themes,Domain,Ip from extdomains"

		golog.Info("syncpushdomain: Start " + sqlstr)

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err(err.Error())
		}
		defer rows.Close()
		
		for rows.Next() {
			
			var extdomain domains.Domaincsv
			rows.Scan(&extdomain.Locale, &extdomain.Themes, &extdomain.Domain, &extdomain.Ip)
			extdomainsarr=append(extdomainsarr,extdomain)

		}
    	rows.Close()
        golog.Info(extdomainsarr[0].Domain)
    	db.Close()
    	
    	
    	   	db, err = sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/pushdomains.db")
	if err != nil {

		golog.Err(err.Error())
	}
    	
    	sqlstr = "select Domain from pushdomains"
    	rows, err = db.Query(sqlstr)
		if err != nil {

			golog.Err(err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var domaintopush string
			rows.Scan(&domaintopush)
			domainstopush=append(domainstopush,domaintopush) 
		
		}
		
		
		
		
    	
    	
    
}

