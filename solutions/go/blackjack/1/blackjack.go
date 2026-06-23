package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11

	case "king", "queen", "jack", "ten":
		return 10

	case "nine":
		return 9

	case "eight":
		return 8

	case "seven":
		return 7

	case "six":
		return 6

	case "five":
		return 5

	case "four":
		return 4

	case "three":
		return 3

	case "two":
		return 2

	case "one":
		return 1

	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	dealerCardVal := ParseCard(dealerCard)

	totalCardsVal := card1Val + card2Val
	isBlackJack := card1Val+card2Val == 21

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"

	case isBlackJack && (dealerCard != "ace" && dealerCard != "king" && dealerCard != "queen" && dealerCard != "jack" && dealerCard != "ten"):
		return "W"

	case isBlackJack:
		return "S"

	case totalCardsVal >= 17 && totalCardsVal <= 20:
		return "S"

	case (totalCardsVal >= 12 && totalCardsVal <= 16) && dealerCardVal < 7:
		return "S"

	case (totalCardsVal >= 12 && totalCardsVal <= 16) && dealerCardVal >= 7:
		return "H"

	default:
		return "H"
	}
}
