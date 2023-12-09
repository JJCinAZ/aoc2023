package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed input.txt
var input []byte

var sample = []byte(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)

type Tuple struct {
	Parts [2]string
}

func main() {
	data, steps := parseInput(input)
	//part1(data, steps)
	part2(data, steps)
}

func parseInput(input []byte) (map[string]Tuple, string) {
	var (
		regex1 = regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)
		steps  string
	)
	linereader := bufio.NewScanner(bytes.NewReader(input))
	linereader.Split(bufio.ScanLines)
	data := make(map[string]Tuple, 8192)
	row := 0
	for linereader.Scan() {
		line := linereader.Text()
		if len(line) > 0 {
			if row == 0 {
				steps = linereader.Text()
			} else {
				parts := regex1.FindStringSubmatch(linereader.Text())
				data[parts[1]] = Tuple{Parts: [2]string{parts[2], parts[3]}}
			}
		}
		row++
	}
	return data, steps
}

func part1(data map[string]Tuple, steps string) {
	stepIdx := 0
	stepCounter := 0
	nextNode := "AAA"
	for {
		if stepIdx >= len(steps) {
			stepIdx = 0
		}
		node, found := data[nextNode]
		if !found {
			panic("Node not found: " + nextNode)
		}
		if steps[stepIdx] == 'L' {
			nextNode = node.Parts[0]
		} else {
			nextNode = node.Parts[1]
		}
		stepCounter++
		stepIdx++
		if nextNode == "ZZZ" {
			break
		}
	}
	fmt.Println(stepCounter)
}

func part2(data map[string]Tuple, steps string) {
	nextNodes := getAllStartingNodes(data)
	zCounts := make([]int, len(nextNodes))
	for i := range nextNodes {
		stepIdx := 0
		stepCounter := 0
		for {
			if stepIdx >= len(steps) {
				stepIdx = 0
			}
			node, found := data[nextNodes[i]]
			if !found {
				panic("Node not found")
			}
			if steps[stepIdx] == 'L' {
				nextNodes[i] = node.Parts[0]
			} else {
				nextNodes[i] = node.Parts[1]
			}
			stepCounter++
			stepIdx++
			if nextNodes[i][2] == 'Z' {
				break
			}
		}
		zCounts[i] = stepCounter
	}
	fmt.Println(zCounts)
	fmt.Println(LCM(zCounts[0], zCounts[1], zCounts[2:]...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func getAllStartingNodes(data map[string]Tuple) []string {
	var startingNodes []string
	for k, _ := range data {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}
	return startingNodes
}
