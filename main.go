package main

import (
	"ascii_art/functions"
	"fmt"
	"strings"
)

func main() {

	color, colorWord, txt, err := functions.Start()
	fmt.Println(*color)
	fmt.Println(colorWord)
	fmt.Println(txt)
	if err != nil {
		fmt.Println(err)
	}
	if txt == "" {
		return
	}

	err = functions.CheckTxt(txt)
	if err != nil {
		fmt.Println(err)
		return
	}

	arrTxt := strings.Split(txt, "\\n")

	alphAscii := functions.ReadAscii()
	functions.PrintAsciiArt(color, colorWord, arrTxt, alphAscii)

}
