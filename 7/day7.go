package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	Hands      []Hand
	ScoreSheet ScoreSheet
	FinalScore []Hand
}

type Hand struct {
	Cards     string
	Bid       int
	Score     int
	Type      int
	TypeScore int
	Rank      int
}

type ScoreSheet struct {
	Fives     []Hand
	Fours     []Hand
	FullHouse []Hand
	Threes    []Hand
	TwoPair   []Hand
	Pair      []Hand
	High      []Hand
}

type Score struct {
	two   int
	three int
	four  int
	five  int
	six   int
	seven int
	eight int
	nine  int
	ten   int
	jack  int
	queen int
	king  int
	ace   int
}

const (
	T         = 10
	J         = 11
	Q         = 12
	K         = 13
	A         = 14
	FIVEKIND  = 7
	FOURKIND  = 6
	FULLHOUSE = 5
	THREEKIND = 4
	TWOPAIR   = 3
	PAIR      = 2
	HIGH      = 1
)

func main() {
	data := new(Data)
	lines := fileToLines("testinput2.txt")
	linesToHands(lines, data)
	fmt.Printf("\n%+v\n", data)
	handsToScores(data)
	fmt.Printf("\nHands: %+v\n", data.Hands)
	arrangeIntoSheets(data)
	fmt.Printf("\nSheets: %+v\n", data.ScoreSheet)
}

func fileToLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return
}

func linesToHands(lines []string, data *Data) {
	for _, line := range lines {
		hand := new(Hand)
		hand.Cards = line[:5]
		hand.Bid, _ = strconv.Atoi(line[6:])
		data.Hands = append(data.Hands, *hand)
	}
}

func handsToScores(data *Data) {
	for i, hand := range data.Hands {
		score := 0
		cards := strings.Split(hand.Cards, "")
		for _, card := range cards {
			score += cardToScore(card)
		}
		data.Hands[i].Score = score
		identifyHandPower(&hand)
		data.Hands[i].Type = hand.Type
		// data.Hands[i].TypeScore = hand.TypeScore
	}
}

func arrangeIntoSheets(data *Data) {
	for _, hand := range data.Hands {
		switch hand.Type {
		case FIVEKIND:
			data.ScoreSheet.Fives = append(data.ScoreSheet.Fives, hand)
		case FOURKIND:
			data.ScoreSheet.Fours = append(data.ScoreSheet.Fours, hand)
		case FULLHOUSE:
			data.ScoreSheet.FullHouse = append(data.ScoreSheet.FullHouse, hand)
		case THREEKIND:
			data.ScoreSheet.Threes = append(data.ScoreSheet.Threes, hand)
		case TWOPAIR:
			data.ScoreSheet.TwoPair = append(data.ScoreSheet.TwoPair, hand)
		case PAIR:
			data.ScoreSheet.Pair = append(data.ScoreSheet.Pair, hand)
		case HIGH:
			data.ScoreSheet.High = append(data.ScoreSheet.High, hand)
		default:
			log.Fatalf("Unprocessable Hand type of %v", hand.Type)
		}
	}
}

func sortSheets(data *Data) {

}

func cardToScore(card string) int {
	if isNumeric(card) {
		cardScore, _ := strconv.Atoi(card)
		return cardScore
	} else {
		return getLetterScore(card)
	}
}

var numeric = regexp.MustCompile("^[0-9]*$")

func isNumeric(str string) bool {
	return numeric.MatchString(str)
}

func getLetterScore(card string) int {
	switch card {
	case "T":
		return T
	case "J":
		return J
	case "Q":
		return Q
	case "K":
		return K
	case "A":
		return A
	default:
		log.Fatalf("unable to find score for card %v", card)
		return 0
	}
}

