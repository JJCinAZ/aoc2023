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

var sample = []byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

func main() {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	sumPrev, sumNext := 0, 0
	for linereader.Scan() {
		data := make([]int, 0)
		a := strings.Fields(linereader.Text())
		for _, v := range a {
			data = append(data, mustParseInt(v))
		}
		f, l := calcDeltas(data)
		prevNumber, nextNumber := data[0]-f, data[len(data)-1]+l
		sumPrev += prevNumber
		sumNext += nextNumber
	}
	fmt.Printf("part1: %d, part2: %d\n", sumNext, sumPrev)
}

func calcDeltas(data []int) (int, int) {
	deltas := make([]int, 0, len(data)-1)
	nonZeroPresent := false
	for i := 1; i < len(data); i++ {
		d := data[i] - data[i-1]
		deltas = append(deltas, d)
		if d != 0 {
			nonZeroPresent = true
		}
	}
	if nonZeroPresent {
		f, l := calcDeltas(deltas)
		return deltas[0] - f, deltas[len(deltas)-1] + l
	}
	return 0, 0
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
