package day_2

import (
	"strings"
	"unicode"
)

type Game struct {
	Cubes map[string]int
}

// A possible game, is a game that at
// no point of the game has more cubes
// than what is defined here
func newPossibleGame() *Game {
	return &Game{
		Cubes: map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		},
	}
}

func T_2_1(lines []string) int {
	sum := 0
	for _, line := range lines {
		gameId, restOfTheLine := extractGameId(line)
		if isGamePossible(restOfTheLine) {
			sum += gameId
		}
	}

	return sum
}

func T_2_2(lines []string) int {
	sum := 0
	for _, line := range lines {
		_, restOfTheLine := extractGameId(line)
		sum += getProductOfFewestCubeInEachColor(restOfTheLine)
	}

	return sum
}

func extractGameId(line string) (int, string) {
	var gameId int
	for idx, char := range line {
		if string(char) == ":" {
			// The biggest number before the semi-colon is
			// min 1, and at max 3 digits long
			var sumOfDigits int
			if unicode.IsDigit(rune(line[idx-2])) {
				if unicode.IsDigit(rune(line[idx-3])) {
					sumOfDigits = int(line[idx-3]-'0')*100 + int(line[idx-2]-'0')*10 + int(line[idx-1]-'0')
					gameId = sumOfDigits
				} else {
					sumOfDigits = int(line[idx-2]-'0')*10 + int(line[idx-1]-'0')
					gameId = sumOfDigits
				}
			} else {
				sumOfDigits = int(line[idx-1] - '0')
				gameId = sumOfDigits
			}

			break
		}
	}

	restOfTheLine := strings.SplitN(line, ":", 2)
	restOfTheLine[1] += ";"

	return gameId, restOfTheLine[1]
}

func isGamePossible(line string) bool {
	cubesInGame := make(map[string]int)
	skipNext := false

	for idx, char := range line {
		if skipNext {
			skipNext = false
			continue
		}

		if string(char) == ";" {
			if !isGameStillPossible(cubesInGame) {
				return false
			}
			// Reset the map for the next game
			cubesInGame = make(map[string]int)
		}

		// If we see a number it is at most 2 digits long
		// and after the last digit there is a space
		// and then the color, r, g or b
		var countOfCubes int
		if unicode.IsDigit(rune(line[idx])) {
			if unicode.IsDigit(rune(line[idx+1])) {
				colorsFirstLetter := string(line[idx+3])
				countOfCubes = int(line[idx]-'0')*10 + int(line[idx+1]-'0')
				cubesInGame[getColor(colorsFirstLetter)] = countOfCubes

				// we need to skip an iteration if a number is
				// 2 digits long
				skipNext = true
				continue
			} else {
				colorsFirstLetter := string(line[idx+2])
				countOfCubes = int(line[idx] - '0')
				cubesInGame[getColor(colorsFirstLetter)] = countOfCubes
			}
		}
	}

	return true
}

func isGameStillPossible(cubesInGame map[string]int) bool {
	possible := newPossibleGame()
	for color, count := range cubesInGame {
		if count > possible.Cubes[color] {
			return false
		}
	}
	return true
}

func getColor(char string) string {
	colors := map[string]string{
		"r": "red",
		"g": "green",
		"b": "blue",
	}

	return colors[char]
}

func getProductOfFewestCubeInEachColor(line string) int {
	cubesInGame := make(map[string]int)
	skipNext := false

	for idx := range line {
		if skipNext {
			skipNext = false
			continue
		}

		// If we see a number it is at most 2 digits long
		// and after the last digit there is a space
		// and then the color, r, g or b
		var countOfCubes int
		if unicode.IsDigit(rune(line[idx])) {
			if unicode.IsDigit(rune(line[idx+1])) {
				colorsFirstLetter := string(line[idx+3])
				countOfCubes = int(line[idx]-'0')*10 + int(line[idx+1]-'0')

				_, ok := cubesInGame[getColor(colorsFirstLetter)]
				if !ok {
					cubesInGame[getColor(colorsFirstLetter)] = countOfCubes
				}

				if IsMoreCubes(cubesInGame, getColor(colorsFirstLetter), countOfCubes) {
					cubesInGame[getColor(colorsFirstLetter)] = countOfCubes
				}

				// we need to skip an iteration if a number is
				// 2 digits long
				skipNext = true
				continue
			} else {
				colorsFirstLetter := string(line[idx+2])
				countOfCubes = int(line[idx] - '0')

				_, ok := cubesInGame[getColor(colorsFirstLetter)]
				if !ok {
					cubesInGame[getColor(colorsFirstLetter)] = countOfCubes
				}

				if IsMoreCubes(cubesInGame, getColor(colorsFirstLetter), countOfCubes) {
					cubesInGame[getColor(colorsFirstLetter)] = countOfCubes
				}
			}
		}
	}
	product := productOfCubes(cubesInGame)

	return product
}

func IsMoreCubes(cubesInGame map[string]int, color string, countOfCubes int) bool {
	return countOfCubes > cubesInGame[color]
}

func productOfCubes(cubesInGame map[string]int) int {
	product := 0
	for _, count := range cubesInGame {
		if product == 0 {
			product = count
		} else {
			product *= count
		}
	}

	return product
}
