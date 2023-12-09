package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

//go:embed sample.txt
var sample []byte

type mapping struct {
	DstStart int
	SrcStart int
	Length   int
}

type translationMap struct {
	Name     string
	Mappings []mapping
	Next     *translationMap
}

func main() {
	translationMaps, seeds := parseInput(input)

	// part 1
	lowestLocation := math.MaxInt64
	for _, seed := range seeds {
		location := translate(seed, translationMaps)
		if location < lowestLocation {
			lowestLocation = location
		}
		fmt.Printf("seed %d: %d\n", seed, location)
	}
	fmt.Printf("part 1 -- lowest location: %d\n", lowestLocation)

	// part 2
	lowestLocation = math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		end := seeds[i] + seeds[i+1]
		for seed := seeds[i]; seed < end; seed++ {
			location := translate(seed, translationMaps)
			if location < lowestLocation {
				lowestLocation = location
			}
		}
	}
	fmt.Printf("part 2 -- lowest location: %d\n", lowestLocation)
}

func translate(location int, translationMaps *translationMap) int {
	for translationMaps != nil {
		for _, mapping := range translationMaps.Mappings {
			if location >= mapping.SrcStart && location < mapping.SrcStart+mapping.Length {
				location = mapping.DstStart + (location - mapping.SrcStart)
				break
			}
		}
		translationMaps = translationMaps.Next
	}
	return location
}

func parseMapLine(line string) mapping {
	regex1 := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
	matches := regex1.FindStringSubmatch(line)
	if matches == nil {
		panic(fmt.Sprintf("no match with mapping line '%s'", line))
	}
	dstStart, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	srcStart, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	length, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}
	return mapping{DstStart: dstStart, SrcStart: srcStart, Length: length}
}

func parseInput(input []byte) (*translationMap, []int) {
	var (
		startingMap, curMap *translationMap
		regex2              = regexp.MustCompile(`^([a-z\-]+) map:$`)
	)
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	seeds := make([]int, 0)
	state := 0
	for linereader.Scan() {
		line := linereader.Text()
		switch state {
		case 0:
			if strings.HasPrefix(line, "seeds: ") {
				for _, n := range strings.Split(line[7:], " ") {
					if len(n) > 0 {
						if i, err := strconv.Atoi(n); err == nil {
							seeds = append(seeds, i)
						} else {
							panic(err)
						}
					}
				}
				state = 1
				continue
			}
		case 1:
			if a := regex2.FindStringSubmatch(line); a != nil {
				newMap := new(translationMap)
				newMap.Name = a[1]
				if startingMap == nil {
					startingMap = newMap
				} else {
					curMap.Next = newMap
				}
				curMap = newMap
				state = 2
				continue
			}
		case 2:
			if len(line) == 0 {
				state = 1
				break
			}
			curMap.Mappings = append(curMap.Mappings, parseMapLine(line))
		}
	}
	return startingMap, seeds
}
