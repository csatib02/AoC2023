package day_7

import (
	"sort"
	"strconv"
	"strings"
)

type AllHands struct {
	AllHands []*Hand
	Length   *int
}

type Hand struct {
	Cards          string
	Bid            int
	Strength       string
	RemainingCards string
	Rank           int
}

type StrengthCounter struct {
	Selector map[string]int
}

type Strength struct {
	Groups []map[string][]*Hand
}

type StrengthOfCards struct {
	Strength map[string]int
}

func NewStrengthCounter() *StrengthCounter {
	return &StrengthCounter{
		Selector: map[string]int{
			"2": 0,
			"3": 0,
			"4": 0,
			"5": 0,
			"6": 0,
			"7": 0,
			"8": 0,
			"9": 0,
			"T": 0,
			"J": 0,
			"Q": 0,
			"K": 0,
			"A": 0,
		},
	}
}

func newStrengthOfCards(part2 bool) *StrengthOfCards {
	if part2 {
		return &StrengthOfCards{
			Strength: map[string]int{
				"J": 1,
				"2": 2,
				"3": 3,
				"4": 4,
				"5": 5,
				"6": 6,
				"7": 7,
				"8": 8,
				"9": 9,
				"T": 10,
				"Q": 11,
				"K": 12,
				"A": 13,
			},
		}
	}

	return &StrengthOfCards{
		Strength: map[string]int{
			"2": 1,
			"3": 2,
			"4": 3,
			"5": 4,
			"6": 5,
			"7": 6,
			"8": 7,
			"9": 8,
			"T": 9,
			"J": 10,
			"Q": 11,
			"K": 12,
			"A": 13,
		},
	}
}

func NewStrengthGroup() *Strength {
	strengths := map[int]string{
		1: "Five of a kind",
		2: "Four of a kind",
		3: "Full house",
		4: "Three of a kind",
		5: "Two pair",
		6: "One pair",
		7: "High card",
	}

	strengthGrps := &Strength{}

	var keys []int
	for key := range strengths {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		cardGroup := strengths[key]
		mapForGroup := map[string][]*Hand{}
		mapForGroup[cardGroup] = []*Hand{}
		strengthGrps.Groups = append(strengthGrps.Groups, mapForGroup)
	}

	return strengthGrps
}

func T_7_1(lines []string) int {
	sum := 0
	var hands AllHands

	for _, line := range lines {
		split := strings.Split(line, " ")
		hand := NewHand(split[0], split[1])
		hands.AllHands = append(hands.AllHands, hand)
	}
	lengthOfHands := len(hands.AllHands)
	strengthGroup := NewStrengthGroup()

	for _, hand := range hands.AllHands {
		hand.addStrength(*strengthGroup, false)
	}

	// Order and rank within groups
	for _, group := range strengthGroup.Groups {
		for label := range group {
			lengthOfHands = orderAndRankWithinGroup(group, lengthOfHands, label, false)
		}
	}

	// Calculate the sum of bids multiplied by ranks
	for _, hand := range hands.AllHands {
		sum += hand.Bid * hand.Rank
	}

	return sum
}

func T_7_2(lines []string) int {
	sum := 0
	var hands AllHands

	for _, line := range lines {
		split := strings.Split(line, " ")
		hand := NewHand(split[0], split[1])
		hands.AllHands = append(hands.AllHands, hand)
	}
	lengthOfHands := len(hands.AllHands)
	strengthGroup := NewStrengthGroup()

	for _, hand := range hands.AllHands {
		hand.addStrength(*strengthGroup, true)
	}

	// Order and rank within groups
	for _, group := range strengthGroup.Groups {
		for label := range group {
			lengthOfHands = orderAndRankWithinGroup(group, lengthOfHands, label, true)
		}
	}

	for _, hand := range hands.AllHands {
		sum += hand.Bid * hand.Rank
	}

	return sum
}

func NewHand(cards, bid string) *Hand {
	bidInt, _ := strconv.Atoi(bid)
	return &Hand{
		Cards: cards,
		Bid:   bidInt,
	}
}

func (hand *Hand) addStrength(strengthGroup Strength, part2 bool) {
	strengthCounter := NewStrengthCounter()
	var strength string
	var labelCombinations []string
	for _, card := range hand.Cards {
		strengthCounter.Selector[string(card)]++
	}

	if part2 {
		// we have to deal with jokers
		if strengthCounter.Selector["J"] > 0 {
			strength, labelCombinations = strengthCounter.HighestCombinations(part2)
		} else {
			// if there aren't any jokers
			// then we can just use everything as is
			strength, labelCombinations = strengthCounter.HighestCombinations(false)
		}
	} else {
		strength, labelCombinations = strengthCounter.HighestCombinations(part2)
	}

	hand.Strength = strength
	strengthGroup.addHand(strength, hand)
	hand.updateHand(labelCombinations)
}

