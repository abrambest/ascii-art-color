package functions

import (
	"fmt"
	"strings"
)

func idWordToChange(wordToChange string, txt []string) string {

	str := ""
	var num []string

	word := chekWordContains(txt[0], wordToChange)

	if word {
		index := 0
		for _, v := range txt {
			for index = strings.Index(txt[0], wordToChange); index < len(v); index++ {
				if have(string(v[index]), wordToChange) {
					str += string(v[index])
					num = append(num, fmt.Sprint(index))

				}
			}
		}
	}
	if !word {
		for _, v := range txt {
			for i, j := range v {
				if have(string(j), wordToChange) {
					str += string(j)
					num = append(num, fmt.Sprint(i))

				}
			}
		}
	}

	if strings.ContainsAny(wordToChange, str) {

		return strings.Join(num, " ")
	}
	return ""

}

func chekWordContains(txt, wordToChange string) bool {
	txtArr := strings.Split(txt, " ")
	for _, v := range txtArr {
		if v == wordToChange {
			return true
		}
	}
	return false
}

func have(j, word string) bool {
	for _, v := range word {
		if j == string(v) {
			return true
		}

	}
	return false
}
