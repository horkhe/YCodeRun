// #15. Speleologist path.
// https://coderun.yandex.ru/problem/speleologist-way?currentPage=2&pageSize=10&rowNumber=15
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var dimension int
	fmt.Scan(&dimension)
	c := newCave(dimension)
	for z := 0; z < dimension; z++ {
		var line string
		fmt.Scanln(&line)
		for y := 0; y < dimension; y++ {
			fmt.Scanln(&line)
			c.setLine(z, y, line)
		}
	}

	pathLen := c.findShortestWayOut()

	fmt.Println(pathLen)
}

type cave struct {
	dimension int
	sections  []int
	fromPos   position
}

func newCave(dimension int) *cave {
	c := cave{
		dimension: dimension,
		sections:  make([]int, dimension*dimension*dimension),
	}
	return &c
}

func (c *cave) setLine(z, y int, line string) {
	for x, ch := range line {
		switch ch {
		case '#':
			c.set(position{x, y, z}, sectionSolid)
		case 'S':
			c.fromPos = position{x, y, z}
			c.set(c.fromPos, sectionStart)
		}
	}
}

func (c *cave) findShortestWayOut() int {
	c.markReachableFrom(c.fromPos)
	//fmt.Print(c.String())
	// Choose the best option.
	minPathLen := math.MaxInt64
	for x := 0; x < c.dimension; x++ {
		for y := 0; y < c.dimension; y++ {
			val := c.get(position{x, y, 0})
			if val <= 0 {
				continue
			}
			if val < minPathLen {
				minPathLen = val
			}
		}
	}
	if minPathLen <= 0 {
		return -1
	}
	return minPathLen
}

func (c *cave) markReachableFrom(fromPos position) {
	c.set(fromPos, -1)
	reachable := c.appendAndMarkReachableFrom(nil, fromPos, 1)
	distance := 2
	for len(reachable) > 0 {
		var nextReachable []position
		for _, pos := range reachable {
			nextReachable = c.appendAndMarkReachableFrom(nextReachable, pos, distance)
		}
		distance++
		reachable = nextReachable
	}
}

func (c *cave) appendAndMarkReachableFrom(positions []position, fromPos position, distance int) []position {
	positions = c.appendAndMarkIfInBounds(positions, position{fromPos.x + 1, fromPos.y, fromPos.z}, distance)
	positions = c.appendAndMarkIfInBounds(positions, position{fromPos.x - 1, fromPos.y, fromPos.z}, distance)
	positions = c.appendAndMarkIfInBounds(positions, position{fromPos.x, fromPos.y + 1, fromPos.z}, distance)
	positions = c.appendAndMarkIfInBounds(positions, position{fromPos.x, fromPos.y - 1, fromPos.z}, distance)
	positions = c.appendAndMarkIfInBounds(positions, position{fromPos.x, fromPos.y, fromPos.z + 1}, distance)
	positions = c.appendAndMarkIfInBounds(positions, position{fromPos.x, fromPos.y, fromPos.z - 1}, distance)
	return positions
}

func (c *cave) appendAndMarkIfInBounds(positions []position, pos position, distance int) []position {
	if c.isOutOfBounds(pos) {
		return positions
	}
	if c.get(pos) != 0 {
		return positions
	}
	c.set(pos, distance)
	positions = append(positions, pos)
	return positions
}

func (c *cave) isOutOfBounds(pos position) bool {
	if pos.x < 0 {
		return true
	}
	if pos.x >= c.dimension {
		return true
	}
	if pos.y < 0 {
		return true
	}
	if pos.y >= c.dimension {
		return true
	}
	if pos.z < 0 {
		return true
	}
	if pos.z >= c.dimension {
		return true
	}
	return false
}

func (c *cave) get(pos position) int {
	return c.sections[(pos.z*c.dimension+pos.y)*c.dimension+pos.x]
}

func (c *cave) set(pos position, val int) {
	c.sections[(pos.z*c.dimension+pos.y)*c.dimension+pos.x] = val
}

func (c *cave) String() string {
	var b strings.Builder
	for z := 0; z < c.dimension; z++ {
		for y := 0; y < c.dimension; y++ {
			for x := 0; x < c.dimension; x++ {
				v := c.get(position{x, y, z})
				if v == sectionSolid {
					b.WriteRune('#')
					continue
				}
				if v == sectionStart {
					b.WriteRune('S')
					continue
				}
				b.WriteString(strconv.Itoa(v))
			}
			b.WriteRune('\n')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

type position struct {
	x, y, z int
}

const (
	sectionStart = -1
	sectionSolid = -2
)
