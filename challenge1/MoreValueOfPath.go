package challenge_1

// challenge 1
func MoreValueOfPath(data [][]int) int {

	var answerList = make([]int, len(data))

	answerList[0] = data[0][0]

	for i := 1; i < len(data); i++ {

		_answerList := make([]int, len(answerList))
		copy(_answerList, answerList)

		for j := 0; j < len(data[i]); j++ {

			var sumMaxAnswerTestcase int

			if j == 0 {
				sumMaxAnswerTestcase = _answerList[0]
			} else if j == i {
				sumMaxAnswerTestcase = _answerList[i]
			} else if _answerList[j-1] > _answerList[j] {
				sumMaxAnswerTestcase = _answerList[j-1]
			} else {
				sumMaxAnswerTestcase = _answerList[j]
			}

			answerList[j] = sumMaxAnswerTestcase + data[i][j]
		}

	}

	return findLargestNumber(answerList)
}

func findLargestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}