func (strength *Strength) addHand(strengthLabel string, hand *Hand) {
	for _, group := range strength.Groups {
		for label, hands := range group {
			if label == strengthLabel {
				group[label] = append(hands, hand)
			}
		}
	}
}

func (sc *StrengthCounter) HighestCombinations(part2 bool) (string, []string) {
	highestCombination, labelOfCombination := sc.getHighestCombination(false)
	var labelCombinations []string
	labelCombinations = append(labelCombinations, labelOfCombination)

	// we have to deal with jokers
	if part2 {
		// 5 of a kind
		if sc.Selector["J"] == 5 {
			return "Five of a kind", labelCombinations
		}

		// if we have 4 jokers
		// they make 5 of a kind
		if sc.Selector["J"] == 4 {
			// only do this to add the other card since everything is stronger than jokers
			labelOfStrongestCard := findStrongestCardInHand(sc.Selector)
			labelCombinations = append(labelCombinations, labelOfStrongestCard)

			// joker cards help the best possible outcome
			sc.Selector[labelOfStrongestCard] += sc.Selector["J"]
			sc.Selector["J"] = 0
			return "Five of a kind", labelCombinations
		}

		if sc.Selector["J"] == 3 {
			ok, label := sc.checkforCertainCombination(2, false)
			// if we have 2 of the same cards and 3 jokers
			// we have 5 of a kind
			if ok {
				labelCombinations = append(labelCombinations, label)
				// jokers help the best possible outcome
				sc.Selector[label] += sc.Selector["J"]
				sc.Selector["J"] = 0
				return "Five of a kind", labelCombinations
			}
			// if we have 3 jokers and 2 different cards
			// we need to check the strongest card
			// but we have 4 of a kind
			labelOfStrongestCard := findStrongestCardInHand(sc.Selector)
			labelCombinations = append(labelCombinations, labelOfStrongestCard)
			// jokers help the best possible outcome
			sc.Selector[labelOfStrongestCard] += sc.Selector["J"]
			sc.Selector["J"] = 0
			return "Four of a kind", labelCombinations
		}

		if sc.Selector["J"] == 2 {
			// if we have 2 jokers and 3 of the same cards
			// we have 5 of a kind
			ok, label := sc.checkforCertainCombination(3, false)
			if ok {
				// since we have 3 of the same cards
				// the jokers aren't added to the combination
				labelCombinations = append(labelCombinations, "J")
				// jokers help the best possible outcome
				sc.Selector[label] += sc.Selector["J"]
				sc.Selector["J"] = 0
				return "Five of a kind", labelCombinations
			}
			ok, label = sc.checkforCertainCombination(2, true)
			// if we have 2 of the same cards and 2 jokers
			// we need to check the strongest card
			// but we have 4 of a kind
			if ok {
				// we don't know what has been added to the combinations
				labelCombinations = []string{}
				labelCombinations = append(labelCombinations, "J")
				labelCombinations = append(labelCombinations, label)
				// jokers help the best possible outcome
				sc.Selector[label] += sc.Selector["J"]
				sc.Selector["J"] = 0
				return "Four of a kind", labelCombinations
			}
			// if we have 2 jokers and 2 different cards
			// we need to check the strongest card
			// but we have 3 of a kind
			labelOfStrongestCard := findStrongestCardInHand(sc.Selector)
			labelCombinations = append(labelCombinations, labelOfStrongestCard)
			// jokers help the best possible outcome
			sc.Selector[labelOfStrongestCard] += sc.Selector["J"]
			sc.Selector["J"] = 0
			return "Three of a kind", labelCombinations
		}

		if sc.Selector["J"] == 1 {
			// if we have 1 joker and 4 of the same cards
			// we have 5 of a kind
			ok, label := sc.checkforCertainCombination(4, false)
			if ok {
				// since we have 4 of the same cards
				// the joker isn't added to the combination
				labelCombinations = append(labelCombinations, "J")
				// jokers help the best possible outcome
				sc.Selector[label] += sc.Selector["J"]
				sc.Selector["J"] = 0
				return "Five of a kind", labelCombinations
			}
			ok, label = sc.checkforCertainCombination(3, false)
			// if we have 1 joker and 3 of the same cards
			// we need to check the strongest card
			// but we have 4 of a kind
			if ok {
				labelCombinations = append(labelCombinations, "J")
				// jokers help the best possible outcome
				sc.Selector[label] += sc.Selector["J"]
				sc.Selector["J"] = 0
				return "Four of a kind", labelCombinations
			}
			// we need to check for full house
			ok, label = sc.checkforCertainCombination(2, false)
			if ok {
				labelCombinations = []string{}
				labelCombinations = append(labelCombinations, "J")
				labelCombinations = append(labelCombinations, label)
				for otherLabel, value := range sc.Selector {
					if value == 2 && label != otherLabel {
						labelCombinations = append(labelCombinations, otherLabel)
						labelOfStrongestCard := findStrongestCardInHand(sc.Selector)
						// joke helps the best possible outcome
						sc.Selector[labelOfStrongestCard] += sc.Selector["J"]
						sc.Selector["J"] = 0
						return "Full house", labelCombinations
					}
				}
				// we have 1 joker and 2 of the same cards and 3 different cards
				// 3 of a kind

				// jokers help the best possible outcome
				sc.Selector[label] += sc.Selector["J"]
				sc.Selector["J"] = 0
				return "Three of a kind", labelCombinations
			}
			// if we have 1 joker and 4 different cards
			// we need to check the strongest card
			// but we have one pair
			labelOfStrongestCard := findStrongestCardInHand(sc.Selector)
			labelCombinations = append(labelCombinations, labelOfStrongestCard)
			// jokers help the best possible outcome
			sc.Selector[labelOfStrongestCard] += sc.Selector["J"]
			sc.Selector["J"] = 0
			return "One pair", labelCombinations
		}
	}

	// 5 of a kind
	if highestCombination == 5 {
		return "Five of a kind", labelCombinations
	}

	// 4 of a kind
	if highestCombination == 4 {
		return "Four of a kind", labelCombinations
	}
	// full house
	if highestCombination == 3 {
		ok, label := sc.checkforCertainCombination(2, false)
		// 3 of a kind
		if !ok {
			return "Three of a kind", labelCombinations
		}
		labelCombinations = append(labelCombinations, label)
		return "Full house", labelCombinations
	}

	// 2 pairs
	if highestCombination == 2 {
		sc.Selector[labelOfCombination] = 0
		ok, label := sc.checkforCertainCombination(2, false)
		// 1 pair
		if !ok {
			return "One pair", labelCombinations
		}
		labelCombinations = append(labelCombinations, label)
		return "Two pair", labelCombinations
	}

	// high card
	return "High card", nil
}

