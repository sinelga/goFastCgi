package titlegen

import (
	"encoding/csv"
	//	"fmt"
	"io"
	"strings"
	//	"testing"
	"unicode/utf8"
)

func GetTitle(locale string, themes string, site string, pathinfo string) string {
	var title string

	if pathinfo == "/" {

		strReader := strings.NewReader(site)
		reader := csv.NewReader(strReader)
		runes, _ := utf8.DecodeRune([]byte("."))
		reader.Comma = runes

		rec := getrecords(reader)

		if len(rec) == 3 {

			title = rec[1]

		} else if len(rec) == 2 {
			title = rec[0]
		}

		//		fmt.Println("Title-> " + title)

	} else {
		strReader := strings.NewReader(pathinfo)
		reader := csv.NewReader(strReader)
		runes, _ := utf8.DecodeRune([]byte("/"))
		reader.Comma = runes
		rec := getrecords(reader)

		if len(rec) == 3 {

			title = rec[1] + " " + strings.Split(rec[2], ".")[0]

		} else if len(rec) == 2 {

			title = strings.Split(rec[1], ".")[0]
		}


	}

	return title
}

func getrecords(reader *csv.Reader) []string {

	var retrecord []string

	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {

			retrecord = nil
		}



		retrecord = record

	}
	return retrecord
}
