package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"strings"
	"unicode"
)

//go:embed input.txt
var input []byte

func main() {
	part1(input)
}

func part1(input []byte) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sum := 0
	for linereader.Scan() {
		digit1, digit2 := getFirstLast(linereader.Text())
		sum += (int(digit1)-'0')*10 + (int(digit2) - '0')
	}
}

func getFirstLast(line string) (rune, rune) {
	var (
		first, last rune
		x           = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	)
	for i, c := range line {
		if unicode.IsDigit(c) {
			first = c
			break
		} else if digit := hasAnyPrefix(line[i:], x); digit != -1 {
			first = rune(digit + '0')
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			last = rune(line[i])
			break
		} else if digit := hasAnyPrefix(line[i:], x); digit != -1 {
			last = rune(digit + '0')
			break
		}
	}
	return first, last
}

func hasAnyPrefix(s string, x []string) int {
	for i, v := range x {
		if strings.HasPrefix(s, v) {
			return i + 1
		}
	}
	return -1
}
