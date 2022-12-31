package service

import (
	"github.com/JeyXeon/poker-easy/dto"
	"sort"
)

func CalculateBestCombination(cardSequence dto.CardsSequence) dto.CardsCombination {
	sort.Slice(cardSequence.Cards, func(i, j int) bool {
		return dto.CardValuesByWeight[cardSequence.Cards[i].CardValue] < dto.CardValuesByWeight[cardSequence.Cards[j].CardValue]
	})

	existingCombinations := make(dto.ExistingCombinations)

	cardValues, cardsByValues, cardsBySuites, highestCard := fillData(cardSequence.Cards)

	for _, suitCardSequence := range cardsBySuites {
		if len(suitCardSequence.Cards) >= 5 {
			existingCombinations.AddAbstractFlush(suitCardSequence)
		}
	}

	for _, cardValue := range cardValues {
		valueCardSequence := cardsByValues[cardValue]
		if len(valueCardSequence.Cards) == 4 {
			existingCombinations.AddFourOfAKind(valueCardSequence)
		}

		if len(valueCardSequence.Cards) == 2 {
			existingCombinations.AddPair(valueCardSequence, cardValue)
		}

		if len(valueCardSequence.Cards) == 3 {
			existingCombinations.AddThreeOfAKind(valueCardSequence, cardValue)
		}
	}

	if !existingCombinations.HasStraightFlush() {
		existingCombinations.AddStraight(cardSequence)
	}

	if existingCombinations.HasFullHouse() {
		existingCombinations.AddFullHouse()
	}

	existingCombinations.AddHighCard(highestCard)

	bestCombination := chooseBestCombination(existingCombinations)
	return bestCombination
}

func fillData(cards []dto.PlayingCard) ([]dto.CardValue, map[dto.CardValue]dto.CardsSequence, map[dto.CardSuit]dto.CardsSequence, dto.PlayingCard) {
	cardValues := make([]dto.CardValue, 0, 7)
	cardsByValues := make(map[dto.CardValue][]dto.PlayingCard)
	cardsBySuites := make(map[dto.CardSuit][]dto.PlayingCard)
	var highestCard dto.PlayingCard
	higherValue := -1

	for _, card := range cards {
		if dto.CardValuesByWeight[card.CardValue] > higherValue {
			highestCard = card
			higherValue = dto.CardValuesByWeight[card.CardValue]
		}

		cardSuit := card.CardSuit
		cardValue := card.CardValue

		_, containsValue := cardsByValues[cardValue]
		_, containsSuit := cardsBySuites[cardSuit]

		if !containsValue {
			cardValues = append(cardValues, cardValue)

			cardsSlice := make([]dto.PlayingCard, 0, 4)
			cardsByValues[cardValue] = cardsSlice
		}

		if !containsSuit {
			cardsSlice := make([]dto.PlayingCard, 0, 7)
			cardsBySuites[cardSuit] = cardsSlice
		}

		cardsByValues[cardValue] = append(cardsByValues[cardValue], card)
		cardsBySuites[cardSuit] = append(cardsBySuites[cardSuit], card)
	}

	sequenceByValue := make(map[dto.CardValue]dto.CardsSequence)
	sequenceBySuit := make(map[dto.CardSuit]dto.CardsSequence)

	for value, cards := range cardsByValues {
		sequenceByValue[value] = dto.NewCardSequence(cards)
	}

	for suit, cards := range cardsBySuites {
		sequenceBySuit[suit] = dto.NewCardSequence(cards)
	}

	return cardValues, sequenceByValue, sequenceBySuit, highestCard
}

func chooseBestCombination(existingCombinations map[string]dto.CardsSequence) dto.CardsCombination {
	var bestCombination dto.CardsCombination
	bestCombinationValue := -1

	for combinationName, cardsSequence := range existingCombinations {
		if dto.CombinationNamesByValues[combinationName] > bestCombinationValue {
			bestCombinationValue = dto.CombinationNamesByValues[combinationName]
			bestCombination = dto.CardsCombination{
				CombinationName:     combinationName,
				CombinationValue:    bestCombinationValue,
				CombinationSequence: cardsSequence,
			}
		}
	}

	return bestCombination
}
