package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input []byte

var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	fmt.Printf("part 1: %d\n", sumDigits(input, nil))
	fmt.Printf("part 2: %d\n", sumDigits(input, numberWords))
}

func sumDigits(input []byte, words []string) int {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sum := 0
	for linereader.Scan() {
		digit1, digit2 := getFirstLast(linereader.Text(), words)
		sum += (int(digit1)-'0')*10 + (int(digit2) - '0')
	}
	return sum
}

func getFirstLast(line string, words []string) (rune, rune) {
	var (
		first, last rune
	)
	for i, c := range line {
		if unicode.IsDigit(c) {
			first = c
			break
		} else if digit := hasAnyPrefix(line[i:], words); digit != -1 {
			first = rune(digit + '0')
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			last = rune(line[i])
			break
		} else if digit := hasAnyPrefix(line[i:], words); digit != -1 {
			last = rune(digit + '0')
			break
		}
	}
	return first, last
}

func hasAnyPrefix(s string, x []string) int {
	if x == nil {
		return -1
	}
	for i, v := range x {
		if strings.HasPrefix(s, v) {
			return i + 1
		}
	}
	return -1
}
