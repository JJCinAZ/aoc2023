package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

var sample = []byte(`2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41`)

type Hand struct {
	Cards    string
	Bid      int
	HandType int
}

const (
	FiveOfAKind  = 6
	FourOfAKind  = 5
	FullHouse    = 4
	ThreeOfAKind = 3
	TwoPair      = 2
	OnePair      = 1
	HighCard     = 0
)

var weights [256]rune

func main() {
	hands := parseInput(input)
	part1(hands)
	sortAndTotal(hands)
	part2(hands)
	sortAndTotal(hands)
}

func sortAndTotal(hands []Hand) {
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.HandType == b.HandType {
			Wa, Wb := weightString(a.Cards), weightString(b.Cards)
			return strings.Compare(Wb, Wa)
		}
		return a.HandType - b.HandType
	})
	total := 0
	for rank, h := range hands {
		fmt.Printf("%d %s %d\n", h.HandType, h.Cards, h.Bid)
		total += h.Bid * (rank + 1)
	}
	fmt.Println("Total:", total)
}

func part1(hands []Hand) {
	weights['2'] = 'M'
	weights['3'] = 'L'
	weights['4'] = 'K'
	weights['5'] = 'J'
	weights['6'] = 'I'
	weights['7'] = 'H'
	weights['8'] = 'G'
	weights['9'] = 'F'
	weights['T'] = 'E'
	weights['J'] = 'D'
	weights['Q'] = 'C'
	weights['K'] = 'B'
	weights['A'] = 'A'
	for i := range hands {
		hands[i].HandType = getHandType1(hands[i].Cards)
	}

}

func part2(hands []Hand) {
	weights['J'] = 'N'
	weights['2'] = 'M'
	weights['3'] = 'L'
	weights['4'] = 'K'
	weights['5'] = 'J'
	weights['6'] = 'I'
	weights['7'] = 'H'
	weights['8'] = 'G'
	weights['9'] = 'F'
	weights['T'] = 'E'
	weights['Q'] = 'C'
	weights['K'] = 'B'
	weights['A'] = 'A'
	for i := range hands {
		hands[i].HandType = getHandType2(hands[i].Cards)
	}
}

func weightString(s string) string {
	var b strings.Builder
	for _, r := range s {
		b.WriteRune(weights[r])
	}
	return b.String()
}

func getHandType1(cards string) int {
	var cardCounts [13]int
	for i := 0; i < len(cards); i++ {
		switch cards[i] {
		case '2':
			cardCounts[0]++
		case '3':
			cardCounts[1]++
		case '4':
			cardCounts[2]++
		case '5':
			cardCounts[3]++
		case '6':
			cardCounts[4]++
		case '7':
			cardCounts[5]++
		case '8':
			cardCounts[6]++
		case '9':
			cardCounts[7]++
		case 'T':
			cardCounts[8]++
		case 'J':
			cardCounts[9]++
		case 'Q':
			cardCounts[10]++
		case 'K':
			cardCounts[11]++
		case 'A':
			cardCounts[12]++
		}
	}
	threes, twos := 0, 0
	for _, count := range cardCounts {
		if count == 5 {
			return FiveOfAKind
		}
		if count == 4 {
			return FourOfAKind
		}
		if count == 3 {
			threes++
		}
		if count == 2 {
			twos++
		}
	}
	if threes == 1 && twos == 1 {
		return FullHouse
	}
	if threes == 1 {
		return ThreeOfAKind
	}
	if twos == 2 {
		return TwoPair
	}
	if twos == 1 {
		return OnePair
	}
	return HighCard
}

func getHandType2(cards string) int {
	var (
		cardCounts [13]int
		jokers     int
	)
	for i := 0; i < len(cards); i++ {
		switch cards[i] {
		case '2':
			cardCounts[0]++
		case '3':
			cardCounts[1]++
		case '4':
			cardCounts[2]++
		case '5':
			cardCounts[3]++
		case '6':
			cardCounts[4]++
		case '7':
			cardCounts[5]++
		case '8':
			cardCounts[6]++
		case '9':
			cardCounts[7]++
		case 'T':
			cardCounts[8]++
		case 'Q':
			cardCounts[9]++
		case 'K':
			cardCounts[10]++
		case 'A':
			cardCounts[11]++
		case 'J':
			jokers++
		}
	}
	fours, threes, twos := 0, 0, 0
	for _, count := range cardCounts {
		if count == 5 {
			return FiveOfAKind
		}
		if count == 4 {
			fours++
		}
		if count == 3 {
			threes++
		}
		if count == 2 {
			twos++
		}
	}
	switch jokers {
	case 4, 5:
		return FiveOfAKind
	case 3:
		switch {
		case twos == 1:
			return FiveOfAKind
		default:
			return FourOfAKind
		}
	case 2:
		switch {
		case threes == 1:
			return FiveOfAKind
		case twos == 1:
			return FourOfAKind
		default:
			return ThreeOfAKind
		}
	case 1:
		switch {
		case fours == 1:
			return FiveOfAKind
		case threes == 1:
			return FourOfAKind
		case twos == 2:
			return FullHouse
		case twos == 1:
			return ThreeOfAKind
		default:
			return OnePair
		}
	}
	if fours == 1 {
		return FourOfAKind
	}
	if threes == 1 && twos == 1 {
		return FullHouse
	}
	if threes == 1 {
		return ThreeOfAKind
	}
	if twos == 2 {
		return TwoPair
	}
	if twos == 1 {
		return OnePair
	}
	return HighCard
}

func parseInput(input []byte) []Hand {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	hands := make([]Hand, 0)
	for linereader.Scan() {
		parts := strings.Fields(linereader.Text())
		h := Hand{Cards: parts[0]}
		if i, err := strconv.Atoi(parts[1]); err != nil {
			panic(err)
		} else {
			h.Bid = i
		}
		hands = append(hands, h)
	}
	return hands
}
