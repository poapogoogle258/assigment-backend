package challenge_2

import (
	"fmt"
	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func YouCanCatchMe(input string) string {
	var decoded = make([]int, len(input)+1)

	for changed := true; changed; {
		changed = false
		for i := 0; i < len(input); i++ {
			if input[i] == 'R' && decoded[i] >= decoded[i+1] {
				decoded[i+1]++
				changed = true
			} else if input[i] == 'L' && decoded[i] <= decoded[i+1] {
				decoded[i]++
				changed = true
			} else if input[i] == '=' && decoded[i] != decoded[i+1] {
				if decoded[i] > decoded[i+1] {
					decoded[i+1] = decoded[i]
				} else {
					decoded[i] = decoded[i+1]
				}
				changed = true
			}
		}
	}

	return arrayToString(decoded, "")
}
