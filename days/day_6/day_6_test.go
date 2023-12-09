package day_6

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
			expected: 288,
			input:    "test_input/test.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_6_1(lines))
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 71503,
			input:    "test_input/test.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_6_2(lines))
	}
}
