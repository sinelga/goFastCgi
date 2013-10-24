package insertsentences

import (
	"cleansentence"
	"log"
	"markgen"
	ml "marklib"
	"strings"
)

func Insert(c *ml.Chain, locale string, themes string) []string {

	var clsentensesarr []string

	sentences := markgen.Generate(c, locale, themes)

	for i := 0; i < len(sentences); i++ {

		words := strings.Split(sentences[i], " ")

		if len(words) < 2 {

			log.Println(sentences[i])

		} else {

			if len(sentences[i]) > 20 {

				clsentence := cleansentence.Clean(sentences[i])
				clsentensesarr = append(clsentensesarr, clsentence)

			}
		}
	}

	return clsentensesarr

}
