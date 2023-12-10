package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input []byte

var sample = []byte(`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`)

var sample2 = []byte(`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`)

type Tile struct {
	Value    rune
	Visited  bool
	Starting bool
	Inside   bool
}

const (
	Starting = 'S'
	Empty    = '.'
	Vpipe    = '│'
	Hpipe    = '─'
	NEcorner = '└'
	NWcorner = '┘'
	SWcorner = '┐'
	SEcorner = '┌'
)

var (
	translation = map[rune]rune{
		'.': Empty,
		'|': Vpipe,
		'-': Hpipe,
		'L': NEcorner,
		'J': NWcorner,
		'7': SWcorner,
		'F': SEcorner,
		'S': Starting,
	}
	area           [][]Tile
	Height, Width  int
	rStart, cStart int
)

func main() {
	parseInput(input)
	part1()
	part2()
	printarea()
}

// part1 will follow the path until it reaches the starting tile again, marking tiles as visited
func part1() {
	var (
		steps, rCur, cCur, rNext, cNext, rFrom, cFrom int
	)
	rCur, cCur = rStart, cStart
	rFrom, cFrom = -9, -9
	for {
		area[rCur][cCur].Visited = true
		rNext, cNext = moveToNext(rFrom, cFrom, rCur, cCur)
		steps++
		if rNext == rStart && cNext == cStart {
			break
		}
		rFrom, cFrom = rCur, cCur
		rCur, cCur = rNext, cNext
	}
	fmt.Printf("Part 1: total loop length=%d, further end=%d\n", steps, steps/2)
}

// part2 will scan across the columns to find where rows are vertically connected and count
// non-visited tiles as within the loop by counting those between two vertical segments.
// This is the Nonzero-Rule algorithm described in https://en.wikipedia.org/wiki/Nonzero-rule
// or the Even-Odd rule https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule because in our case, the
// curve never crosses itself.
// Warning: the tile map must have already been processed with part1() to have set the
// visit flags.
func part2() {
	var (
		inLoop      bool
		insideCount int
	)
	for r := 0; r < (Height - 1); r++ {
		for c := 0; c < Width; c++ {
			if area[r][c].Visited && area[r+1][c].Visited {
				// If this tile is vertically connected to the tile below it, then toggle inLoop flag
				v := area[r+1][c].Value
				switch area[r][c].Value {
				case Vpipe:
					if v == Vpipe || v == NWcorner || v == NEcorner {
						inLoop = !inLoop
					}
				case SWcorner, SEcorner:
					if v == Vpipe || v == NWcorner || v == NEcorner {
						inLoop = !inLoop
					}
				}
			}
			if inLoop && !area[r][c].Visited {
				area[r][c].Inside = true
				insideCount++
			}
		}
	}
	fmt.Printf("Part 2: %d blocks within the loop\n", insideCount)
}

func moveToNext(rFrom, cFrom, rCur, cCur int) (int, int) {
	directions := map[rune]struct{ dr1, dc1, dr2, dc2 int }{
		Vpipe:    {-1, 0, 1, 0},  // North & South
		Hpipe:    {0, -1, 0, 1},  // West & East
		NEcorner: {-1, 0, 0, 1},  // L
		NWcorner: {-1, 0, 0, -1}, // J
		SWcorner: {0, -1, 1, 0},  // 7
		SEcorner: {1, 0, 0, 1},   // F
	}

	d := directions[area[rCur][cCur].Value]
	// Check one direction
	if rNext, cNext := rCur+d.dr1, cCur+d.dc1; rNext != rFrom || cNext != cFrom {
		return rNext, cNext
	}
	// check the other direction
	if rNext, cNext := rCur+d.dr2, cCur+d.dc2; rNext != rFrom || cNext != cFrom {
		return rNext, cNext
	}
	panic("No next tile found")
	return -1, -1
}

// parseInput will parse the input into a 2D array of Tiles, translating the tile characters for the path
// from simple ASCII to Unicode box drawing characters for a prettier display later
func parseInput(input []byte) {
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	row := 0
	for linereader.Scan() {
		line := linereader.Text()
		if row == 0 {
			Width = len(line)
			Height = Width
			fmt.Printf("Width=%d, Height=%d\n", Width, Height)
			area = make([][]Tile, Height)
			for i := range area {
				area[i] = make([]Tile, Width)
			}
		}
		for col, c := range line {
			area[row][col].Value = translation[c]
			if c == Starting {
				rStart = row
				cStart = col
				area[row][col].Starting = true
			}
		}
		row++
	}
	// Patch starting tile so we have a valid loop
	n, s, e, w := Empty, Empty, Empty, Empty
	if rStart > 0 {
		n = area[rStart-1][cStart].Value
	}
	if rStart < (Height - 1) {
		s = area[rStart+1][cStart].Value
	}
	if cStart > 0 {
		w = area[rStart][cStart-1].Value
	}
	if cStart < (Width - 1) {
		e = area[rStart][cStart+1].Value
	}
	switch {
	case n == Vpipe || n == SWcorner || n == SEcorner:
		switch {
		case s == Vpipe || s == NWcorner || s == NEcorner:
			area[rStart][cStart].Value = Vpipe
		case e == Hpipe || e == NWcorner || e == SWcorner:
			area[rStart][cStart].Value = NEcorner
		case w == Hpipe || w == NEcorner || w == SEcorner:
			area[rStart][cStart].Value = NWcorner
		default:
			panic("No direction found")
		}
	case s == Vpipe || s == NWcorner || s == NEcorner:
		switch {
		case e == Hpipe || e == NWcorner || e == SWcorner:
			area[rStart][cStart].Value = SEcorner
		case w == Hpipe || w == NEcorner || w == SEcorner:
			area[rStart][cStart].Value = SWcorner
		default:
			panic("No direction found")
		}
	case e == Hpipe || e == NWcorner || e == SWcorner:
		switch {
		case w == Hpipe || w == NEcorner || w == SEcorner:
			area[rStart][cStart].Value = Hpipe
		default:
			panic("No direction found")

		}
	default:
		panic("No direction found")
	}
}

func printarea() {
	for _, row := range area {
		s := strings.Builder{}
		for _, tile := range row {
			switch {
			case tile.Starting:
				s.WriteString("\033[42m")
				s.WriteRune(tile.Value)
				s.WriteString("\033[0m")
			case tile.Visited:
				s.WriteString("\033[31m")
				s.WriteRune(tile.Value)
				s.WriteString("\033[0m")
			case tile.Inside:
				s.WriteString("\033[30m\033[46m")
				s.WriteRune(' ')
				s.WriteString("\033[0m")
			default:
				s.WriteRune(' ')
			}
		}
		fmt.Println(s.String())
	}
}
