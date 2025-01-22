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
