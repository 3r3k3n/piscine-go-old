package piscine

func intfor2(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func BasicAtoi(s string) int {
	stroka := []rune(s)
	n := 0
	for range stroka {
		n++
	}
	ans := 0
	for i := 0; i < n; i++ {
		if stroka[i] < '0' || stroka[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += intfor2(stroka[i])
		}
	}
	return ans
}
