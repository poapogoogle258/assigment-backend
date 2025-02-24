package challenge_1

// challenge 1
func MoreValueOfPath(data [][]int) int {

	answer := make([]int, len(data))
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0][0]
	}

	copy(answer, data[len(data)-1])

	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j < len(data[i]); j++ {
			if answer[j] > answer[j+1] {
				answer[j] = data[i][j] + answer[j]
			} else {
				answer[j] = data[i][j] + answer[j+1]
			}
		}

	}

	return answer[0]
}
