package day_1

import (
	"strconv"
	"strings"
	"unicode"
)

type DigitNumbers struct {
	Numbers map[string]int
}

func newDigitNumbers() *DigitNumbers {
	return &DigitNumbers{
		Numbers: map[string]int{
			"zero":  0,
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		},
	}
}

func T_1_1(lines []string) int {
	sum := 0
	for _, line := range lines {
		firstNumber, lastNumber := processLines(line)
		sum += firstNumber*10 + lastNumber
	}

	return sum
}

func T_1_2(lines []string) int {
	sum := 0
	for _, line := range lines {
		processedLine := replaceSpelledOutNumbers(line)
		firstNumber, lastNumber := processLines(processedLine)
		sum += firstNumber*10 + lastNumber
	}

	return sum
}

func replaceSpelledOutNumbers(line string) string {
	digit := newDigitNumbers()
	var result []string

	var currentNum string
	for idx, char := range line {
		currentNum += string(char)

		if char >= '0' && char <= '9' {
			result = append(result, string(char))
			currentNum = ""
		}

		if shouldCheckForSuffix(idx, char) {
			for word, number := range digit.Numbers {
				if strings.HasSuffix(currentNum, word) {
					result = append(result, strconv.Itoa(number))
					currentNum = string(char)
					break
				}
			}
		}
	}
	result = append(result, currentNum)

	return strings.Join(result, "")
}

func shouldCheckForSuffix(idx int, char rune) bool {
	return idx+1 >= 3 && strings.Contains("eorxnt", string(char))
}

func processLines(line string) (int, int) {
	var firstNumber, lastNumber int

	for _, char := range line {
		if unicode.IsDigit(char) {
			digit, _ := strconv.Atoi(string(char))
			if firstNumber == 0 {
				firstNumber = digit
			} else {
				lastNumber = digit
			}
		}
	}

	if lastNumber == 0 {
		lastNumber = firstNumber
	}

	return firstNumber, lastNumber
}
