package challenge_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYouCanCatchMe(t *testing.T) {
	testcase := []struct {
		name   string
		data   string
		expect string
	}{
		{
			name:   "testcase 1",
			data:   "LLRR=",
			expect: "210122",
		},
		{
			name:   "testcase 2",
			data:   "==RLL",
			expect: "000210",
		},
		{
			name:   "testcase 3",
			data:   "=LLRR",
			expect: "221012",
		},
		{
			name:   "testcase 4",
			data:   "RRL=R",
			expect: "012001",
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			answer := YouCanCatchMe(c.data)
			assert.Equal(t, c.expect, answer)
		})
	}
}

func TestArrayToString(t *testing.T) {
	testcase := []struct {
		name   string
		data   []int
		expect string
		delim  string
	}{
		{
			name:   "testcase 1",
			data:   []int{1, 2, 3, 4, 5, 6},
			delim:  "",
			expect: "123456",
		},
		{
			name:   "testcase 2",
			data:   []int{-1, -1, -1, -1, -2},
			delim:  "",
			expect: "-1-1-1-1-2",
		},
		{
			name:   "testcase 3",
			data:   []int{3, 4, 5, 6, 1, 2},
			delim:  "+",
			expect: "3+4+5+6+1+2",
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			answer := arrayToString(c.data, c.delim)
			assert.Equal(t, c.expect, answer)
		})
	}
}

func TestValidate(t *testing.T) {

	testcase := []struct {
		name   string
		data   []int
		input  string
		expect bool
	}{
		{
			name:   "testcase 1",
			data:   []int{0, 0, 0, 2, 1, 0},
			input:  "==RLL",
			expect: true,
		},
		{
			name:   "testcase 2",
			data:   []int{2, 2, 1, 0, 1, 2},
			input:  "=LLRR",
			expect: true,
		}, {
			name:   "testcase 3",
			data:   []int{0, 1, 2, 0, 0, 0},
			input:  "==RLL",
			expect: false,
		}, {
			name:   "testcase 4",
			data:   []int{1, 2, 0, 0, 2, 1},
			input:  "=LLRR",
			expect: false,
		},
	}

	for _, tt := range testcase {
		t.Run(tt.name, func(t *testing.T) {
			answer := Validate(tt.data, tt.input)

			assert.Equal(t, tt.expect, answer)
		})
	}

	assert.Panics(t, func() { Validate([]int{2, 2, 1, 0, 1, 2}, "+LLRR") })

}
