package challenge_1

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadTestCaseFromFileJson(filename string) [][]int {
	var jsonData = [][]int{}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panic("Error read file :", err)
	}
	if err = json.Unmarshal(data, &jsonData); err != nil {
		log.Panic("Error unmarshal :", err)
	}

	return jsonData
}

func TestMoreValueOfPath(t *testing.T) {
	testcase := []struct {
		name   string
		data   [][]int
		expect int
	}{
		{
			name:   "testcase 1 : example testcase",
			data:   loadTestCaseFromFileJson("files/testcase_example.json"),
			expect: 237,
		},
		{
			name:   "testcase 2 : hard testcase",
			data:   loadTestCaseFromFileJson("files/testcase_hard.json"),
			expect: 7273,
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			answer := MoreValueOfPath(c.data)
			assert.Equal(t, c.expect, answer)
		})
	}
}

func TestFindLargestNumber(t *testing.T) {
	testcase := []struct {
		name   string
		data   []int
		expect int
	}{
		{
			name:   "testcase 1",
			data:   []int{1, 2, 3, 4, 5, 6, 7},
			expect: 7,
		},
		{
			name:   "testcase 2",
			data:   []int{7, 6, 5, 4, 3, 2, 1},
			expect: 7,
		},
		{
			name:   "testcase 3",
			data:   []int{-1, 0},
			expect: 0,
		},
		{
			name:   "testcase 4",
			data:   []int{},
			expect: 0,
		},
		{
			name:   "testcase 5",
			data:   []int{-2, -3, -1},
			expect: -1,
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			answer := findLargestNumber(c.data)
			assert.Equal(t, c.expect, answer)
		})
	}

}
