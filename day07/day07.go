package day07

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day07/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var pokerHands = []PokerHand{}

	for scanner.Scan() {
		line := scanner.Text()
		var cards, bid = ParsePokerHandString(line)
		pokerHands = append(pokerHands, NewPokerHand(cards, bid))
	}

	SortPokerHands(pokerHands)
	var totalWinnings = CalculateTotalWinnings(pokerHands)

	answer = totalWinnings
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day07/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var pokerHands = []PokerHand{}

	for scanner.Scan() {
		line := scanner.Text()
		var cards, bid = ParsePokerHandString(line)
		pokerHands = append(pokerHands, NewPokerHandWithWilds(cards, bid))
	}

	SortPokerHandsWithWilds(pokerHands)
	var totalWinnings = CalculateTotalWinnings(pokerHands)

	answer = totalWinnings
	return answer, time.Since(now)
}

type HandType int

const (
	FiveOfAKind  HandType = 6
	FourOfAKind  HandType = 5
	FullHouse    HandType = 4
	ThreeOfAKind HandType = 3
	TwoPair      HandType = 2
	OnePair      HandType = 1
	HighCard     HandType = 0
)

type PokerHand struct {
	Type  HandType
	Cards string
	Bid   int
}

func ParsePokerHandString(line string) (cards string, bid int) {
	// 32T3K 765
	var fields = strings.Fields(line)
	cards = fields[0]
	bid, _ = strconv.Atoi(fields[1])
	return
}

func NewPokerHand(cards string, bid int) PokerHand {
	var hand = PokerHand{}
	hand.Cards = cards
	hand.Bid = bid

	var counter = map[rune]int{}
	for _, card := range cards {
		counter[card] = counter[card] + 1
	}

	if len(counter) == 1 {
		hand.Type = FiveOfAKind
	} else if len(counter) == 2 {
		// Four of a kind or a full house
		for _, v := range counter {
			if v == 4 || v == 1 {
				hand.Type = FourOfAKind
				break
			} else if v == 3 || v == 2 {
				hand.Type = FullHouse
			}
		}
	} else if len(counter) == 3 {
		// Three of a kind or two pair
		for _, v := range counter {
			if v == 2 {
				hand.Type = TwoPair
				break
			} else if v == 3 || v == 1 {
				hand.Type = ThreeOfAKind
			}
		}
	} else if len(counter) == 4 {
		hand.Type = OnePair
	} else {
		hand.Type = HighCard
	}

	return hand
}

func NewPokerHandWithWilds(cards string, bid int) PokerHand {
	var hand = NewPokerHand(cards, bid)

	var counter = map[rune]int{}
	for _, card := range cards {
		counter[card] = counter[card] + 1
	}

	var count = counter[rune('J')]

	// upgrade hand for each Joker (wild)
	for i := 0; i < count; i++ {
		switch hand.Type {
		case HighCard:
			hand.Type = OnePair
		case OnePair:
			hand.Type = ThreeOfAKind
		case TwoPair:
			hand.Type = FullHouse
		case ThreeOfAKind:
			hand.Type = FourOfAKind
		case FullHouse:
			hand.Type = FourOfAKind
		case FourOfAKind:
			hand.Type = FiveOfAKind
		}
	}

	return hand
}

var cardRanks = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardRanksWithWild = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func SortPokerHands(hands []PokerHand) {
	sort.Slice(hands, func(i, j int) bool {
		var leftHand = hands[i]
		var rightHand = hands[j]

		if leftHand.Type < rightHand.Type {
			return true
		} else if leftHand.Type > rightHand.Type {
			return false
		}

		// Types are equal, compare cards within hands
		for i = 0; i < len(leftHand.Cards); i++ {
			var leftCard = rune(leftHand.Cards[i])
			var rightCard = rune(rightHand.Cards[i])

			if cardRanks[leftCard] < cardRanks[rightCard] {
				return true
			} else if cardRanks[leftCard] > cardRanks[rightCard] {
				return false
			}

			// card ranks are equal, check the next card
		}

		// if we got this far, then the hands are equal
		return false
	})
}

func SortPokerHandsWithWilds(hands []PokerHand) {
	sort.Slice(hands, func(i, j int) bool {
		var leftHand = hands[i]
		var rightHand = hands[j]

		if leftHand.Type < rightHand.Type {
			return true
		} else if leftHand.Type > rightHand.Type {
			return false
		}

		// Types are equal, compare cards within hands
		for i = 0; i < len(leftHand.Cards); i++ {
			var leftCard = rune(leftHand.Cards[i])
			var rightCard = rune(rightHand.Cards[i])

			if cardRanksWithWild[leftCard] < cardRanksWithWild[rightCard] {
				return true
			} else if cardRanksWithWild[leftCard] > cardRanksWithWild[rightCard] {
				return false
			}

			// card ranks are equal, check the next card
		}

		// if we got this far, then the hands are equal
		return false
	})
}

func CalculateTotalWinnings(hands []PokerHand) int {
	var total = 0
	for i, card := range hands {
		var rank = i + 1
		total += rank * card.Bid
	}

	return total
}
