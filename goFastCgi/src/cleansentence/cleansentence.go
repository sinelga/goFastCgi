package cleansentence

import (
	"strings"
)

func Clean(sentence string) string {

	restring := strings.Replace(sentence, "-", " ", -1)
	restring = strings.Replace(restring, "∇", " ", -1)
	restring = strings.Replace(restring, "*", " ", -1)

	restring = strings.Replace(restring, "  ", " ", -1)
	restring = strings.Replace(restring, "  ", " ", -1)

	if restring[:1] == " " {

		restring = strings.Replace(restring, " ", "", 1)
	}

	if restring[:1] == " " {

		restring = strings.Replace(restring, " ", "", 1)
	}

	restring = restring +"."

	return restring
}
