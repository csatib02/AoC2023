package day_8

import (
	"regexp"
	"strings"
)

type Instruction struct {
	left  string
	right string
}

func newInstruction(line string) *map[string]Instruction {
	instruction := make(map[string]Instruction)
	split := strings.Split(line, " ")
	name := split[0]

	split[2] = strings.TrimLeft(split[2], "(")
	split[2] = strings.TrimRight(split[2], ",")
	left := split[2]

	split[3] = strings.TrimRight(split[3], ")")
	right := split[3]

	instruction[name] = Instruction{left, right}

	return &instruction
}

func T_8_1(lines []string) int {
	steps := 0
	var directions []string
	var instructions []*map[string]Instruction
	for i, line := range lines {
		if i == 0 {
			directions = strings.Split(line, "")
		}
		if i > 1 {
			instruction := newInstruction(line)
			instructions = append(instructions, instruction)
		}
	}

	steps = stepsToReachZZZ(directions, instructions, steps)

	return steps
}

func T_8_2(lines []string) int {
	var directions []string
	var instructions []*map[string]Instruction
	for i, line := range lines {
		if i == 0 {
			directions = strings.Split(line, "")
		}
		if i > 1 {
			instruction := newInstruction(line)
			instructions = append(instructions, instruction)
		}
	}

	currents := getInstructionPart2(instructions, regexp.MustCompile(`.*A$`))
	minEndZ := minStepsToFindEndingInZ(currents, directions, instructions)
	result := findlcm(minEndZ)

	return result
}

func getNext(current string, idx int, directions []string, instructions []*map[string]Instruction) string {
	var next string

	currentInstruction := getInstruction(current, instructions)

	if directions[idx] == "R" {
		next = currentInstruction.right
	} else {
		next = currentInstruction.left
	}

	return next
}

func getInstruction(current string, instructions []*map[string]Instruction) *Instruction {
	for _, instruction := range instructions {
		if instruction, ok := (*instruction)[current]; ok {
			return &instruction
		}
	}
	return nil
}

func getInstructionPart2(instructions []*map[string]Instruction, re *regexp.Regexp) []string {
	var currents []string
	for _, instruction := range instructions {
		for key := range *instruction {
			if re.MatchString(key) {
				currents = append(currents, key)
			}
		}
	}

	return currents
}

func stepsToReachZZZ(directions []string, instructions []*map[string]Instruction, steps int) int {
	current := "AAA"
	idx := 0
	var next string
	for current != "ZZZ" {
		if idx == len(directions) {
			idx = 0
		}

		next = getNext(current, idx, directions, instructions)
		current = next

		steps++
		idx++
	}

	return steps
}

func minStepsToFindEndingInZ(currents []string, directions []string, instructions []*map[string]Instruction) []int {
	idx := 0
	stepforOne := 0
	var currentOne string
	var next string
	var minEndZ []int

	for _, current := range currents {
		currentOne = current
		for {
			if idx == len(directions) {
				idx = 0
			}

			next = getNext(currentOne, idx, directions, instructions)
			if endInZ(next) {
				stepforOne++
				minEndZ = append(minEndZ, stepforOne)
				break
			}
			currentOne = next
			stepforOne++
			idx++
		}
		stepforOne = 0
		idx = 0
	}

	return minEndZ
}

func endInZ(current string) bool {
	re := regexp.MustCompile(`.*Z$`)

	return re.MatchString(current)
}

func findlcm(lcm []int) int {
	result := lcm[0]
	for i := 1; i < len(lcm); i++ {
		result = findlcmTwo(result, lcm[i])
	}

	return result
}

func findlcmTwo(result, i int) int {
	return (result * i) / findGCD(result, i)
}

func findGCD(result, i int) int {
	if result == 0 {
		return i
	}

	return findGCD(i%result, result)
}
