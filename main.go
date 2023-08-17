package main

import (
	"ascii_art/functions"
	"fmt"
	"strings"
)

func main() {

	color, colorWord, txt, err := functions.Start()
	if err != nil {
		fmt.Println(err)
	}
	if txt == "" {
		return
	}

	fmt.Println("color flag - ", *color)
	fmt.Println("word to be colored - ", colorWord)
	fmt.Println("txt - ", txt)

	err = functions.CheckTxt(txt)
	if err != nil {
		fmt.Println(err)
		return
	}

	arrTxt := strings.Split(txt, "\\n")

	alphAscii := functions.ReadAscii()

	fmt.Println("arrTxt - ", arrTxt)

	functions.PrintAsciiArt(color, colorWord, arrTxt, alphAscii)

}
