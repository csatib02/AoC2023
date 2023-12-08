package day_5

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Maps struct {
	AllMaps []*Map
}

type Map struct {
	Name      string
	Order     int
	Converter []*Converter
}

type Converter struct {
	Destination      int
	Source           int
	Length           int
	numbersInBetween []int
}

// Absolute bruteforce solution
func T_5_1(lines []string) int {
	var result int
	order := 0
	var seeds []int
	newMap := false
	var maps Maps
	var currentMap *Map
	var mapping []int

	for i, line := range lines {
		if i == 0 {
			seeds = getSeeds(line, false)
			continue
		}

		if strings.TrimSpace(line) != "" && unicode.IsLetter(rune(line[0])) {
			currentMap = createNewMap(line, order)
			maps.AllMaps = append(maps.AllMaps, currentMap)
			newMap = true
			continue
		}

		if newMap {
			if strings.TrimSpace(line) != "" {
				currentMap.addDataToMap(line, &seeds)
			}
		}

		if strings.TrimSpace(line) == "" {
			order++
			newMap = false
		}

	}

	maps.sortMaps()

	for i, m := range maps.AllMaps {
		if i == 0 {
			mapping = calculateMapping(m, &seeds, true)
			continue
		}
		mapping = calculateMapping(m, &mapping, false)
	}

	result = findMin(mapping)

	return result
}

// Absolute bruteforce solution
func T_5_2(lines []string) int {
	var result int
	order := 0
	var seeds []int
	newMap := false
	var maps Maps
	var currentMap *Map
	var mapping []int

	for i, line := range lines {
		if i == 0 {
			seeds = getSeeds(line, true)
			continue
		}

		if strings.TrimSpace(line) != "" && unicode.IsLetter(rune(line[0])) {
			currentMap = createNewMap(line, order)
			maps.AllMaps = append(maps.AllMaps, currentMap)
			newMap = true
			continue
		}

		if newMap {
			if strings.TrimSpace(line) != "" {
				currentMap.addDataToMap(line, &seeds)
			}
		}

		if strings.TrimSpace(line) == "" {
			order++
			newMap = false
		}

	}

	maps.sortMaps()

	for i, m := range maps.AllMaps {
		if i == 0 {
			mapping = calculateMapping(m, &seeds, true)
			continue
		}
		mapping = calculateMapping(m, &mapping, false)
	}
	result = findMin(mapping)

	return result
}

func getSeeds(line string, part2 bool) []int {
	var seeds []int
	split := strings.Split(line, " ")
	split = split[1:]
	if part2 {
		for i := 0; i < len(split); i += 2 {
			end_1, _ := strconv.Atoi(split[i])
			end_2, _ := strconv.Atoi(split[i+1])
			end := end_1 + end_2
			currentSeedSource, _ := strconv.Atoi(split[i])
			for currentSeed := currentSeedSource; currentSeed < end; currentSeed++ {
				seeds = append(seeds, currentSeed)
			}
		}
		return seeds
	}

	for _, s := range split {
		seedInt, _ := strconv.Atoi(s)
		seeds = append(seeds, seedInt)
	}

	return seeds
}

func createNewMap(line string, order int) *Map {
	split := strings.Split(line, " ")
	name := split[0]

	return &Map{Name: name, Order: order, Converter: []*Converter{}}
}

func (m *Map) addDataToMap(line string, seeds *[]int) {
	split := strings.Split(line, " ")

	var destinationInt, sourceInt, rangeInt int

	for i, s := range split {
		if i == 0 {
			destinationInt, _ = strconv.Atoi(s)
		}
		if i == 1 {
			sourceInt, _ = strconv.Atoi(s)
		}
		if i == 2 {
			rangeInt, _ = strconv.Atoi(s)
		}
	}

	converter := &Converter{Destination: destinationInt, Source: sourceInt, Length: rangeInt, numbersInBetween: []int{}}

	if m.Name == "seed-to-soil" {
		seedsInBetween := areInBetween(converter, seeds)
		if len(seedsInBetween) > 0 {
			converter.numbersInBetween = seedsInBetween
			m.Converter = append(m.Converter, converter)
		}
	} else {
		m.Converter = append(m.Converter, converter)
	}
}

func areInBetween(converter *Converter, seeds *[]int) []int {
	begin := converter.Source
	end := converter.Source + converter.Length
	var seedsInBetween []int
	for _, seed := range *seeds {
		if seed >= begin && seed < end {
			seedsInBetween = append(seedsInBetween, seed)
		}
	}

	return seedsInBetween
}

func (aMaps *Maps) sortMaps() {
	n := len(aMaps.AllMaps)
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if aMaps.AllMaps[j].Order > aMaps.AllMaps[j+1].Order {
				aMaps.AllMaps[j], aMaps.AllMaps[j+1] = aMaps.AllMaps[j+1], aMaps.AllMaps[j]
			}
		}
	}
}

func calculateMapping(currentMap *Map, mapping *[]int, first bool) []int {
	var newMapping []int
	for _, c := range currentMap.Converter {
		result := calculateMappingForConverter(c, mapping, first)
		if len(result) > 0 {
			newMapping = append(newMapping, result...)
		}
	}

	newMapping = append(newMapping, *mapping...)

	return newMapping
}

func calculateMappingForConverter(c *Converter, mapping *[]int, first bool) []int {
	source := c.Source
	end := c.Source + c.Length
	destination := c.Destination
	var foundElements []int

	for source < end {
		if foundMatching(c, mapping, source, first) {
			foundElements = append(foundElements, destination)
		}

		source++
		destination++
	}

	return foundElements
}

func foundMatching(c *Converter, mapping *[]int, currentSource int, first bool) bool {
	numbersInBetween := c.numbersInBetween
	if first {
		for i, num := range numbersInBetween {
			if num == currentSource {
				c.numbersInBetween = append(numbersInBetween[:i], numbersInBetween[i+1:]...)
				removeElementFromMapping(mapping, currentSource)
				return true
			}
		}
	} else {
		for i := range *mapping {
			if (*mapping)[i] == currentSource {
				removeElementFromMapping(mapping, currentSource)
				return true
			}
		}
	}

	return false
}

func removeElementFromMapping(mapping *[]int, currentSource int) {
	for i, num := range *mapping {
		if num == currentSource {
			*mapping = append((*mapping)[:i], (*mapping)[i+1:]...)
			return
		}
	}
}

func findMin(mapping []int) int {
	sort.Ints(mapping)

	return mapping[0]
}
