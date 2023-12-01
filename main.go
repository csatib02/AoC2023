package main

import (
	"AoC/days/day_1"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, err := getData("days/day_1/day_1_input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	result1 := day_1.T_1(lines)
	result2 := day_1.T_2(lines)

	fmt.Printf("Day 1.1: %v\n", result1)
	fmt.Printf("Day 1.2: %v\n", result2)
}

func getData(filepath string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("file reading error %v", err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	return lines, nil
}
