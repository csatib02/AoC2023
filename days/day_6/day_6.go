package day_6

import (
	"regexp"
	"strconv"
	"strings"
)

type Races struct {
	allRaces []*Race
}

type Race struct {
	Time     int
	Distance int
}

func T_6_1(lines []string) int {
	var product int
	times, _ := getData(lines[0], false)
	distances, _ := getData(lines[1], false)
	races := NewRaces(times, distances)

	records := getRecordBreaks(races)
	product = getProduct(records)

	return product
}

func T_6_2(lines []string) int {
	var product int
	var races Races
	_, time := getData(lines[0], true)
	_, distance := getData(lines[1], true)

	race := NewRace(time, distance)
	races.allRaces = append(races.allRaces, race)

	records := getRecordBreaks(races)
	product = getProduct(records)

	return product
}

func getData(s string, part2 bool) (map[int]int, int) {
	times := make(map[int]int)
	order := 1
	split := strings.Split(s, " ")
	re := regexp.MustCompile(`\d{1,4}`)

	var badlyKernedNumber []string

	for _, v := range split {
		numbers := re.FindAllString(v, -1)
		if len(numbers) == 1 {
			if part2 {
				badlyKernedNumber = append(badlyKernedNumber, numbers[0])
				continue
			}

			time, _ := strconv.Atoi(numbers[0])
			times[order] = time
			order++
		}
	}
	formedNumber := strings.Join(badlyKernedNumber, "")
	number, _ := strconv.Atoi(formedNumber)

	return times, number
}

func NewRaces(times, distances map[int]int) Races {
	var races Races
	for order, time := range times {
		distance, ok := distances[order]
		if ok {
			race := NewRace(time, distance)
			races.allRaces = append(races.allRaces, race)
		}
	}

	return races
}

func NewRace(time, distance int) *Race {
	return &Race{Time: time, Distance: distance}
}

func getRecordBreaks(races Races) []int {
	var sumOfRecordBreaksPerRace []int
	for _, race := range races.allRaces {
		possibleDistances := runRace(race)
		numberOfRecordBreaks := evalRaceOptions(possibleDistances, race.Distance)
		sumOfRecordBreaksPerRace = append(sumOfRecordBreaksPerRace, numberOfRecordBreaks)
	}

	return sumOfRecordBreaksPerRace
}

func runRace(race *Race) []int {
	var possibleDistances []int

	raceTime := race.Time
	time := 0
	for {
		if time == raceTime {
			break
		}
		distance := time * (raceTime - time)
		possibleDistances = append(possibleDistances, distance)

		time++
	}

	return possibleDistances
}

func evalRaceOptions(currentRaceDistances []int, raceDistance int) int {
	var numberOfRecordBreaks int
	for _, distance := range currentRaceDistances {
		if distance > raceDistance {
			numberOfRecordBreaks++
		}
	}

	return numberOfRecordBreaks
}

func getProduct(record []int) int {
	product := 1
	for _, v := range record {
		product *= v
	}

	return product
}
