package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _,makePrint := range str {
		z01.PrintRune(makePrint)
	}
}