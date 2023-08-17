package functions

import (
	"fmt"
	"strings"
)

func idWordToChange(wordToChange string, txt []string) string {

	str := ""
	var num []string

	for _, v := range txt {
		for i, j := range v {
			if have(string(j), wordToChange) {
				str += string(j)
				num = append(num, fmt.Sprint(i))

			}
		}
	}

	if strings.ContainsAny(wordToChange, str) {

		return strings.Join(num, " ")
	}
	return ""

}

func have(j, word string) bool {
	for _, v := range word {
		if j == string(v) {
			return true
		}

	}
	return false
}
