package functions

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Start() (*string, string, string, error) {

	args := os.Args[1:]

	if len(args) == 3 {
		colPtr := flag.String("color", "Black", "specify a color (options: Black, Red, Green, Yellow, Blue)")
		flag.Parse()
		return *&colPtr, args[1], args[2], nil
	}
	if len(args) == 2 {
		colPtr := flag.String("color", "Black", "specify a color (options: Black, Red, Green, Yellow, Blue)")
		flag.Parse()
		return *&colPtr, "", args[1], nil
	}
	return nil, "", "", errors.New("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")

}

func ReadAscii() []string {

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
	}

	str0 := ""
	arrSplit := strings.Split(string(file), "\n\n")
	for i, v := range arrSplit[0] {
		if i > 0 {
			str0 += string(v)
		}

	}
	arrSplit[0] = str0

	return arrSplit

}

func PrintAsciiArt(color *string, colorWord string, txt, alhpAscii []string) {
	var setColor ColorType

	lineCharAscii := ""
	mapNumWord := ""

	fmt.Println("color word - ", len(colorWord))

	if colorWord != "" {

		mapNumWord = idWordToChange(colorWord, txt)
	}

	if txt[0] == "" {
		fmt.Println()
		return
	}

	for _, word := range txt {

		if word == "" {
			fmt.Println()
			continue
		}
		firstLine := true

		for k := 0; k < 8; k++ {
			str := ""

			for i, s := range word {

				if mapNumWord != "" || colorWord == "" {

					if findIndexChar(mapNumWord, fmt.Sprint(i)) || colorWord == "" {

						setColor = Colorize(color)

						strCut := strings.Split(alhpAscii[s-32], "\n")

						for _, ascii := range strCut[k] {

							lineCharAscii += string(ascii)

						}

						str += string(setColor) + lineCharAscii + string(ColorReset)

						lineCharAscii = ""

					} else {
						strCut := strings.Split(alhpAscii[s-32], "\n")
						for _, ascii := range strCut[k] {

							str += string(ascii)

						}

					}

				} else {
					strCut := strings.Split(alhpAscii[s-32], "\n")
					for _, ascii := range strCut[k] {

						str += string(ascii)

					}
				}

			}

			if firstLine {
				firstLine = false

				fitConsole(len(word)*8 + 20)
			}
			fmt.Print(str)

			fmt.Println()
		}

	}

}

func findIndexChar(mapNumWord, i string) bool {
	arr := strings.Split(mapNumWord, " ")
	for _, n := range arr {
		if n == i {
			return true
		}
	}
	return false
}

func CheckTxt(str string) error {

	noChar := ""
	for _, s := range str {
		if s < 0 || s > 127 {
			noChar += string(s)
		}

	}
	if noChar != "" {
		return errors.New(fmt.Sprintf("a character \"%v\" is not available.", noChar))
	}
	return nil
}
func check(errMsg string, err error) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}

func fitConsole(s int) {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	check("Error measuring console size:", err)

	outStr := string(out)
	outStr = strings.TrimSpace(outStr)
	heightWidth := strings.Split(outStr, " ")
	width, err := strconv.Atoi(heightWidth[1])
	check("Error measuring console size:", err)

	fmt.Println("S - ", s, "W - ", width)

	if s > width {
		fmt.Println("The input string doesn't fit into terminal.")
		os.Exit(1)
	}
}
