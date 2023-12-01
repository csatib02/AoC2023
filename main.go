package main

import (
	"fmt"
	"os"
	"strings"

	"AoC/days"
)

func main() {
	getResultForDay(1)
}

func getResultForDay(day int) {
	lines, _ := getData(fmt.Sprintf("inputs/day_%v.txt", day))

	switch day {
	case 1:
		fmt.Printf("Day 1.1: %v\n", days.T_1_1(lines))
		fmt.Printf("Day 1.2: %v\n", days.T_1_2(lines))
	case 2:
		fmt.Printf("Day 2.1: %v\n", days.T_2_1(lines))
		fmt.Printf("Day 2.2: %v\n", days.T_2_2(lines))

	}
}

func getData(filepath string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("file reading error %v", err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	return lines, nil
}
