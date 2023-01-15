package dto

import (
	"fmt"
	"strings"
)

type CardValue struct {
	ValueName   string
	ValueWeight int
}

type CardSuit string

type PlayingCard struct {
	ID        int
	CardValue CardValue
	CardSuit  CardSuit
}

func (playingCard PlayingCard) ToString() string {
	return fmt.Sprintf("{Id: %d, Suit: %s, Value: %s}", playingCard.ID, playingCard.CardSuit, playingCard.CardValue.ValueName)
}

type CardsSequence struct {
	CardsByIds map[int]PlayingCard
	Cards      []PlayingCard
}

func NewCardSequence(cards []PlayingCard) CardsSequence {
	cardsByIds := make(map[int]PlayingCard)
	for _, card := range cards {
		cardsByIds[card.ID] = card
	}
	return CardsSequence{CardsByIds: cardsByIds, Cards: cards}
}

func (cardsSequence CardsSequence) Equals(another CardsSequence) bool {
	if len(cardsSequence.Cards) != len(another.Cards) {
		return false
	}
	for i := range cardsSequence.CardsByIds {
		_, anotherContains := another.CardsByIds[i]
		if !anotherContains {
			return false
		}
	}
	return true
}

func (cardsSequence CardsSequence) ToString() string {
	sb := strings.Builder{}
	sb.WriteString("{ ")
	for _, card := range cardsSequence.Cards {
		sb.WriteString(card.ToString())
		sb.WriteString(" ")
	}
	sb.WriteString("}")
	return sb.String()
}

var SpadeSuit = CardSuit("SPADE")
var ClubSuit = CardSuit("CLUB")
var DiamondSuit = CardSuit("DIAMOND")
var HeartSuit = CardSuit("HEART")

var TwoValue = CardValue{ValueName: "TWO", ValueWeight: 2}
var ThreeValue = CardValue{ValueName: "THREE", ValueWeight: 3}
var FourValue = CardValue{ValueName: "FOUR", ValueWeight: 4}
var FiveValue = CardValue{ValueName: "FIVE", ValueWeight: 5}
var SixValue = CardValue{ValueName: "SIX", ValueWeight: 6}
var SevenValue = CardValue{ValueName: "SEVEN", ValueWeight: 7}
var EightValue = CardValue{ValueName: "EIGHT", ValueWeight: 8}
var NineValue = CardValue{ValueName: "NINE", ValueWeight: 9}
var TenValue = CardValue{ValueName: "TEN", ValueWeight: 10}
var JackValue = CardValue{ValueName: "JACK", ValueWeight: 11}
var QueenValue = CardValue{ValueName: "QUEEN", ValueWeight: 12}
var KingValue = CardValue{ValueName: "KING", ValueWeight: 13}
var AceValue = CardValue{ValueName: "ACE", ValueWeight: 14}

