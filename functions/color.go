package functions

import "fmt"

type ColorType string

const (
	ColorBlack  ColorType = "\u001b[30m"
	ColorRed              = "\033[38;2;255;0;0m"
	ColorGreen            = "\u001b[32m"
	ColorYellow           = "\u001b[33m"
	ColorBlue             = "\u001b[34m"
	ColorReset            = "\u001b[0m"
	ColorOrange           = "\033[38;2;255;165;0m"
	ColorPurple           = "\033[38;2;120;81;169m"
	ColorPink             = "\033[38;2;255;192;203m"
	ColorBrown            = "\033[38;2;244;164;96m"
	ColorGrey             = "\033[38;2;112;112;112m"
)

var colorMap = map[string]ColorType{
	"black":  ColorBlack,
	"red":    ColorRed,
	"green":  ColorGreen,
	"yellow": ColorYellow,
	"blue":   ColorBlue,
	"orange": ColorOrange,
	"purple": ColorPurple,
	"pink":   ColorPink,
	"brown":  ColorBrown,
	"grey":   ColorGrey,
}

func Colorize(color *string) ColorType {

	setColor, found := colorMap[*color]
	if !found {
		fmt.Println("Invalid color. Using default Color.")
		setColor = ColorReset
	}
	return setColor

}
