package dto

import (
	"fmt"
	"strings"
)

type CardValue string
type CardSuit string

type PlayingCard struct {
	ID        uint
	CardValue CardValue
	CardSuit  CardSuit
}

func (playingCard PlayingCard) ToString() string {
	return fmt.Sprintf("{Id: %d, Suit: %s, Value: %s}", playingCard.ID, playingCard.CardSuit, playingCard.CardValue)
}

type CardsSequence struct {
	CardsByIds map[uint]PlayingCard
	Cards      []PlayingCard
}

func NewCardSequence(cards []PlayingCard) CardsSequence {
	cardsByIds := make(map[uint]PlayingCard)
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
	sb.WriteString("Sequence { ")
	for _, card := range cardsSequence.Cards {
		sb.WriteString(card.ToString())
		sb.WriteString(" ")
	}
	sb.WriteString("}")
	return sb.String()
}

var CardValuesByWeight = map[CardValue]int{
	"TWO":   2,
	"THREE": 3,
	"FOUR":  4,
	"FIVE":  5,
	"SIX":   6,
	"SEVEN": 7,
	"EIGHT": 8,
	"NINE":  9,
	"TEN":   10,
	"JACK":  11,
	"QUEEN": 12,
	"KING":  13,
	"ACE":   14,
}

var DiamondTwo = PlayingCard{0, "TWO", "DIAMOND"}
var DiamondThree = PlayingCard{1, "THREE", "DIAMOND"}
var DiamondFour = PlayingCard{2, "FOUR", "DIAMOND"}
var DiamondFive = PlayingCard{3, "FIVE", "DIAMOND"}
var DiamondSix = PlayingCard{4, "SIX", "DIAMOND"}
var DiamondSeven = PlayingCard{5, "SEVEN", "DIAMOND"}
var DiamondEight = PlayingCard{6, "EIGHT", "DIAMOND"}
var DiamondNine = PlayingCard{7, "NINE", "DIAMOND"}
var DiamondTen = PlayingCard{8, "TEN", "DIAMOND"}
var DiamondJack = PlayingCard{9, "JACK", "DIAMOND"}
var DiamondQueen = PlayingCard{10, "QUEEN", "DIAMOND"}
var DiamondKing = PlayingCard{11, "KING", "DIAMOND"}
var DiamondAce = PlayingCard{12, "ACE", "DIAMOND"}
var HeartTwo = PlayingCard{13, "TWO", "HEART"}
var HeartThree = PlayingCard{14, "THREE", "HEART"}
var HeartFour = PlayingCard{15, "FOUR", "HEART"}
var HeartFive = PlayingCard{16, "FIVE", "HEART"}
var HeartSix = PlayingCard{17, "SIX", "HEART"}
var HeartSeven = PlayingCard{18, "SEVEN", "HEART"}
var HeartEight = PlayingCard{19, "EIGHT", "HEART"}
var HeartNine = PlayingCard{20, "NINE", "HEART"}
var HeartTen = PlayingCard{21, "TEN", "HEART"}
var HeartJack = PlayingCard{22, "JACK", "HEART"}
var HeartQueen = PlayingCard{23, "QUEEN", "HEART"}
var HeartKing = PlayingCard{24, "KING", "HEART"}
var HeartAce = PlayingCard{25, "ACE", "HEART"}
var ClubTwo = PlayingCard{26, "TWO", "CLUB"}
var ClubThree = PlayingCard{27, "THREE", "CLUB"}
var ClubFour = PlayingCard{28, "FOUR", "CLUB"}
var ClubFive = PlayingCard{29, "FIVE", "CLUB"}
var ClubSix = PlayingCard{30, "SIX", "CLUB"}
var ClubSeven = PlayingCard{31, "SEVEN", "CLUB"}
var ClubEight = PlayingCard{32, "EIGHT", "CLUB"}
var ClubNine = PlayingCard{33, "NINE", "CLUB"}
var ClubTen = PlayingCard{34, "TEN", "CLUB"}
var ClubJack = PlayingCard{35, "JACK", "CLUB"}
var ClubQueen = PlayingCard{36, "QUEEN", "CLUB"}
var ClubKing = PlayingCard{37, "KING", "CLUB"}
var ClubAce = PlayingCard{38, "ACE", "CLUB"}
var SpadeTwo = PlayingCard{39, "TWO", "SPADE"}
var SpadeThree = PlayingCard{40, "THREE", "SPADE"}
var SpadeFour = PlayingCard{41, "FOUR", "SPADE"}
var SpadeFive = PlayingCard{42, "FIVE", "SPADE"}
var SpadeSix = PlayingCard{43, "SIX", "SPADE"}
var SpadeSeven = PlayingCard{44, "SEVEN", "SPADE"}
var SpadeEight = PlayingCard{45, "EIGHT", "SPADE"}
var SpadeNine = PlayingCard{46, "NINE", "SPADE"}
var SpadeTen = PlayingCard{47, "TEN", "SPADE"}
var SpadeJack = PlayingCard{48, "JACK", "SPADE"}
var SpadeQueen = PlayingCard{49, "QUEEN", "SPADE"}
var SpadeKing = PlayingCard{50, "KING", "SPADE"}
var SpadeAce = PlayingCard{51, "ACE", "SPADE"}

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
