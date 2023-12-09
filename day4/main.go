package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

var sample = []byte(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`)

type Card struct {
	CardNum    int
	Winning    []int
	Picked     []int
	Matching   int
	Value      int
	Duplicates int
}

func main() {
	cards := parseInput(input)
	sum := 0
	for i := range cards {
		cards[i].playGame()
		sum += cards[i].Value
		fmt.Printf("Card %d: matching %d value %d\n", cards[i].CardNum, cards[i].Matching, cards[i].Value)
	}
	fmt.Printf("Part 1: %d\n", sum)
	for i := range cards {
		if cards[i].Matching > 0 {
			for j := i + 1; j < i+cards[i].Matching+1; j++ {
				cards[j].Duplicates += cards[i].Duplicates + 1
			}
		}
	}
	sum = 0
	for i := range cards {
		sum += cards[i].Duplicates + 1
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func (c *Card) playGame() {
	c.Matching = 0
	for _, p := range c.Picked {
		for _, w := range c.Winning {
			if p == w {
				c.Matching++
				break
			}
		}
	}
	if c.Matching > 0 {
		c.Value = 1
		for i := 0; i < c.Matching-1; i++ {
			c.Value *= 2
		}
	}
}

func parseInput(input []byte) []Card {
	regex1 := regexp.MustCompile(`Card \s*(\d+): ([0-9 ]+) \| ([0-9 ]+)`)
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	cards := make([]Card, 0)
	for linereader.Scan() {
		if matches := regex1.FindStringSubmatch(linereader.Text()); matches != nil {
			card := Card{CardNum: len(cards) + 1}
			for _, n := range strings.Split(matches[2], " ") {
				if len(n) == 0 {
					continue
				}
				if i, err := strconv.Atoi(n); err == nil {
					card.Winning = append(card.Winning, i)
				} else {
					panic(err)
				}
			}
			for _, n := range strings.Split(matches[3], " ") {
				if len(n) == 0 {
					continue
				}
				if i, err := strconv.Atoi(n); err == nil {
					card.Picked = append(card.Picked, i)
				} else {
					panic(err)
				}
			}
			cards = append(cards, card)
		} else {
			fmt.Println(linereader.Text())
			panic("no match")
		}
	}
	return cards
}
