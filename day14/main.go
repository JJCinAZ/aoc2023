package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var sample = string(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)

var (
	northCache, southCache, westCache, eastCache map[string]string
)

func main() {
	day14()
	northCache = make(map[string]string, 8192)
	southCache = make(map[string]string, 8192)
	westCache = make(map[string]string, 8192)
	eastCache = make(map[string]string, 8192)
	part1()
	part2()
}

func part1() {
	grid := input
	grid = slideNorth(grid)
	weight := getLoad(grid)
	fmt.Printf("Part 1: %d\n", weight)
}

func part2() {
	grid := input
	for i := 0; i < 1000000000; i++ {
		grid = slideNorth(grid)
		grid = slideWest(grid)
		grid = slideSouth(grid)
		grid = slideEast(grid)
		if i%1000000 == 0 {
			fmt.Printf("%f\r", float64(i)/1000000000*100)
		}
	}
	weight := getLoad(grid)
	fmt.Printf("Part 2: %d\n", weight)
}

func slideNorth(grid string) string {
	if val, ok := northCache[grid]; ok {
		return val
	}
	x := strings.Split(grid, "\n")
	x = rotateCCW(x)
	slideLeft(x)
	val := strings.Join(rotateCW(x), "\n")
	northCache[grid] = val
	return val
}

func slideSouth(grid string) string {
	if val, ok := southCache[grid]; ok {
		return val
	}
	x := strings.Split(grid, "\n")
	x = rotateCCW(x)
	slideRight(x)
	val := strings.Join(rotateCW(x), "\n")
	southCache[grid] = val
	return val
}

func slideWest(grid string) string {
	if val, ok := westCache[grid]; ok {
		return val
	}
	x := strings.Split(grid, "\n")
	slideLeft(x)
	val := strings.Join(x, "\n")
	westCache[grid] = val
	return val
}

func slideEast(grid string) string {
	if val, ok := eastCache[grid]; ok {
		return val
	}
	x := strings.Split(grid, "\n")
	slideRight(x)
	val := strings.Join(x, "\n")
	eastCache[grid] = val
	return val
}

func slideLeft(grid []string) {
	for i := 0; i < len(grid); i++ {
		parts := strings.Split(grid[i], "#")
		for j := 0; j < len(parts); j++ {
			c := 0
			for k := 0; k < len(parts[j]); k++ {
				if parts[j][k] == 'O' {
					c++
				}
			}
			parts[j] = strings.Repeat("O", c) + strings.Repeat(".", len(parts[j])-c)
		}
		grid[i] = strings.Join(parts, "#")
	}
}

func slideRight(grid []string) {
	for i := 0; i < len(grid); i++ {
		parts := strings.Split(grid[i], "#")
		for j := 0; j < len(parts); j++ {
			c := 0
			for k := 0; k < len(parts[j]); k++ {
				if parts[j][k] == 'O' {
					c++
				}
			}
			parts[j] = strings.Repeat(".", len(parts[j])-c) + strings.Repeat("O", c)
		}
		grid[i] = strings.Join(parts, "#")
	}
}

func getLoad(grid string) int {
	var weight int
	lines := strings.Split(grid, "\n")
	height := len(lines)
	for i := 0; i < len(lines); i++ {
		weight += strings.Count(lines[i], "O") * (height - i)
	}
	return weight
}

func rotateCCW(grid []string) []string {
	var newGrid []string
	for w := len(grid[0]) - 1; w >= 0; w-- {
		var line []byte
		for j := 0; j < len(grid); j++ {
			line = append(line, grid[j][w])
		}
		newGrid = append(newGrid, string(line))
	}
	return newGrid
}

func rotateCW(grid []string) []string {
	var newGrid []string
	for j := 0; j < len(grid[0]); j++ {
		var line []byte
		for l := len(grid) - 1; l >= 0; l-- {
			line = append(line, grid[l][j])
		}
		newGrid = append(newGrid, string(line))
	}
	return newGrid
}

func printGrid(grid string) {
	fmt.Println(grid, "\n------------------")
}
