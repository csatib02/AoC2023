package day_5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AoC/util"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 35,
			input:    "test_input/test.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_5_1(lines))
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 46,
			input:    "test_input/test.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_5_2(lines))
	}
}
