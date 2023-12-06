package day_3

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Matrix[x][y] == Matrix[row][column]
func T_3_1(lines []string) int {
	sum := 0
	matrix := createMatrixFromInput(lines)
	for rowIdx, row := range matrix {
		for colIdx, char := range row {
			if isSymbol(string(char)) {
				numberToAdd := checkForNumbersAround(matrix, rowIdx, colIdx, false)
				sum += numberToAdd
			}
		}
	}

	return sum
}

func T_3_2(lines []string) int {
	sum := 0
	matrix := createMatrixFromInput(lines)
	for rowIdx, row := range matrix {
		for colIdx, char := range row {
			if isGearSymbol(string(char)) {
				numberToAdd := checkForNumbersAround(matrix, rowIdx, colIdx, true)
				sum += numberToAdd
			}
		}
	}

	return sum
}

func createMatrixFromInput(lines []string) [][]string {
	var matrix [][]string
	for _, line := range lines {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	return matrix
}

// Everything that is not a dot or a number counts as a symbol
func isSymbol(char string) bool {
	return char != "." && !unicode.IsDigit(rune(char[0]))
}

func checkForNumbersAround(matrix [][]string, rowIdx, colIdx int, part2 bool) int {
	var numbersToCount = make(map[int]int)
	// Check that we are not on the edges
	// Altough there should be no symbols on the edges...
	if rowIdx != 0 && rowIdx != len(matrix)-1 {
		if colIdx != 0 && colIdx != len(matrix[rowIdx])-1 {
			// Check the column to the left
			if unicode.IsDigit(rune(matrix[rowIdx][colIdx-1][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx, colIdx-1)
				numbersToCount[index] = num
			}
			// Check diagonal up left
			if unicode.IsDigit(rune(matrix[rowIdx-1][colIdx-1][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx-1, colIdx-1)
				numbersToCount[index] = num
			}
			// Check the row above
			if unicode.IsDigit(rune(matrix[rowIdx-1][colIdx][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx-1, colIdx)
				numbersToCount[index] = num
			}
			// Check diagonal up right
			if unicode.IsDigit(rune(matrix[rowIdx-1][colIdx+1][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx-1, colIdx+1)
				numbersToCount[index] = num
			}
			// Check the column to the right
			if unicode.IsDigit(rune(matrix[rowIdx][colIdx+1][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx, colIdx+1)
				numbersToCount[index] = num
			}
			// Check diagonal down right
			if unicode.IsDigit(rune(matrix[rowIdx+1][colIdx+1][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx+1, colIdx+1)
				numbersToCount[index] = num
			}
			// Check the row below
			if unicode.IsDigit(rune(matrix[rowIdx+1][colIdx][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx+1, colIdx)
				numbersToCount[index] = num
			}
			// Check diagonal down left
			if unicode.IsDigit(rune(matrix[rowIdx+1][colIdx-1][0])) {
				num, index := getNumberFromMatrix(matrix, rowIdx+1, colIdx-1)
				numbersToCount[index] = num
			}
		}
	}
	if part2 {
		// A real gear symbol should have 2 numbers around it
		if len(numbersToCount) == 2 {
			return multplyNumbers(numbersToCount)
		}

		return 0
	}

	return addNumbers(numbersToCount)
}

// If we see a number we should check both ways
func getNumberFromMatrix(matrix [][]string, rowIdx, colIdx int) (int, int) {
	var numberWithIndex = make(map[int]string)
	numberWithIndex[colIdx] = matrix[rowIdx][colIdx]

	// Check the row to the right until we see a dot
	i := 1
	for {
		// Check that we are not on the edges
		if (colIdx + i) >= len(matrix[rowIdx]) {
			break
		}

		if unicode.IsDigit(rune(matrix[rowIdx][colIdx+i][0])) {
			numberWithIndex[colIdx+i] = matrix[rowIdx][colIdx+i]
		} else {
			break
		}
		i++
	}

	// Check the row to the left until we see a dot
	i = 1
	for {
		// Check that we are not on the edges
		if (colIdx - i) < 0 {
			break
		}

		if unicode.IsDigit(rune(matrix[rowIdx][colIdx-i][0])) {
			numberWithIndex[colIdx-i] = matrix[rowIdx][colIdx-i]
		} else {
			break
		}
		i++
	}

	// Assemble the number
	num, idx := assembleNumber(numberWithIndex, rowIdx)
	return num, idx
}

func addNumbers(numbersToAdd map[int]int) int {
	sum := 0
	for _, number := range numbersToAdd {
		sum += number
	}

	return sum
}

func getKeys(numberWithIndex map[int]string) []int {
	var keys []int
	for k := range numberWithIndex {
		keys = append(keys, k)
	}

	return keys
}

func assembleNumber(numberWithIndex map[int]string, rowIdx int) (int, int) {
	var number []string
	// Sort the numbers by the index
	// they had in the matrix
	keys := getKeys(numberWithIndex)
	sort.Ints(keys)
	var idxStr string
	for _, key := range keys {
		idxStr += strconv.Itoa(key)
		// Put together the keys into a number
		number = append(number, numberWithIndex[key])
	}
	// Assemble index and add row index to it
	// so we don't have a case where we miss out on a number
	// because it is in the same column as another
	idx, _ := strconv.Atoi(idxStr)
	idx += rowIdx

	// Assemble the number
	num, _ := strconv.Atoi(strings.Join(number, ""))

	return num, idx
}

// We only care about the gear symbols
func isGearSymbol(char string) bool {
	return char == "*"
}

func multplyNumbers(numbersToAdd map[int]int) int {
	product := 1
	for _, number := range numbersToAdd {
		product *= number
	}

	return product
}