var DiamondTwo = PlayingCard{0, TwoValue, DiamondSuit}
var DiamondThree = PlayingCard{1, ThreeValue, DiamondSuit}
var DiamondFour = PlayingCard{2, FourValue, DiamondSuit}
var DiamondFive = PlayingCard{3, FiveValue, DiamondSuit}
var DiamondSix = PlayingCard{4, SixValue, DiamondSuit}
var DiamondSeven = PlayingCard{5, SevenValue, DiamondSuit}
var DiamondEight = PlayingCard{6, EightValue, DiamondSuit}
var DiamondNine = PlayingCard{7, NineValue, DiamondSuit}
var DiamondTen = PlayingCard{8, TenValue, DiamondSuit}
var DiamondJack = PlayingCard{9, JackValue, DiamondSuit}
var DiamondQueen = PlayingCard{10, QueenValue, DiamondSuit}
var DiamondKing = PlayingCard{11, KingValue, DiamondSuit}
var DiamondAce = PlayingCard{12, AceValue, DiamondSuit}
var HeartTwo = PlayingCard{13, TwoValue, HeartSuit}
var HeartThree = PlayingCard{14, ThreeValue, HeartSuit}
var HeartFour = PlayingCard{15, FourValue, HeartSuit}
var HeartFive = PlayingCard{16, FiveValue, HeartSuit}
var HeartSix = PlayingCard{17, SixValue, HeartSuit}
var HeartSeven = PlayingCard{18, SevenValue, HeartSuit}
var HeartEight = PlayingCard{19, EightValue, HeartSuit}
var HeartNine = PlayingCard{20, NineValue, HeartSuit}
var HeartTen = PlayingCard{21, TenValue, HeartSuit}
var HeartJack = PlayingCard{22, JackValue, HeartSuit}
var HeartQueen = PlayingCard{23, QueenValue, HeartSuit}
var HeartKing = PlayingCard{24, KingValue, HeartSuit}
var HeartAce = PlayingCard{25, AceValue, HeartSuit}
var ClubTwo = PlayingCard{26, TwoValue, ClubSuit}
var ClubThree = PlayingCard{27, ThreeValue, ClubSuit}
var ClubFour = PlayingCard{28, FourValue, ClubSuit}
var ClubFive = PlayingCard{29, FiveValue, ClubSuit}
var ClubSix = PlayingCard{30, SixValue, ClubSuit}
var ClubSeven = PlayingCard{31, SevenValue, ClubSuit}
var ClubEight = PlayingCard{32, EightValue, ClubSuit}
var ClubNine = PlayingCard{33, NineValue, ClubSuit}
var ClubTen = PlayingCard{34, TenValue, ClubSuit}
var ClubJack = PlayingCard{35, JackValue, ClubSuit}
var ClubQueen = PlayingCard{36, QueenValue, ClubSuit}
var ClubKing = PlayingCard{37, KingValue, ClubSuit}
var ClubAce = PlayingCard{38, AceValue, ClubSuit}
var SpadeTwo = PlayingCard{39, TwoValue, SpadeSuit}
var SpadeThree = PlayingCard{40, ThreeValue, SpadeSuit}
var SpadeFour = PlayingCard{41, FourValue, SpadeSuit}
var SpadeFive = PlayingCard{42, FiveValue, SpadeSuit}
var SpadeSix = PlayingCard{43, SixValue, SpadeSuit}
var SpadeSeven = PlayingCard{44, SevenValue, SpadeSuit}
var SpadeEight = PlayingCard{45, EightValue, SpadeSuit}
var SpadeNine = PlayingCard{46, NineValue, SpadeSuit}
var SpadeTen = PlayingCard{47, TenValue, SpadeSuit}
var SpadeJack = PlayingCard{48, JackValue, SpadeSuit}
var SpadeQueen = PlayingCard{49, QueenValue, SpadeSuit}
var SpadeKing = PlayingCard{50, KingValue, SpadeSuit}
var SpadeAce = PlayingCard{51, AceValue, SpadeSuit}

var AvailableCards = []PlayingCard{
	DiamondTwo,
	DiamondThree,
	DiamondFour,
	DiamondFive,
	DiamondSix,
	DiamondSeven,
	DiamondEight,
	DiamondNine,
	DiamondTen,
	DiamondJack,
	DiamondQueen,
	DiamondKing,
	DiamondAce,
	HeartTwo,
	HeartThree,
	HeartFour,
	HeartFive,
	HeartSix,
	HeartSeven,
	HeartEight,
	HeartNine,
	HeartTen,
	HeartJack,
	HeartQueen,
	HeartKing,
	HeartAce,
	ClubTwo,
	ClubThree,
	ClubFour,
	ClubFive,
	ClubSix,
	ClubSeven,
	ClubEight,
	ClubNine,
	ClubTen,
	ClubJack,
	ClubQueen,
	ClubKing,
	ClubAce,
	SpadeTwo,
	SpadeThree,
	SpadeFour,
	SpadeFive,
	SpadeSix,
	SpadeSeven,
	SpadeEight,
	SpadeNine,
	SpadeTen,
	SpadeJack,
	SpadeQueen,
	SpadeKing,
	SpadeAce,
}
