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
	CombinationSequence []PlayingCard
}

type ExistingCombinations map[string][]PlayingCard

func (existingCombinations ExistingCombinations) AddAbstractFlush(flushSequence []PlayingCard) {
	hasStraight, maxStraightSequence := hasStraight(flushSequence)
	if hasStraight {
		existingCombinations.AddStraightFlush(maxStraightSequence)
	} else {
		existingCombinations.AddFlush(flushSequence)
	}
}

func (existingCombinations ExistingCombinations) AddStraightFlush(maxStraightSequence []PlayingCard) {
	if maxStraightSequence[len(maxStraightSequence)-1].CardValue == "ACE" {
		existingCombinations[FlushRoyalName] = maxStraightSequence
	} else {
		existingCombinations[StraightFlushName] = maxStraightSequence
	}
}

func (existingCombinations ExistingCombinations) AddFlush(cardsSequence []PlayingCard) {
	firstSequenceCardIdx := len(cardsSequence) - 5
	lastSequenceCardIdx := len(cardsSequence) - 1
	maxFlushSequence := cardsSequence[firstSequenceCardIdx:lastSequenceCardIdx]
	existingCombinations[FlushName] = maxFlushSequence
}

func (existingCombinations ExistingCombinations) AddFourOfAKind(playingCards []PlayingCard) {
	existingCombinations[FourOfAKindName] = playingCards
}

func (existingCombinations ExistingCombinations) AddFullHouse() {
	threeSequence := existingCombinations[PairName]
	pairSequence := existingCombinations[ThreeOfAKindName]
	fullHouseSequence := append(pairSequence, threeSequence...)
	existingCombinations[FullHouseName] = fullHouseSequence

}

func (existingCombinations ExistingCombinations) AddStraight(cards []PlayingCard) {
	hasStraight, sequence := hasStraight(cards)
	if hasStraight {
		existingCombinations[StraightName] = sequence
	}
}

func (existingCombinations ExistingCombinations) AddPair(playingCards []PlayingCard, cardValue CardValue) {
	existingPair, hasPair := existingCombinations[PairName]
	if !hasPair || (hasPair && existingPair[0].CardValue < cardValue) {
		existingCombinations[PairName] = playingCards
	}
	if hasPair {
		twoPairsSequence := append(existingPair, playingCards...)
		existingCombinations[TwoPairsName] = twoPairsSequence
	}
}

func (existingCombinations ExistingCombinations) AddThreeOfAKind(playingCards []PlayingCard, cardValue CardValue) {
	_, hasPair := existingCombinations[PairName]
	if !hasPair {
		existingCombinations[PairName] = playingCards[0:1]
	}

	existingThreeOfAKind, hasThreeOfAKind := existingCombinations[ThreeOfAKindName]
	if !hasThreeOfAKind || (hasThreeOfAKind && existingThreeOfAKind[0].CardValue < cardValue) {
		existingCombinations[ThreeOfAKindName] = playingCards
	}
}

func (existingCombinations ExistingCombinations) AddHighCard(highestCard PlayingCard) {
	if len(existingCombinations) == 0 {
		existingCombinations[HighCardName] = []PlayingCard{highestCard}
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
	return hasPair && hasThreeOfAKind && threeSequence[0].CardValue != pairSequence[0].CardValue
}

func hasStraight(sortedCards []PlayingCard) (bool, []PlayingCard) {
	counter := 5
	sequence := make([]PlayingCard, 5)
	lastValue := -1
	for i := len(sortedCards) - 1; i >= 0; i-- {
		card := sortedCards[i]
		cardValue := card.CardValue
		if lastValue-CardValuesByWeight[cardValue] == 1 {
			sequence[counter-1] = card
			lastValue = CardValuesByWeight[cardValue]
			counter--
		} else {
			sequence = make([]PlayingCard, 5)
			counter = 4
			sequence[counter] = card
			lastValue = CardValuesByWeight[cardValue]
		}

		if counter == 0 {
			return true, sequence
		}
	}

	return false, nil
}
