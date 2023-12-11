package day_8

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
			expected: 6,
			input:    "test_input/test_1.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_8_1(lines))
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 6,
			input:    "test_input/test_2.txt",
		},
	}

	for _, test := range tests {
		lines, err := util.GetData(test.input)

		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, T_8_2(lines))
	}
}
