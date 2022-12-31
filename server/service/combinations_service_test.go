package service

import (
	"github.com/JeyXeon/poker-easy/dto"
	"testing"
)

func TestRoyalFlash(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartTwo,
		dto.ClubJack,
		dto.ClubQueen,
		dto.ClubKing,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.ClubTen,
		dto.ClubJack,
		dto.ClubQueen,
		dto.ClubKing,
		dto.ClubAce,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.FlushRoyalName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FlushRoyalName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestStraightFlush(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartTwo,
		dto.ClubJack,
		dto.ClubQueen,
		dto.ClubKing,
		dto.ClubNine,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.ClubNine,
		dto.ClubTen,
		dto.ClubJack,
		dto.ClubQueen,
		dto.ClubKing,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.StraightFlushName,
		CombinationValue:    dto.CombinationNamesByValues[dto.StraightFlushName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestFlush(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartTwo,
		dto.ClubJack,
		dto.ClubEight,
		dto.ClubKing,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.ClubEight,
		dto.ClubTen,
		dto.ClubJack,
		dto.ClubKing,
		dto.ClubAce,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.FlushName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FlushName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestFourOfAKind(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartTwo,
		dto.ClubTwo,
		dto.DiamondTwo,
		dto.ClubKing,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartTwo,
		dto.ClubTwo,
		dto.DiamondTwo,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.FourOfAKindName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FourOfAKindName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestStraight(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartThree,
		dto.ClubFour,
		dto.DiamondFive,
		dto.ClubSix,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartThree,
		dto.ClubFour,
		dto.DiamondFive,
		dto.ClubSix,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.StraightName,
		CombinationValue:    dto.CombinationNamesByValues[dto.StraightName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestFullHouse(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartThree,
		dto.ClubThree,
		dto.DiamondThree,
		dto.ClubTwo,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.ClubTwo,
		dto.SpadeTwo,
		dto.HeartThree,
		dto.ClubThree,
		dto.DiamondThree,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.FullHouseName,
		CombinationValue:    dto.CombinationNamesByValues[dto.FullHouseName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestThreeOfAKind(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeJack,
		dto.HeartThree,
		dto.ClubThree,
		dto.DiamondThree,
		dto.ClubTwo,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.HeartThree,
		dto.ClubThree,
		dto.DiamondThree,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.ThreeOfAKindName,
		CombinationValue:    dto.CombinationNamesByValues[dto.ThreeOfAKindName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestTwoPairs(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeJack,
		dto.HeartThree,
		dto.ClubThree,
		dto.DiamondTwo,
		dto.ClubTwo,
		dto.ClubAce,
		dto.HeartQueen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.HeartThree,
		dto.ClubThree,
		dto.DiamondTwo,
		dto.ClubTwo,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.TwoPairsName,
		CombinationValue:    dto.CombinationNamesByValues[dto.TwoPairsName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestPair(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartThree,
		dto.ClubFour,
		dto.DiamondFive,
		dto.ClubTwo,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.ClubTwo,
		dto.SpadeTwo,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.PairName,
		CombinationValue:    dto.CombinationNamesByValues[dto.PairName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func TestHighCard(t *testing.T) {
	cards := dto.NewCardSequence([]dto.PlayingCard{
		dto.SpadeTwo,
		dto.HeartThree,
		dto.ClubJack,
		dto.DiamondFive,
		dto.ClubQueen,
		dto.ClubAce,
		dto.ClubTen,
	})

	cardSequence := dto.NewCardSequence([]dto.PlayingCard{
		dto.ClubAce,
	})
	expected := dto.CardsCombination{
		CombinationName:     dto.HighCardName,
		CombinationValue:    dto.CombinationNamesByValues[dto.HighCardName],
		CombinationSequence: cardSequence,
	}
	got := CalculateBestCombination(cards)

	assertCombinations(t, got, expected)
}

func assertCombinations(t *testing.T, got dto.CardsCombination, expected dto.CardsCombination) {
	if expected.CombinationName != got.CombinationName {
		t.Errorf("\ngot %s;\nexpected %s", got.CombinationName, expected.CombinationName)
	}
	if expected.CombinationValue != got.CombinationValue {
		t.Errorf("\ngot %d;\nexpected %d", got.CombinationValue, expected.CombinationValue)
	}

	if !expected.CombinationSequence.Equals(got.CombinationSequence) {
		t.Errorf("\ngot %s;\nexpected %s", got.CombinationSequence.ToString(), expected.CombinationSequence.ToString())
	}
}
