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

var sample = []byte(`Time:      7  15   30
Distance:  9  40  200`)

func main() {
	part1()
	part2()
}

func part1() {
	times, distances := parseInput(sample)
	product := 1
	for i := range times {
		winCount := countWins(times[i], distances[i])
		fmt.Println("Time:", times[i], "Distance:", distances[i], "Wins:", winCount)
		product *= winCount
	}
	fmt.Println("Product:", product)
}

func countWins(time int, distance int) int {
	wins := 0
	for i := 1; i < time; i++ {
		speed := i
		traveled := speed * (time - i)
		if traveled > distance {
			wins++
		}
	}
	return wins
}

func parseInput(input []byte) ([]int, []int) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	times := make([]int, 0)
	distances := make([]int, 0)
	row := 0
	for linereader.Scan() {
		parts := strings.Split(linereader.Text(), ":")
		a := strings.Fields(parts[1])
		for _, v := range a {
			if i, err := strconv.Atoi(v); err != nil {
				panic(err)
			} else {
				switch row {
				case 0:
					times = append(times, i)
				case 1:
					distances = append(distances, i)
				default:
					panic("too many rows")
				}
			}
		}
		row++
	}
	return times, distances
}

func part2() {
	time, distance := parseInput2(input)
	winCount := countWins2(time, distance)
	fmt.Println("Time:", time, "Distance:", distance, "Wins:", winCount)
}

// countWins2 runs a binary search for the first value that wins
// This is a hack as it turns out I could have used the quadratic equation that
// gives the distance traveled for any hold time and solved for the time that
// gives the distance
func countWins2(time int, distance int) int {
	// Pretend time is an array of integers from 1 to time
	// Do a binary search for the first value that wins
	span := time / 2
	t := span
	for t > 1 {
		T0, Tp := t*(time-t) > distance, (t-1)*(time-(t-1)) > distance
		// fmt.Printf("span=%d, t=%d, %t %t\n", span, t, Tp, T0)
		if !Tp && T0 {
			return time - 1 - (t-1)*2
		}
		span /= 2
		if span == 0 {
			span = 1
		}
		if !Tp && !T0 {
			t += span
			continue
		}
		t -= span
	}
	panic("not found")
	return 0
}

func parseInput2(input []byte) (int, int) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	row := 0
	time, distance := 0, 0
	for linereader.Scan() {
		parts := strings.Split(linereader.Text(), ":")
		a := strings.Fields(parts[1])
		if i, err := strconv.Atoi(strings.Join(a, "")); err != nil {
			panic(err)
		} else {
			switch row {
			case 0:
				time = i
			case 1:
				distance = i
			default:
				panic("too many rows")
			}
		}
		row++
	}
	return time, distance
}
