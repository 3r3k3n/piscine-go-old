<<<<<<< HEAD:printcomb.go
package main
=======
package printcomb
>>>>>>> f329ff2b6c93e8e5bfeddc71795e85d3af04e28c:printcomb/printcomb.go

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := a + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune(10)
}

func main() {
<<<<<<< HEAD:printcomb.go
	PrintComb
=======
	PrintComb()
>>>>>>> f329ff2b6c93e8e5bfeddc71795e85d3af04e28c:printcomb/printcomb.go
}
