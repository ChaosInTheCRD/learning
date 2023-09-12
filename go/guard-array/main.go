package main

import (
	"fmt"
)

const (
	empty  = -1
	locked = -2
	guard  = 0
	// sets the size of the 2D array (x and y)
	axis_max = 5
)

type position struct {
	x int
	y int
}

// helpful function to figure out what the path of the steps were
func flashArray(c position, lockedPos map[position]bool, in [][]int) {
	fmt.Println("")
	o := in[c.x][c.y]
	if isValid(in, lockedPos, c) {
		in[c.x][c.y] = 99
	}

	for _, row := range in {
		fmt.Println(row)
	}

	fmt.Println("")
	in[c.x][c.y] = o
}

func isLocked(lockedPos map[position]bool, p position) bool {
	return lockedPos[position{x: p.x, y: p.y}]
}

func isValid(in [][]int, lockedPos map[position]bool, p position) bool {
	if p.x >= 0 && p.y >= 0 && p.x < axis_max && p.y < axis_max && !isLocked(lockedPos, p) {
		return true
	}

	return false
}

func isQuick(c int, n int) bool {
	if n < c || c == empty {
		return true
	}

	return false
}

func findGuards(in [][]int) ([][]int, []position) {
	guards := []position{}
	for i := range in {
		for j := range in {
			if in[i][j] == guard {
				guards = append(guards, position{x: i, y: j})
				in[i][j] = 0
			}
		}
	}

	return in, guards
}

func findLockedDoors(in [][]int) ([][]int, map[position]bool) {
	doors := map[position]bool{}
	for i := range in {
		for j := range in {
			if in[i][j] == locked {
				doors[position{x: i, y: j}] = true
				in[i][j] = 0
			}
		}
	}

	return in, doors
}

func takeSteps(in [][]int, lockedDoors map[position]bool, guard position) [][]int {
	pos := []position{{x: guard.x, y: guard.y}}

	for m := 0; m <= axis_max*2; m++ {
		pos, in = takeStep(in, pos, lockedDoors, m)
	}

	return in
}

func takeStep(in [][]int, pos []position, lockedDoors map[position]bool, m int) ([]position, [][]int) {
	newPos := []position{}
	for _, c := range pos {
		if !isLocked(lockedDoors, position{x: c.x + 1, y: c.y}) && isValid(in, lockedDoors, position{x: c.x + 1, y: c.y}) {
			if isQuick(in[c.x+1][c.y], m) {
				in[c.x+1][c.y] = m + 1
			}
			newPos = append(newPos, position{x: c.x + 1, y: c.y})
		}

		if !isLocked(lockedDoors, position{x: c.x - 1, y: c.y}) && isValid(in, lockedDoors, position{x: c.x - 1, y: c.y}) {
			if isQuick(in[c.x-1][c.y], m) {
				in[c.x-1][c.y] = m + 1
			}
			newPos = append(newPos, position{x: c.x - 1, y: c.y})
		}

		if !isLocked(lockedDoors, position{x: c.x, y: c.y + 1}) && isValid(in, lockedDoors, position{x: c.x, y: c.y + 1}) {
			if isQuick(in[c.x][c.y+1], m) {
				in[c.x][c.y+1] = m + 1
			}
			newPos = append(newPos, position{x: c.x, y: c.y + 1})
		}

		if !isLocked(lockedDoors, position{x: c.x, y: c.y - 1}) && isValid(in, lockedDoors, position{x: c.x, y: c.y - 1}) {
			if isQuick(in[c.x][c.y-1], m) {
				in[c.x][c.y-1] = m + 1
			}
			newPos = append(newPos, position{x: c.x, y: c.y - 1})
		}
	}

	return newPos, in
}

func solve(in [][]int) [][]int {
	in, guards := findGuards(in)
	in, lockedDoors := findLockedDoors(in)

	for _, guard := range guards {
		in = takeSteps(in, lockedDoors, guard)
	}

	return in
}

func main() {
	in := [][]int{
		{-1, -1, -1, 0, -1},
		{-1, -2, -1, -1, -1},
		{-1, -2, -1, -1, -1},
		{-1, -2, -1, -1, -1},
		{-1, -2, 0, -1, -1},
	}
	for _, row := range in {
		fmt.Println(row)
	}
	fmt.Println("")
	solution := solve(in)
	for _, row := range solution {
		fmt.Println(row)
	}
}
