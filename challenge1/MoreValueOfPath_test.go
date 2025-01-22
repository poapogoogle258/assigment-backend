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
		{
			name:   "testcase 3 : empty testcase",
			data:   loadTestCaseFromFileJson("files/testcase_empty.json"),
			expect: 0,
		}, {
			name:   "testcase 4 : height 1 testcase",
			data:   loadTestCaseFromFileJson("files/testcase_empty_height_1.json"),
			expect: 15,
		}, {
			name:   "testcase 5 : height 2 testcase",
			data:   loadTestCaseFromFileJson("files/testcase_empty_height_2.json"),
			expect: 17,
		},
		{
			name:   "testcase 6 : height 3 testcase",
			data:   loadTestCaseFromFileJson("files/testcase_empty_height_3.json"),
			expect: 22,
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			answer := MoreValueOfPath(c.data)
			assert.Equal(t, c.expect, answer)
		})
	}
}
