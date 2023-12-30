package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

// Number is a custom type set of constraints extending the Float and Integer type set from the experimental constraints package.
type Number interface {
	constraints.Float | constraints.Integer
}

//go:embed input.txt
var input string

var sample = string(`.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`)

/*
.??..??...?##. 1,1,3

*/

var arr [][]string

func main() {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		total += part1(line)
		break
	}
	fmt.Printf("Part 1: %d\n", total)
}

func part1(data string) int {
	a := strings.Split(data, " ")
	pattern := a[0]
	numbers := csvToInts(a[1], 1)
	maxInterval := len(pattern) - (sum(numbers) + len(numbers) - 1) + 1
	arr = buildInput(numbers, maxInterval)
	possibleArrangements := 0
	combine(0, pattern, func(combination string) {
		fmt.Println(pattern, combination)
		if matcher(pattern, combination) {
			possibleArrangements++
		}
	})
	fmt.Printf("%-32.32s\t Possible arrangements: %d\n", data, possibleArrangements)
	return possibleArrangements
}

func csvToInts(data string, repetitions int) []int {
	var numbers []int
	for i := 0; i < repetitions; i++ {
		for _, value := range strings.Split(data, ",") {
			if i, err := strconv.Atoi(value); err != nil {
				panic(err)
			} else {
				numbers = append(numbers, i)
			}
		}
	}
	return numbers
}

func buildInput(numbers []int, maxInterval int) [][]string {
	var input [][]string
	intervals := make([]string, maxInterval+1)
	for i := 0; i <= maxInterval; i++ {
		intervals[i] = strings.Repeat(".", i)
	}
	input = append(input, intervals)
	for _, value := range numbers {
		input = append(input, []string{strings.Repeat("#", value)}, intervals[1:])
	}
	input = append(input[0:len(input)-1], intervals)
	return input
}

// matcher is a function that takes a pattern and an input and returns true if the input matches the pattern.
// Pattern characters are either a '?' (unknown), '.' (operation), or '#' (damaged)
// and the input characters are either a '.', or '#'.
func matcher(pattern string, input string) bool {
	if len(pattern) != len(input) {
		return false
	}
	for index, value := range pattern {
		if value != '?' && value != rune(input[index]) {
			return false
		}
	}
	return true
}

func sum[T Number](numbers []T) T {
	var total T
	for _, value := range numbers {
		total += value
	}
	return total
}

func combine(index int, pattern string, f func(string)) {
	if index == len(arr) {
		f("")
		return
	}

	for _, value := range arr[index] {
		combine(index+1, pattern, func(combination string) {
			f(value + combination)
		})
	}
}
