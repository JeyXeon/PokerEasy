package combinations

import (
	"github.com/JeyXeon/poker-easy/dto"
	"testing"
)

func TestRoyalFlash(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "TWO", CardSuit: "HEARTS"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "QUEEN", CardSuit: "CLUB"},
		{CardValue: "KING", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "TEN", CardSuit: "CLUB"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "QUEEN", CardSuit: "CLUB"},
		{CardValue: "KING", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.FlushRoyalName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FlushRoyalName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestStraightFlush(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "TWO", CardSuit: "HEARTS"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "QUEEN", CardSuit: "CLUB"},
		{CardValue: "KING", CardSuit: "CLUB"},
		{CardValue: "NINE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "NINE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "QUEEN", CardSuit: "CLUB"},
		{CardValue: "KING", CardSuit: "CLUB"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.StraightFlushName,
		CombinationValue:    dto.CombinationNamesByValues[dto.StraightFlushName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestFlush(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "TWO", CardSuit: "HEARTS"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "EIGHT", CardSuit: "CLUB"},
		{CardValue: "KING", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "EIGHT", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "KING", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.FlushName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FlushName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestFourOfAKind(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "TWO", CardSuit: "HEARTS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "TWO", CardSuit: "DIAMONDS"},
		{CardValue: "KING", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "TWO", CardSuit: "HEARTS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "TWO", CardSuit: "DIAMONDS"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.FourOfAKindName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FourOfAKindName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestStraight(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "FOUR", CardSuit: "CLUB"},
		{CardValue: "FIVE", CardSuit: "DIAMONDS"},
		{CardValue: "SIX", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "FOUR", CardSuit: "CLUB"},
		{CardValue: "FIVE", CardSuit: "DIAMONDS"},
		{CardValue: "SIX", CardSuit: "CLUB"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.StraightName,
		CombinationValue:    dto.CombinationNamesByValues[dto.StraightName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestFullHouse(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "THREE", CardSuit: "CLUB"},
		{CardValue: "THREE", CardSuit: "DIAMONDS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "THREE", CardSuit: "CLUB"},
		{CardValue: "THREE", CardSuit: "DIAMONDS"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.FullHouseName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FullHouseName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestThreeOfAKind(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "JACK", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "THREE", CardSuit: "CLUB"},
		{CardValue: "THREE", CardSuit: "DIAMONDS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "THREE", CardSuit: "CLUB"},
		{CardValue: "THREE", CardSuit: "DIAMONDS"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.ThreeOfAKindName,
		CombinationValue:    dto.CombinationNamesByValues[dto.ThreeOfAKindName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestTwoPairs(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "JACK", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "THREE", CardSuit: "CLUB"},
		{CardValue: "TWO", CardSuit: "DIAMONDS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "DIAMONDS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "THREE", CardSuit: "CLUB"},
		{CardValue: "THREE", CardSuit: "DIAMONDS"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.TwoPairsName,
		CombinationValue:    dto.CombinationNamesByValues[dto.TwoPairsName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestPair(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "FOUR", CardSuit: "CLUB"},
		{CardValue: "FIVE", CardSuit: "DIAMONDS"},
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "CLUB"},
		{CardValue: "TWO", CardSuit: "SPADE"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.PairName,
		CombinationValue:    dto.CombinationNamesByValues[dto.PairName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}

func TestHighCard(t *testing.T) {
	cards := []dto.PlayingCard{
		{CardValue: "TWO", CardSuit: "SPADE"},
		{CardValue: "THREE", CardSuit: "HEARTS"},
		{CardValue: "JACK", CardSuit: "CLUB"},
		{CardValue: "FIVE", CardSuit: "DIAMONDS"},
		{CardValue: "QUEEN", CardSuit: "CLUB"},
		{CardValue: "ACE", CardSuit: "CLUB"},
		{CardValue: "TEN", CardSuit: "CLUB"},
	}

	cardSequence := []dto.PlayingCard{
		{CardValue: "ACE", CardSuit: "CLUB"},
	}
	expected := dto.CardsCombination{
		CombinationName:     dto.HighCardName,
		CombinationValue:    dto.CombinationNamesByValues[dto.HighCardName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	if expected.CombinationName != got.CombinationName {
		t.Errorf("got %s; expected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("got %d; expected %d", got.CombinationValue, expected.CombinationValue)
	}
}
