package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/constraints"
	"gonum.org/v1/gonum/stat/combin"
	"strings"
)

//go:embed input.txt
var input string

var sample = string(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`)

type Point struct {
	x, y int
}

const Galaxy = '#'

func main() {
	universe := strings.Split(input, "\n")
	fmt.Printf("Part 1: %d\n", totalGalaxyDistances(universe, 2))
	fmt.Printf("Part 2: %d\n", totalGalaxyDistances(universe, 1_000_000))
}

func totalGalaxyDistances(universe []string, expansionRatio int) int {
	emptyRows := findEmptyRows(universe)
	emptyCols := findEmptyCols(universe)

	// Build list of galaxies positions with expansion taken into account
	galaxies := make([]Point, 0)
	for r, row := range universe {
		for c, char := range row {
			if char == Galaxy {
				emptyRowsBefore := countBefore(emptyRows, r) * (expansionRatio - 1)
				emptyColsBefore := countBefore(emptyCols, c) * (expansionRatio - 1)
				galaxies = append(galaxies, Point{r + emptyRowsBefore, c + emptyColsBefore})
			}
		}
	}

	// Sum up all distances between galaxies
	sum := 0
	for _, v := range combin.Combinations(len(galaxies), 2) {
		sum += distance(galaxies[v[0]], galaxies[v[1]])
	}
	return sum
}

func countBefore(elements []int, idx int) int {
	c := 0
	for _, v := range elements {
		if v < idx {
			c++
		}
	}
	return c
}

func findEmptyRows(galaxy []string) []int {
	emptyRows := make([]int, 0)
	emptyRow := strings.Repeat(".", len(galaxy[0]))
	for r, row := range galaxy {
		if row == emptyRow {
			emptyRows = append(emptyRows, r)
		}
	}
	return emptyRows
}

func findEmptyCols(galaxy []string) []int {
	emptyCols := make([]int, 0)
	for c := 0; c < len(galaxy[0]); c++ {
		for r := 0; r < len(galaxy); r++ {
			if galaxy[r][c] != '.' {
				break
			}
			if r == len(galaxy)-1 {
				emptyCols = append(emptyCols, c)
			}
		}
	}
	return emptyCols
}

func distance(a, b Point) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
