package day_9

import (
	"sort"
	"strconv"
	"strings"
)

type Difference struct {
	order  int
	values []int
}

func T_9_1(lines []string) int {
	sum := 0
	for _, line := range lines {
		values := strings.Split(line, " ")
		extrapolatedValue := processLine(values, false)
		sum += extrapolatedValue
	}

	return sum
}

func T_9_2(lines []string) int {
	sum := 0
	for _, line := range lines {
		values := strings.Split(line, " ")
		extrapolatedValue := processLine(values, true)
		sum += extrapolatedValue
	}

	return sum
}

func processLine(values []string, part2 bool) int {
	differences := []Difference{}
	order := 1
	extrapolatedValue := 0
	nextValues := convertToInt(values)
	for {
		difference := newDifference(order, nextValues)
		order++
		differences = append(differences, difference)
		if checkIfAllZero(nextValues) {
			if part2 {
				extrapolatedValue = getExtrapolatedValueP2(differences)
			} else {
				extrapolatedValue = getExtrapolatedValue(differences)
			}
			break
		}

		nextValues = getNextInput(nextValues)
	}

	return extrapolatedValue
}

func newDifference(order int, nextValues []int) Difference {
	return Difference{order: order, values: nextValues}
}

func convertToInt(values []string) []int {
	var intValues []int
	for _, value := range values {
		intValue, _ := strconv.Atoi(value)
		intValues = append(intValues, intValue)
	}

	return intValues
}

func checkIfAllZero(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}

	return true
}

func getNextInput(values []int) []int {
	var differences []int
	for i, value := range values {
		if i == len(values)-1 {
			break
		}

		firstValue := value
		secondValue := values[i+1]
		difference := getdifference(firstValue, secondValue)
		differences = append(differences, difference)
	}

	return differences
}

func getdifference(firstValue, secondValue int) int {
	return secondValue - firstValue
}

func getExtrapolatedValue(differences []Difference) int {
	differences = sortslice(differences)
	extrapolatedValue := 0
	for i := len(differences) - 1; i >= 0; i-- {
		if i == len(differences)-1 {
			differences[i].values = append(differences[i].values, 0)
			continue
		}
		extrapolate := addExtrapolatedValue(differences[i], differences[i+1])
		differences[i].values = append(differences[i].values, extrapolate)
	}
	extrapolatedValue = differences[0].values[len(differences[0].values)-1]

	return extrapolatedValue
}

func sortslice(differences []Difference) []Difference {
	sort.Slice(differences, func(i, j int) bool {
		return differences[i].order < differences[j].order
	})

	return differences
}

func addExtrapolatedValue(difference1, difference2 Difference) int {
	return difference1.values[len(difference1.values)-1] + difference2.values[len(difference2.values)-1]
}

func getExtrapolatedValueP2(differences []Difference) int {
	differences = sortslice(differences)
	extrapolatedValue := 0
	for i := range differences {
		differences[i].values = reverseSlice(differences[i].values)
	}

	for i := len(differences) - 1; i >= 0; i-- {
		if i == len(differences)-1 {
			differences[i].values = append(differences[i].values, 0)
			continue
		}
		extrapolate := addExtrapolatedValueBehind(differences[i+1], differences[i])
		differences[i].values = append(differences[i].values, extrapolate)
	}
	extrapolatedValue = differences[0].values[len(differences[0].values)-1]

	return extrapolatedValue
}

func reverseSlice(nextValues []int) []int {
	reversed := make([]int, len(nextValues))

	for i, value := range nextValues {
		reversed[len(nextValues)-1-i] = value
	}

	return reversed
}

func addExtrapolatedValueBehind(difference1, difference2 Difference) int {
	value1 := difference1.values[len(difference1.values)-1]
	value2 := difference2.values[len(difference2.values)-1]
	var value int
	if value1 > value2 {
		value = value1 - value2
	} else {
		value = value2 - value1
	}

	if value1+value != value2 {
		value = value * -1
	}

	return value
}
