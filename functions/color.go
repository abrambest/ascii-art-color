package functions

import "fmt"

type ColorType string

const (
	ColorBlack  ColorType = "\u001b[30m"
	ColorRed    ColorType = "\033[38;2;255;0;0m"
	ColorGreen  ColorType = "\u001b[32m"
	ColorYellow ColorType = "\u001b[33m"
	ColorBlue   ColorType = "\u001b[34m"
	ColorReset  ColorType = "\u001b[0m"
)

var colorMap = map[string]ColorType{
	"black":  ColorBlack,
	"red":    ColorRed,
	"green":  ColorGreen,
	"yellow": ColorYellow,
	"blue":   ColorBlue,
}

func Colorize(color *string) ColorType {

	setColor, found := colorMap[*color]
	if !found {
		fmt.Println("Invalid color. Using default Color.")
		setColor = ColorReset
	}
	return setColor

}
