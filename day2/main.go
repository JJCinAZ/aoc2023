package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

type reveal struct {
	r, g, b int
}

func main() {
	part1(input)
	part2(input)
}

func part1(input []byte) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sum := 0
	for linereader.Scan() {
		id, reveals := getID(linereader.Text())
		fmt.Printf("ID: %d, Reveals: %v\n", id, reveals)
		if allPossible(reveals, 12, 13, 14) {
			sum += id
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func part2(input []byte) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sum := 0
	for linereader.Scan() {
		_, reveals := getID(linereader.Text())
		mr, mg, mb := minimums(reveals)
		sum += mr * mg * mb
	}
	fmt.Printf("Sum: %d\n", sum)
}

func minimums(reveals []reveal) (int, int, int) {
	var (
		minR, minG, minB int
	)
	for _, reveal := range reveals {
		if reveal.r > minR {
			minR = reveal.r
		}
		if reveal.g > minG {
			minG = reveal.g
		}
		if reveal.b > minB {
			minB = reveal.b
		}
	}
	return minR, minG, minB
}

func allPossible(reveals []reveal, r, g, b int) bool {
	for _, reveal := range reveals {
		if reveal.r > r {
			return false
		}
		if reveal.g > g {
			return false
		}
		if reveal.b > b {
			return false
		}
	}
	return true
}

func getID(line string) (int, []reveal) {
	var (
		id int
	)
	parts := strings.Split(line, ":")
	id, _ = strconv.Atoi(parts[0][5:])
	reveals := strings2reveal(strings.Split(parts[1], ";"))
	return id, reveals
}

func strings2reveal(parts []string) []reveal {
	var reveals []reveal
	for _, part := range parts {
		var reveal reveal
		colors := strings.Split(part, ",")
		for _, color := range colors {
			switch {
			case strings.HasSuffix(color, " red"):
				reveal.r, _ = strconv.Atoi(color[1 : len(color)-4])
			case strings.HasSuffix(color, " green"):
				reveal.g, _ = strconv.Atoi(color[1 : len(color)-6])
			case strings.HasSuffix(color, " blue"):
				reveal.b, _ = strconv.Atoi(color[1 : len(color)-5])
			}
		}
		reveals = append(reveals, reveal)
	}
	return reveals
}
