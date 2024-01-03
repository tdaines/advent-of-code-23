package day07_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day07"
)

func TestParsePokerHandString(t *testing.T) {
	var cards, bid = day07.ParsePokerHandString("32T3K 765")
	assert.Equal(t, "32T3K", cards)
	assert.Equal(t, 765, bid)
}

func TestNewPokerHand(t *testing.T) {
	var hand = day07.NewPokerHand("32T3K", 765)
	assert.Equal(t, "32T3K", hand.Cards)
	assert.Equal(t, day07.OnePair, hand.Type)
	assert.Equal(t, 765, hand.Bid)

	hand = day07.NewPokerHand("T55J5", 684)
	assert.Equal(t, "T55J5", hand.Cards)
	assert.Equal(t, day07.ThreeOfAKind, hand.Type)
	assert.Equal(t, 684, hand.Bid)

	hand = day07.NewPokerHand("KK677", 28)
	assert.Equal(t, "KK677", hand.Cards)
	assert.Equal(t, day07.TwoPair, hand.Type)
	assert.Equal(t, 28, hand.Bid)

	hand = day07.NewPokerHand("KTJJT", 220)
	assert.Equal(t, "KTJJT", hand.Cards)
	assert.Equal(t, day07.TwoPair, hand.Type)
	assert.Equal(t, 220, hand.Bid)

	hand = day07.NewPokerHand("QQQJA", 483)
	assert.Equal(t, "QQQJA", hand.Cards)
	assert.Equal(t, day07.ThreeOfAKind, hand.Type)
	assert.Equal(t, 483, hand.Bid)
}

func TestSortPokerHands(t *testing.T) {
	var hands = []day07.PokerHand{
		day07.NewPokerHand("32T3K", 765),
		day07.NewPokerHand("T55J5", 684),
		day07.NewPokerHand("KK677", 28),
		day07.NewPokerHand("KTJJT", 220),
		day07.NewPokerHand("QQQJA", 483),
	}

	day07.SortPokerHands(hands)
	assert.Equal(t, "32T3K", hands[0].Cards)
	assert.Equal(t, "KTJJT", hands[1].Cards)
	assert.Equal(t, "KK677", hands[2].Cards)
	assert.Equal(t, "T55J5", hands[3].Cards)
	assert.Equal(t, "QQQJA", hands[4].Cards)
}

func TestCalculateTotalWinnings(t *testing.T) {
	var hands = []day07.PokerHand{
		day07.NewPokerHand("32T3K", 765),
		day07.NewPokerHand("KTJJT", 220),
		day07.NewPokerHand("KK677", 28),
		day07.NewPokerHand("T55J5", 684),
		day07.NewPokerHand("QQQJA", 483),
	}

	assert.Equal(t, 6440, day07.CalculateTotalWinnings(hands))
}

func TestNewPokerHandWithWilds(t *testing.T) {
	var hand = day07.NewPokerHandWithWilds("32T3K", 765)
	assert.Equal(t, "32T3K", hand.Cards)
	assert.Equal(t, day07.OnePair, hand.Type)
	assert.Equal(t, 765, hand.Bid)

	hand = day07.NewPokerHandWithWilds("T55J5", 684)
	assert.Equal(t, "T55J5", hand.Cards)
	assert.Equal(t, day07.FourOfAKind, hand.Type)
	assert.Equal(t, 684, hand.Bid)

	hand = day07.NewPokerHandWithWilds("KK677", 28)
	assert.Equal(t, "KK677", hand.Cards)
	assert.Equal(t, day07.TwoPair, hand.Type)
	assert.Equal(t, 28, hand.Bid)

	hand = day07.NewPokerHandWithWilds("KTJJT", 220)
	assert.Equal(t, "KTJJT", hand.Cards)
	assert.Equal(t, day07.FourOfAKind, hand.Type)
	assert.Equal(t, 220, hand.Bid)

	hand = day07.NewPokerHandWithWilds("QQQJA", 483)
	assert.Equal(t, "QQQJA", hand.Cards)
	assert.Equal(t, day07.FourOfAKind, hand.Type)
	assert.Equal(t, 483, hand.Bid)
}

func TestSortPokerHandsWithWilds(t *testing.T) {
	var hands = []day07.PokerHand{
		day07.NewPokerHandWithWilds("32T3K", 765),
		day07.NewPokerHandWithWilds("T55J5", 684),
		day07.NewPokerHandWithWilds("KK677", 28),
		day07.NewPokerHandWithWilds("KTJJT", 220),
		day07.NewPokerHandWithWilds("QQQJA", 483),
	}

	day07.SortPokerHandsWithWilds(hands)
	assert.Equal(t, "32T3K", hands[0].Cards)
	assert.Equal(t, "KK677", hands[1].Cards)
	assert.Equal(t, "T55J5", hands[2].Cards)
	assert.Equal(t, "QQQJA", hands[3].Cards)
	assert.Equal(t, "KTJJT", hands[4].Cards)
}

func TestCalculateTotalWinnings_WithWilds(t *testing.T) {
	var hands = []day07.PokerHand{
		day07.NewPokerHandWithWilds("32T3K", 765),
		day07.NewPokerHandWithWilds("KTJJT", 220),
		day07.NewPokerHandWithWilds("KK677", 28),
		day07.NewPokerHandWithWilds("T55J5", 684),
		day07.NewPokerHandWithWilds("QQQJA", 483),
	}

	day07.SortPokerHandsWithWilds(hands)
	assert.Equal(t, 5905, day07.CalculateTotalWinnings(hands))
}
