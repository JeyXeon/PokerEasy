package dto

const (
	FlushRoyalName    = "FLUSH_ROYAL"
	StraightFlushName = "STRAIGHT_FLUSH"
	FourOfAKindName   = "FOUR_OF_A_KIND"
	FullHouseName     = "FULL_HOUSE"
	FlushName         = "FLUSH"
	StraightName      = "STRAIGHT"
	ThreeOfAKindName  = "THREE_OF_A_KIND"
	TwoPairsName      = "TWO_PAIRS"
	PairName          = "PAIR"
	HighCardName      = "HIGH_CARD"
)

var CombinationNamesByValues = map[string]int{
	FlushRoyalName:    9,
	StraightFlushName: 8,
	FourOfAKindName:   7,
	FullHouseName:     6,
	FlushName:         5,
	StraightName:      4,
	ThreeOfAKindName:  3,
	TwoPairsName:      2,
	PairName:          1,
	HighCardName:      0,
}

type CardsCombination struct {
	CombinationName     string
	CombinationValue    int
	CombinationSequence CardsSequence
}

type ExistingCombinations map[string]CardsSequence

func (existingCombinations ExistingCombinations) AddAbstractFlush(flushSequence CardsSequence) {
	hasStraight, maxStraightSequence := hasStraight(flushSequence.Cards)
	if hasStraight {
		existingCombinations.AddStraightFlush(maxStraightSequence)
	} else {
		existingCombinations.AddFlush(flushSequence)
	}
}

func (existingCombinations ExistingCombinations) AddStraightFlush(maxStraightSequence CardsSequence) {
	if maxStraightSequence.Cards[len(maxStraightSequence.Cards)-1].CardValue.ValueName == "ACE" {
		existingCombinations[FlushRoyalName] = maxStraightSequence
	} else {
		existingCombinations[StraightFlushName] = maxStraightSequence
	}
}

func (existingCombinations ExistingCombinations) AddFlush(cardsSequence CardsSequence) {
	firstSequenceCardIdx := len(cardsSequence.Cards) - 5
	lastSequenceCardIdx := len(cardsSequence.Cards)
	maxFlushSequence := cardsSequence.Cards[firstSequenceCardIdx:lastSequenceCardIdx]
	existingCombinations[FlushName] = NewCardSequence(maxFlushSequence)
}

func (existingCombinations ExistingCombinations) AddFourOfAKind(cardsSequence CardsSequence) {
	existingCombinations[FourOfAKindName] = cardsSequence
}

func (existingCombinations ExistingCombinations) AddFullHouse() {
	threeSequence := existingCombinations[PairName]
	pairSequence := existingCombinations[ThreeOfAKindName]
	fullHouseSequence := append(pairSequence.Cards, threeSequence.Cards...)
	existingCombinations[FullHouseName] = NewCardSequence(fullHouseSequence)

}

func (existingCombinations ExistingCombinations) AddStraight(cardsSequence CardsSequence) {
	hasStraight, sequence := hasStraight(cardsSequence.Cards)
	if hasStraight {
		existingCombinations[StraightName] = sequence
	}
}

func (existingCombinations ExistingCombinations) AddPair(cardsSequence CardsSequence, cardValue CardValue) {
	existingPair, hasPair := existingCombinations[PairName]
	if !hasPair || (hasPair && existingPair.Cards[0].CardValue.ValueWeight < cardValue.ValueWeight) {
		existingCombinations[PairName] = cardsSequence
	}
	if hasPair {
		twoPairsSequence := append(existingPair.Cards, cardsSequence.Cards...)
		existingCombinations[TwoPairsName] = NewCardSequence(twoPairsSequence)
	}
}

func (existingCombinations ExistingCombinations) AddThreeOfAKind(cardsSequence CardsSequence, cardValue CardValue) {
	_, hasPair := existingCombinations[PairName]
	if !hasPair {
		existingCombinations[PairName] = NewCardSequence(cardsSequence.Cards[0:1])
	}

	existingThreeOfAKind, hasThreeOfAKind := existingCombinations[ThreeOfAKindName]
	if !hasThreeOfAKind || (hasThreeOfAKind && existingThreeOfAKind.Cards[0].CardValue.ValueWeight < cardValue.ValueWeight) {
		existingCombinations[ThreeOfAKindName] = cardsSequence
	}
}

func (existingCombinations ExistingCombinations) AddHighCard(highestCard PlayingCard) {
	if len(existingCombinations) == 0 {
		existingCombinations[HighCardName] = NewCardSequence([]PlayingCard{highestCard})
	}
}

func (existingCombinations ExistingCombinations) HasStraightFlush() bool {
	_, hasFlushRoyal := existingCombinations[FlushRoyalName]
	_, hasStraightFlush := existingCombinations[StraightFlushName]
	return hasFlushRoyal || hasStraightFlush
}

func (existingCombinations ExistingCombinations) HasFullHouse() bool {
	threeSequence, hasPair := existingCombinations[PairName]
	pairSequence, hasThreeOfAKind := existingCombinations[ThreeOfAKindName]
	return hasPair && hasThreeOfAKind && threeSequence.Cards[0].CardValue != pairSequence.Cards[0].CardValue
}

func hasStraight(sortedCards []PlayingCard) (bool, CardsSequence) {
	counter := 5
	sequence := make([]PlayingCard, 5)
	lastValue := -1
	for i := len(sortedCards) - 1; i >= 0; i-- {
		card := sortedCards[i]
		cardValue := card.CardValue
		if lastValue-cardValue.ValueWeight == 1 {
			sequence[counter-1] = card
			lastValue = cardValue.ValueWeight
			counter--
		} else {
			sequence = make([]PlayingCard, 5)
			counter = 4
			sequence[counter] = card
			lastValue = cardValue.ValueWeight
		}

		if counter == 0 {
			return true, NewCardSequence(sequence)
		}
	}

	return false, CardsSequence{}
}
