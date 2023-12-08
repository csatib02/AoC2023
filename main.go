package main

import (
	"fmt"

	"AoC/days/day_1"
	"AoC/days/day_2"
	"AoC/days/day_3"
	"AoC/days/day_4"
	"AoC/days/day_5"
	"AoC/days/day_6"
	"AoC/util"
)

func main() {
	getResultForDay(5)
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
	case 4:
		fmt.Printf("Day 4.1: %v\n", day_4.T_4_1(lines))
		fmt.Printf("Day 4.2: %v\n", day_4.T_4_2(lines))
	case 5:
		fmt.Printf("Day 5.1: %v\n", day_5.T_5_1(lines))
		fmt.Printf("Day 5.2: %v\n", day_5.T_5_2(lines))
	case 6:
		fmt.Printf("Day 6.1: %v\n", day_6.T_6_1(lines))
		fmt.Printf("Day 6.2: %v\n", day_6.T_6_2(lines))
	}
}
