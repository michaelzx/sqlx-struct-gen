package utils

import "fmt"

var (
	greenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green     = string([]byte{27, 91, 51, 50, 109})
	white     = string([]byte{27, 91, 51, 55, 109})
	yellow    = string([]byte{27, 91, 51, 51, 109})
	red       = string([]byte{27, 91, 51, 49, 109})
	blue      = string([]byte{27, 91, 51, 52, 109})
	magenta   = string([]byte{27, 91, 51, 53, 109})
	cyan      = string([]byte{27, 91, 51, 54, 109})
	rest      = string([]byte{27, 91, 48, 109})
)

func PrintGreenBg(str string) {
	fmt.Println(greenBg, str, rest)
}

func PrintWhiteBg(str string) {
	fmt.Println(whiteBg, str, rest)
}
func PrintYellowBg(str string) {
	fmt.Println(yellowBg, str, rest)
}
func PrintRedBg(str string) {
	fmt.Println(redBg, str, rest)
}
func PrintBlueBg(str string) {
	fmt.Println(blueBg, str, rest)
}
func PrintMagentaBg(str string) {
	fmt.Println(magentaBg, str, rest)
}
func PrintCyanBg(str string) {
	fmt.Println(cyanBg, str, rest)
}
func PrintGreen(str string) {
	fmt.Println(green, str, rest)
}
func PrintWhite(str string) {
	fmt.Println(white, str, rest)
}
func PrintYellow(str string) {
	fmt.Println(yellow, str, rest)
}
func PrintRed(str string) {
	fmt.Println(red, str, rest)
}
func PrintBlue(str string) {
	fmt.Println(blue, str, rest)
}
func PrintMagenta(str string) {
	fmt.Println(magenta, str, rest)
}
func PrintCyan(str string) {
	fmt.Println(cyan, str, rest)
}
