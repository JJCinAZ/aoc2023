package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var sample = string(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)

func main() {
	part1()
	part2()

}
func part1() {
	parts := strings.Split(input, ",")
	sum := 0
	for _, part := range parts {
		sum += calcHash(part)
	}
	fmt.Printf("Part1: %d\n", sum)
}

type Lens struct {
	code     string
	focalLen int
}

type Box struct {
	lenses *list.List
}

func part2() {
	boxes := make([]Box, 256)
	for i := range boxes {
		boxes[i].lenses = list.New()
	}
	parts := strings.Split(input, ",")
	for _, part := range parts {
		if part[len(part)-1] == '-' { // part looks like "xxx-"?
			code := part[0 : len(part)-1]
			hash := calcHash(code)
			l := boxes[hash].lenses
			for e := l.Front(); e != nil; e = e.Next() {
				if e.Value.(Lens).code == code {
					l.Remove(e)
					break
				}
			}
		} else { // assume part looks like "xxx=N"
			code := part[0 : len(part)-2]
			focalLen := int(part[len(part)-1] - '0')
			hash := calcHash(code)
			l := boxes[hash].lenses
			found := false
			for e := l.Front(); e != nil; e = e.Next() {
				if e.Value.(Lens).code == code {
					e.Value = Lens{code: code, focalLen: focalLen}
					found = true
					break
				}
			}
			if !found {
				l.PushBack(Lens{code: code, focalLen: focalLen})
			}
		}
	}
	totalFocalLength := 0
	for i, b := range boxes {
		l := b.lenses
		if l.Len() > 0 {
			slot := 1
			for e := l.Front(); e != nil; e = e.Next() {
				x := (i + 1) * slot * e.Value.(Lens).focalLen
				totalFocalLength += x
				slot++
			}
		}
	}
	fmt.Printf("Part2: %d\n", totalFocalLength)
}

func calcHash(data string) int {
	if len(data) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(data); i++ {
		sum = ((sum + int(data[i])) * 17) % 256
	}
	return sum
}
