package main

import (
	"bufio"
	fmt "log"
	"os"
)

// Positions
const (
	up      = '^'
	down    = 'v'
	left    = '<'
	right   = '>'
	wall    = '#'
	empty   = '.'
	visited = 'X'
)

func main() {
	ctn := readRunes("input.txt")
	yCount := len(ctn)
	xCount := len(ctn[0])
	playing := true
	turnAround := map[rune]rune{
		up:    right,
		right: down,
		down:  left,
		left:  up,
	}

	for playing {
		for y := 0; y < xCount; y++ {
			newLine := make([]rune, xCount)
			for x := 0; x < yCount; x++ {
				position := rune(ctn[y][x])
				if playerFound(position) {
					playerView := position
					switch playerView {
					case up:
						if !stepAllowed(y-1, x, xCount, yCount) {
							playing = false
							break
						}
						nextStep := rune(ctn[y-1][x])
						ctn[y][x] = visited
						if nextStep != wall {
							ctn[y-1][x] = up
						} else {
							ctn[y][x+1] = turnAround[up]
						}
					case right:
						if !stepAllowed(y, x+1, xCount, yCount) {
							playing = false
							break
						}
						nextStep := rune(ctn[y][x+1])
						ctn[y][x] = visited
						if nextStep != wall {
							ctn[y][x+1] = right
						} else {
							ctn[y][x] = turnAround[right]
						}
					case down:
						if !stepAllowed(y+1, x, xCount, yCount) {
							playing = false
							break
						}
						nextStep := rune(ctn[y+1][x])
						ctn[y][x] = visited
						if nextStep != wall {
							ctn[y+1][x] = down
						} else {
							ctn[y][x] = turnAround[down]
						}
					case left:
						if !stepAllowed(y, x-1, xCount, yCount) {
							playing = false
							break
						}
						nextStep := rune(ctn[y][x-1])
						ctn[y][x] = visited
						if nextStep != wall {
							ctn[y][x-1] = left
						} else {
							ctn[y][x] = turnAround[left]
						}
					}
				}
				newLine[x] = position
			}
			// fmt.Println(string(newLine))
		}

	}

	steps := 0
	for _, row := range ctn {
		for _, position := range row {
			if position != empty && position != wall {
				steps++
			}
		}
	}
	fmt.Printf("Steps taken: %v", steps)
}

func move(x, y int, ctn *[][]rune) {

}

func stepAllowed(x, y, xMax, yMax int) bool {
	if x < 0 || x > xMax-1 {
		return false
	}
	if y < 0 || y > yMax-1 {
		return false
	}
	return true
}

func playerFound(p rune) bool {
	if p == up || p == down || p == left || p == right {
		return true
	}
	return false
}

func readRunes(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return matrix
}
