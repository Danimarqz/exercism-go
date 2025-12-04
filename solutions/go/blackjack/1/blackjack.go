package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int;
	switch  card{
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}
	return value;
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	valueDealer := ParseCard(dealerCard)
	if (value1 == 11 && value1 == value2) {
		return "P"
	}
	if (value1 + value2 == 21) {
		if (valueDealer < 10) {
			return "W"
		} else {
			return "S"
		}
	}
	if (value1 + value2 >= 17) {
		return "S"
	}
	if (value1 + value2 >= 12 && valueDealer < 7) {
		return "S"
	}
	return "H"
}
