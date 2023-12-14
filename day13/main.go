package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

var sample = []byte(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`)

// Sample answers:
// part1: 405, part2: 400

func main() {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sum := process(linereader, 0)
	fmt.Println("PART 1: ", sum)
	//assert(sum == 37718)
	linereader = bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sum = process(linereader, 1)
	fmt.Println("PART 2: ", sum)
	//assert(sum == 40995)
}

func process(linereader *bufio.Scanner, allowableSumdges int) int {
	sum := 0
	for {
		if puzzle, good := getNextPuzzle(linereader); !good {
			break
		} else {
			var smudgeCount, rowCount, colCount int
			smudgeCount = allowableSumdges
			rowCount = findReflection(puzzle, &smudgeCount)
			if rowCount == 0 {
				smudgeCount = allowableSumdges
				colCount = findReflection(rotate(puzzle), &smudgeCount)
			}
			sum += colCount + rowCount*100
		}
	}
	return sum
}

func fuzzyCompare(a, b string, smudgeCount *int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if *smudgeCount > 0 {
				*smudgeCount--
			} else {
				return false
			}
		}
	}
	return true
}

// findReflection returns the row number of the reflection if found, otherwise 0
// smudgeCount is the number of characters that can be different between the two rows
// and if the number of differences must be exactly smudgeCount, otherwise 0 is returned
func findReflection(puzzle []string, smudgeCount *int) int {
	startingCount := *smudgeCount
	for r := 0; r < len(puzzle)-1; r++ {
		*smudgeCount = startingCount
		if fuzzyCompare(puzzle[r], puzzle[r+1], smudgeCount) {
			if x := followReflection(puzzle, r, smudgeCount); x > 0 && *smudgeCount == 0 {
				return x
			}
		}
	}
	return 0
}

func followReflection(puzzle []string, r int, smudgeCount *int) int {
	a := r - 1
	b := r + 2
	for a >= 0 && b < len(puzzle) {
		if !fuzzyCompare(puzzle[a], puzzle[b], smudgeCount) {
			return 0
		}
		a--
		b++
	}
	return r + 1 // we index by zero but the puzzle rows are numbered from 1
}

func rotate(puzzle []string) []string {
	newPuzzle := make([]string, 0)
	for c := 0; c < len(puzzle[0]); c++ {
		newPuzzle = append(newPuzzle, getColumn(puzzle, c))
	}
	return newPuzzle
}

func getColumn(puzzle []string, col int) string {
	s := make([]byte, len(puzzle))
	c := 0
	for r := len(puzzle) - 1; r >= 0; r-- {
		s[c] = puzzle[r][col]
		c++
	}
	return string(s)
}

func getNextPuzzle(linereader *bufio.Scanner) ([]string, bool) {
	gotlines := false
	pair := make([]string, 0)
	for linereader.Scan() {
		line := linereader.Text()
		if len(line) == 0 {
			break
		}
		pair = append(pair, line)
		gotlines = true
	}
	return pair, gotlines
}
