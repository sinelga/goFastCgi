package markgen

import (
	//	"bufio"
	//	"bytes"
	//	"encoding/json"
	////	"flag"
	//	"fmt"
	//	"io"
//	"io/ioutil"
//	"math/rand"
	//	"net/http"
	////	"os"
	"strings"
	//	//	"strconv"
//	"log"
	ml "marklib"
//	"time"
)

//type Tsentences struct {
//	Sentences []string
//}

var markfile string

func Generate(c *ml.Chain,locale string, themes string) []string {
	numWords := 1000
//	prefixLen := 1

//	if locale == "fi_FI" && themes == "finance" {
//
//		markfile = "/home/juno/git/goFastCgi/goFastCgi/fi_FI_finance.txt"
//		//		remotehosturl = "http://singo.jntlserv15.appspot.com/mark"
//	}
//
//	rand.Seed(time.Now().UnixNano())
//
////	c := ml.NewChain(prefixLen)
//
//	fData, err := ioutil.ReadFile(markfile)
//
//	if err != nil {
//
//		log.Fatalln(err.Error())
//		return nil
//	}
//	c.Write(fData)

	for i := 0; i < 5; i++ {

		text := c.Generate(numWords) // Generate text.

		sentences := strings.Split(text, ". ")

		sentenses_1 := make([]string, len(sentences)-1)

		for i, sentence := range sentences {

			if i > 0 && i < len(sentences) {

				sentenses_1[i-1] = sentence
//				log.Println(sentenses_1[i-1]) 

			}
		}
		return sentenses_1
	}
	
	return nil
	}


