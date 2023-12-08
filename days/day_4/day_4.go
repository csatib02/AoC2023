package day_4

import (
	"regexp"
	"strconv"
	"strings"
)

type Cards struct {
	AllCards []Card
}

type Card struct {
	Order                int
	CountOfWinnerNumbers int
}

func T_4_1(lines []string) int {
	sum := 0
	for _, line := range lines {
		line, _ = ProcessOrderNumber(line, false)
		winningNumbers, numbersHave := splitBySeperator(line)
		winningNumbersMap := makeMapOfNumbers(winningNumbers)
		numbersHaveMap := makeMapOfNumbers(numbersHave)

		countOfWinnerNumbers := checkForWinningNumbers(numbersHaveMap, winningNumbersMap)
		sum += calculatePoints(countOfWinnerNumbers)
	}

	return sum
}

func T_4_2(lines []string) int {
	sum := 0
	cards := Cards{}
	for _, line := range lines {
		// Make a struct of cards, so we can easily find a card by order number
		line, card := ProcessOrderNumber(line, true)
		winningNumbers, numbersHave := splitBySeperator(line)
		winningNumbersMap := makeMapOfNumbers(winningNumbers)
		numbersHaveMap := makeMapOfNumbers(numbersHave)
		card.CountOfWinnerNumbers = checkForWinningNumbers(numbersHaveMap, winningNumbersMap)

		cards.AllCards = append(cards.AllCards, card)
	}

	for _, card := range cards.AllCards {
		// Instead of points we get more cards below
		// the order number of the winning card
		// the amount of cards won is determined by the count of winning numbers
		scratchcardsWon := cards.processCard(card)
		sum += scratchcardsWon
	}

	return sum
}

func ProcessOrderNumber(line string, part2 bool) (string, Card) {
	split := strings.SplitN(line, ":", 2)
	split[1] = strings.TrimLeft(split[1], " ")

	if part2 {
		re := regexp.MustCompile(`\d{1,3}`)
		Card := Card{}
		Card.Order, _ = strconv.Atoi(re.FindAllString(split[0], -1)[0])

		return split[1], Card
	}

	return split[1], Card{}
}

func splitBySeperator(line string) ([]string, []string) {
	re := regexp.MustCompile(`\d{1,2}`)
	split := strings.SplitN(line, "|", 2)

	winningNumbers := re.FindAllString(split[0], -1)
	numbersHave := re.FindAllString(split[1], -1)

	return winningNumbers, numbersHave
}

func makeMapOfNumbers(numbers []string) map[string]int {
	numbersMap := make(map[string]int)
	for _, number := range numbers {
		intNumber, _ := strconv.Atoi(number)
		numbersMap[number] = intNumber
	}

	return numbersMap
}

func checkForWinningNumbers(numbersHaveMap, winningNumbersMap map[string]int) int {
	countOfWinnerNumbers := 0
	for number := range numbersHaveMap {
		if _, ok := winningNumbersMap[number]; ok {
			countOfWinnerNumbers++
		}
	}

	return countOfWinnerNumbers
}

func calculatePoints(countOfWinnerNumbers int) int {
	if countOfWinnerNumbers > 0 {
		return 1 << (countOfWinnerNumbers - 1)
	}

	return 0
}

func (Cards Cards) findByOrderNumber(orderNumber int) Card {
	for _, card := range Cards.AllCards {
		if card.Order == orderNumber {
			return card
		}
	}

	return Card{}
}

func (Cards Cards) processCard(card Card) int {
	orderNumber := card.Order
	count := 1
	if card.CountOfWinnerNumbers > 0 {
		// Get all cards below the winning card
		for i := 0; i < card.CountOfWinnerNumbers; i++ {
			orderNumber++
			count += Cards.processCard(Cards.findByOrderNumber(orderNumber))
		}

		return count
	}

	return count
}
