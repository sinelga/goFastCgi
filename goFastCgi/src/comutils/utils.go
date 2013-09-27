package comutils

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"unicode"
)

func UpcaseInitial(str string) string {

	//	for i, v := range str {
	////		return string(unicode.ToUpper(v)) + str[i+1:]
	//		return
	//	}

	runes := []rune(str)

	for i, v := range runes {
	
		return string(unicode.ToUpper(v)) + string(runes[i+1:])

	}

	return ""

}

func Clone(a, b interface{}) {

	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(a)
	dec.Decode(b)
}

func Random(min, max int) int {
	//    rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
