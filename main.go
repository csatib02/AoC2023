package main

import (
	"fmt"

	"AoC/days/day_1"
	"AoC/days/day_2"
	"AoC/days/day_3"
	"AoC/util"
)

func main() {
	getResultForDay(3)
}

func getResultForDay(day int) {
	lines, _ := util.GetData(fmt.Sprintf("inputs/day_%v.txt", day))

	switch day {
	case 1:
		fmt.Printf("Day 1.1: %v\n", day_1.T_1_1(lines))
		fmt.Printf("Day 1.2: %v\n", day_1.T_1_2(lines))
	case 2:
		fmt.Printf("Day 2.1: %v\n", day_2.T_2_1(lines))
		fmt.Printf("Day 2.2: %v\n", day_2.T_2_2(lines))
	case 3:
		fmt.Printf("Day 3.1: %v\n", day_3.T_3_1(lines))
		fmt.Printf("Day 3.2: %v\n", day_3.T_3_2(lines))
	}
}
