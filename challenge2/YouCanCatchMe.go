package challenge_2

import (
	"fmt"
	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func Validate(decode []int, code string) bool {
	for i, c := range code {
		if c != 'R' && c != 'L' && c != '=' {
			panic(`input only "L", "R", "="`)
		}

		if c == 'R' && !(decode[i] < decode[i+1]) {
			return false
		} else if c == 'L' && !(decode[i] > decode[i+1]) {
			return false
		} else if c == '=' && decode[i] != decode[i+1] {
			return false
		}
	}

	return true
}

func YouCanCatchMe(input string) string {
	var decoded = make([]int, len(input)+1)

	for !Validate(decoded, input) {
		for i := 0; i < len(input); i++ {
			if input[i] == 'R' && decoded[i] >= decoded[i+1] {
				decoded[i+1] = decoded[i] + 1
			} else if input[i] == 'L' && decoded[i] <= decoded[i+1] {
				decoded[i] = decoded[i+1] + 1
			} else if input[i] == '=' && decoded[i] != decoded[i+1] {
				if decoded[i] > decoded[i+1] {
					decoded[i+1] = decoded[i]
				} else {
					decoded[i] = decoded[i+1]
				}
			}
		}
	}

	return arrayToString(decoded, "")
}
