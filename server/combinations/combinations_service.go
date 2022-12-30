package combinations

import (
	"github.com/JeyXeon/poker-easy/dto"
	"sort"
)

func CalculateBestCombination(cards []dto.PlayingCard) dto.CardsCombination {
	sort.Slice(cards, func(i, j int) bool {
		return dto.CardValuesByWeight[cards[i].CardValue] < dto.CardValuesByWeight[cards[j].CardValue]
	})

	existingCombinations := make(dto.ExistingCombinations)

	cardValues, cardsByValues, cardsBySuites, highestCard := fillData(cards)

	for _, playingCards := range cardsBySuites {
		if len(playingCards) >= 5 {
			existingCombinations.AddAbstractFlush(playingCards)
		}
	}

	for _, cardValue := range cardValues {
		playingCards := cardsByValues[cardValue]
		if len(playingCards) == 4 {
			existingCombinations.AddFourOfAKind(playingCards)
		}

		if len(playingCards) == 2 {
			existingCombinations.AddPair(playingCards, cardValue)
		}

		if len(playingCards) == 3 {
			existingCombinations.AddThreeOfAKind(playingCards, cardValue)
		}
	}

	if !existingCombinations.HasStraightFlush() {
		existingCombinations.AddStraight(cards)
	}

	if existingCombinations.HasFullHouse() {
		existingCombinations.AddFullHouse()
	}

	existingCombinations.AddHighCard(highestCard)

	bestCombination := chooseBestCombination(existingCombinations)
	return bestCombination
}

func fillData(cards []dto.PlayingCard) ([]dto.CardValue, map[dto.CardValue][]dto.PlayingCard, map[dto.CardSuit][]dto.PlayingCard, dto.PlayingCard) {
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

	return cardValues, cardsByValues, cardsBySuites, highestCard
}

func chooseBestCombination(existingCombinations map[string][]dto.PlayingCard) dto.CardsCombination {
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
