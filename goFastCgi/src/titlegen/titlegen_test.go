package titlegen

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"testing"
	"unicode/utf8"

//	"strings"
//	"pkg/comutils"
)

func Test_GetTitle(t *testing.T) {

	//	locale := "fi_FI"
	//	themes := "finance"
	//	site := "www.raha.me"
	var title string

	pathinfo := "/"

	//	strReader := strings.NewReader(pathinfo)
	//	reader := csv.NewReader(strReader)
	if pathinfo == "/" {

		strReader := strings.NewReader("www.raha.me")
		reader := csv.NewReader(strReader)
		runes, _ := utf8.DecodeRune([]byte("."))
		reader.Comma = runes

		rec := getrecords(reader)

		if len(rec) == 3 {

			title = rec[1]

		} else if len(rec) == 2 {
			title = rec[0]
		}

		fmt.Println("Title-> " + title)

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

		fmt.Println("Title-> " + title)

	}

}
func getrecords(reader *csv.Reader) []string {

	var retrecord []string

	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			retrecord = nil
		}

		fmt.Println(record)
		fmt.Println(len(record))

		for _, rec := range record {
			fmt.Println(rec)

		}

		retrecord = record

	}
	return retrecord
}
