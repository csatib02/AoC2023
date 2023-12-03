package day_2

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
			expected: 8,
			input:    "test_input/test.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_2_1(lines))
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 2286,
			input:    "test_input/test.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_2_2(lines))
	}
}
