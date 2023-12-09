package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

var sample = []byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

type Position struct {
	Row         int
	Column      int
	Symbol      rune
	AdjNumCount int
}

type Number struct {
	Starting     Position
	Length       int
	Value        int
	IsPartNumber bool
}

func main() {
	symbols, numbers := parseInput(input)
	markPartNumbers(symbols, numbers)
	sum := 0
	for _, n := range numbers {
		if n.IsPartNumber {
			sum += n.Value
		}
	}
	fmt.Printf("part 1 answer: %d\n", sum)
	sum = 0
	for i := range symbols {
		if symbols[i].Symbol == '*' {
			product := 1
			count := 0
			for _, n := range numbers {
				p := n.Starting
				for j := 0; j < n.Length; j++ {
					if symbolAdjacent(p, symbols[i:i+1]) {
						product *= n.Value
						count++
						break
					}
					p.Column++
				}
			}
			// If gear was adjacent to exactly two numbers, sum the product of those numbers
			if count == 2 {
				sum += product
			}
		}
	}
	fmt.Printf("part 2 answer: %d\n", sum)
}

func markPartNumbers(symbols []Position, numbers []Number) {
	for i := range numbers {
		p := numbers[i].Starting
		for j := 0; j < numbers[i].Length; j++ {
			if symbolAdjacent(p, symbols) {
				numbers[i].IsPartNumber = true
				break
			}
			p.Column++
		}
	}
}

func symbolAdjacent(p Position, symbols []Position) bool {
	for _, s := range symbols {
		if p.Row == s.Row && (p.Column-1 == s.Column || p.Column+1 == s.Column) {
			return true
		}
		if p.Row-1 == s.Row && (p.Column-1 == s.Column || p.Column == s.Column || p.Column+1 == s.Column) {
			return true
		}
		if p.Row+1 == s.Row && (p.Column-1 == s.Column || p.Column == s.Column || p.Column+1 == s.Column) {
			return true
		}
	}
	return false
}

func parseInput(input []byte) ([]Position, []Number) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	row := 0
	symbols := make([]Position, 0)
	numbers := make([]Number, 0)
	for linereader.Scan() {
		var (
			column     int
			n          Number
			collecting bool
		)
		for _, char := range linereader.Text() {
			switch char {
			case '.':
				if collecting {
					collecting = false
					n.Length = column - n.Starting.Column
					numbers = append(numbers, n)
					n = Number{}
				}
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !collecting {
					collecting = true
					n.Starting = Position{Row: row, Column: column}
					n.Value = int(char - '0')
				} else {
					n.Value = n.Value*10 + int(char-'0')
				}
			default:
				if collecting {
					collecting = false
					n.Length = column - n.Starting.Column
					numbers = append(numbers, n)
					n = Number{}
				}
				symbols = append(symbols, Position{Row: row, Column: column, Symbol: char})
			}
			column++
		}
		if collecting {
			collecting = false
			n.Length = column - n.Starting.Column
			numbers = append(numbers, n)
			n = Number{}
		}
		row++
	}
	return symbols, numbers
}