func identifyHandPower(hand *Hand) {
	// TODO tidy this into a switch
	fmt.Printf("\nChecking hand of cards %v\n", hand.Cards)
	// check for 5 of a kind
	if isFiveOfKind(hand) {
		fmt.Print("FIVE of KIND")
		hand.Type = FIVEKIND
		return
	}

	if isFourOfKind(hand) {
		fmt.Print("FOUR of KIND")
		hand.Type = FOURKIND
		return
	}

	if isFullHouse(hand) {
		fmt.Print("FULL HOUSE")
		hand.Type = FULLHOUSE
		return
	}

	if isThreeOfKind(hand) {
		fmt.Print("THREE of KIND")
		hand.Type = THREEKIND
		return
	}

	if hasTwoPair(hand) {
		fmt.Print("TWO PAIR")
		hand.Type = TWOPAIR
		return
	}

	if hasSinglePair(hand) {
		fmt.Print("PAIR")
		hand.Type = PAIR
		return
	}
	fmt.Print("SINGLE")
	fmt.Print("\n")
	hand.Type = HIGH
}

func isFiveOfKind(h *Hand) bool {
	hand := h.Cards
	firstCard := string(hand[0])
	// fmt.Printf("checking for %v\n", firstCard)
	if strings.Count(hand, firstCard) == 5 {
		h.TypeScore = cardToScore(firstCard)
		return true
	}
	return false
}

func isFourOfKind(h *Hand) bool {
	hand := h.Cards
	firstCard := string(hand[0])
	secondCard := string(hand[1])
	// fmt.Printf("checking for %v against %v found %v\n", firstCard, hand, strings.Count(hand, firstCard))
	// fmt.Printf("checking for %v against %v found %v\n", secondCard, hand, strings.Count(hand, secondCard))
	if strings.Count(hand, firstCard) == 4 {
		h.TypeScore = cardToScore(firstCard)
		return true
	}
	if strings.Count(hand, secondCard) == 4 {
		h.TypeScore = cardToScore(secondCard)
		return true
	}
	return false
}

func isFullHouse(h *Hand) bool {
	if isThreeOfKind(h) && hasSinglePair(h) {
		return true
	}
	return false
}

func isThreeOfKind(h *Hand) bool {
	hand := h.Cards
	firstCard := string(hand[0])
	secondCard := string(hand[1])
	thirdCard := string(hand[2])
	// fmt.Printf("checking for %v against %v found %v\n", firstCard, hand, strings.Count(hand, firstCard))
	// fmt.Printf("checking for %v against %v found %v\n", secondCard, hand, strings.Count(hand, secondCard))
	// fmt.Printf("checking for %v against %v found %v\n", thirdCard, hand, strings.Count(hand, thirdCard))
	if strings.Count(hand, firstCard) == 3 {
		h.TypeScore = cardToScore(firstCard)
		return true
	}

	if strings.Count(hand, secondCard) == 3 {
		h.TypeScore = cardToScore(secondCard)
		return true
	}

	if strings.Count(hand, thirdCard) == 3 {
		h.TypeScore = cardToScore(thirdCard)
		return true
	}
	return false
}

func hasSinglePair(h *Hand) (onePair bool) {
	hand := h.Cards
	// Move along the string
	for i := 0; i < len(hand); i++ {
		// Check for one pair
		if strings.Count(hand, string(hand[i])) == 2 {
			onePair = true
			// Now move along the rest of the string and check for other pairs
			for n := i + 1; n < len(hand); n++ {
				// Only check cards that are different to the one we're checking
				if hand[i] != hand[n] {
					if strings.Count(hand, string(hand[n])) == 2 {
						onePair = false
					}
				}
			}
		}
	}
	return
}

func hasTwoPair(h *Hand) (twoPair bool) {
	hand := h.Cards
	// Move along the string
	for i := 0; i < len(hand); i++ {
		// Check for one pair
		if strings.Count(hand, string(hand[i])) == 2 {
			// we have one pair
			// Now move along the rest of the string and check for other pairs
			for n := i + 1; n < len(hand); n++ {
				// Only check cards that are different to the one we're checking
				if hand[i] != hand[n] {
					if strings.Count(hand, string(hand[n])) == 2 {
						twoPair = true
					}
				}
			}
		}
	}
	return
}
