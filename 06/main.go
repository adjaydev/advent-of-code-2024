package main

import (
	"bufio"
	"io"
	fmt "log"
	"os"
)

// Positions
const up = "^"
const down = "v"
const left = "<"
const right = ">"
const wall = "#"
const empty = "."
const visited = "X"

func main() {

	ctn := readFileContents("test.txt")
	lCount := len(ctn)
	pCount := len(ctn[0])

	iters := 3

	for i := 0; i < iters; i++ {
		for i := 0; i < pCount; i++ {
			newLine := make([]string, pCount)
			for j := 0; j < lCount; j++ {
				position := string(ctn[i][j])
				if playerFound(position) {
					playerView := playerView(position)
					if playerView == up {
						// nextStep := string(ctn[i-1][j])
						// if nextStep != wall {
						ctn[i-1][j] = up
						ctn[i][j] = visited
						// }
					}
					fmt.Println("Player found!")
				}
				newLine[j] = position
			}
			fmt.Println(newLine)
		}
		fmt.Println()
	}
}

// func playerStep(i, j int, view string, ctn *[]string) {
// 	if view == up {
// 		nextStep := string(ctn[i-1][j])
// 		if nextStep != wall {
// 			ctn[i-1][j] = nextStep
// 			ctn[i][j] = visited
// 		}
// 	}
// }

func playerFound(p string) bool {
	if p == up || p == down || p == left || p == right {
		return true
	}
	return false
}

func playerView(p string) string {
	if p == up {
		return up
	}

	if p == down {
		return down
	}
	if p == left {
		return left
	}
	if p == up {
		return right
	}
	panic("Error finding player view.")
}

func readFileContents(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	if err := scanner.Err(); err != nil {
		fmt.Fatal(err)
	}

	// Count total lines
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	f.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(f)

	// Read and process lines
	lines := make([]string, lineCount)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines[i] = line
		i++
	}

	return lines
}