func findStrongestCardInHand(sc map[string]int) string {
	// we can recreate a hand from the map
	// and then find the strongest card
	StrengthOfCards := newStrengthOfCards(true)
	var hand []string
	for label, value := range sc {
		for i := 0; i < value; i++ {
			hand = append(hand, label)
		}
	}

	sort.Slice(hand, func(i, j int) bool {
		return StrengthOfCards.Strength[hand[i]] > StrengthOfCards.Strength[hand[j]]
	})

	return hand[0]
}

func (sc *StrengthCounter) getHighestCombination(part2 bool) (int, string) {
	highestCombination := 0
	labelOfCombination := ""

	if part2 {
		for label, value := range sc.Selector {
			if value > highestCombination && label != "J" {
				highestCombination = value
				labelOfCombination = label
			}
		}

		return highestCombination, labelOfCombination
	}

	for label, value := range sc.Selector {
		if value > highestCombination {
			highestCombination = value
			labelOfCombination = label
		}
	}

	return highestCombination, labelOfCombination
}

func (sc *StrengthCounter) checkforCertainCombination(combination int, nonJoker bool) (bool, string) {
	// if non joker we want to omit the jokers
	if nonJoker {
		for label, value := range sc.Selector {
			if value == combination && label != "J" {
				return true, label
			}
		}
	}

	for label, value := range sc.Selector {
		if value == combination {
			return true, label
		}
	}

	return false, ""
}

func (hand *Hand) updateHand(labelCombinations []string) {
	updatedHand := hand.Cards

	for _, label := range labelCombinations {
		updatedHand = strings.Replace(updatedHand, label, "", -1)
	}
	hand.RemainingCards = updatedHand
}

func orderAndRankWithinGroup(group map[string][]*Hand, length int, groupLabel string, part2 bool) int {
	sort.Slice(group[groupLabel], func(i, j int) bool {
		return compareHands(group[groupLabel][i], group[groupLabel][j], part2)
	})

	for i, hand := range group[groupLabel] {
		hand.Rank = length - i
	}

	sort.Slice(group[groupLabel], func(i, j int) bool {
		if group[groupLabel][i].Strength != group[groupLabel][j].Strength {
			return false
		}
		return group[groupLabel][i].RemainingCards > group[groupLabel][j].RemainingCards
	})

	return length - len(group[groupLabel])
}

func compareHands(hand1, hand2 *Hand, part2 bool) bool {
	strengthMap := newStrengthOfCards(part2).Strength

	// first, order hands within the current group
	// based on the order of the combination of strength
	for i := 0; i < 5; i++ {
		card1Strength := strengthMap[string(hand1.Cards[i])]
		card2Strength := strengthMap[string(hand2.Cards[i])]

		// compare the strengths of the cards
		if card1Strength != card2Strength {
			return card1Strength > card2Strength
		}
	}

	// if the strengths are the same, compare the remaining cards
	for i := 0; i < 5; i++ {
		if hand1.Cards[i] != hand2.Cards[i] {
			return hand1.Cards[i] > hand2.Cards[i]
		}
	}

	// no swap
	return false

}
