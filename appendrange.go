package piscine

func AppendRange(min, max int) []int {
	var answer []int

	if min < max {
		for i := min - 1; i < max; i++ {
			answer = append(answer, i+1)
		}
		return answer
	} else {
		return nil
	}
}
