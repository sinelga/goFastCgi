package prtitlegen

import (
	"comutils"
//	"math/rand"
//	"time"
//	"log"
)

func Generate(keywordsarr []string) string {

	prtitle := keywordsarr[comutils.Random(0, len(keywordsarr))]
	prtitle =comutils.UpcaseInitial(prtitle) +" " +keywordsarr[comutils.Random(0, len(keywordsarr))] +"."
	
	return prtitle

}
