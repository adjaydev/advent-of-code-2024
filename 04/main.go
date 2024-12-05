package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	ctn := readFileContents("input.txt")
	lineLength := len(ctn[0])
	ctnLength := len(ctn)

	horizontalCount := checkHorizontal(ctn, lineLength)
	verticalCount := checkVertical(ctn, lineLength, ctnLength)
	slash := checkSlash(ctn, lineLength, ctnLength)
	backslash := checkBackslash(ctn, lineLength, ctnLength)

	log.Println(slash + backslash + verticalCount + horizontalCount)
}

func checkBackslash(ctn []string, lineLength, lines int) int {
	maxRowPos := lines - 3
	c := 0
	for i := 0; i < maxRowPos; i++ {
		for j := lineLength; j > 3; j-- {
			word := ctn[i][j-1:j] + ctn[i+1][j-2:j-1] + ctn[i+2][j-3:j-2] + ctn[i+3][j-4:j-3]
			if word == "XMAS" || word == "SAMX" {
				c++
			}
		}
	}
	return c
}

func checkSlash(ctn []string, lineLength, lines int) int {
	maxRowPos := lines - 3
	maxCharPos := lineLength - 3
	c := 0
	for i := 0; i < maxRowPos; i++ {
		for j := 0; j < maxCharPos; j++ {
			word := ctn[i][j:j+1] + ctn[i+1][j+1:j+2] + ctn[i+2][j+2:j+3] + ctn[i+3][j+3:j+4]
			if word == "XMAS" || word == "SAMX" {
				c++
			}
		}
	}
	return c
}

func checkVertical(ctn []string, lineLength, lines int) int {
	maxRowPos := lines - 3
	c := 0
	for i := 0; i < maxRowPos; i++ {
		for j := 0; j < lineLength; j++ {
			word := ctn[i][j:j+1] + ctn[i+1][j:j+1] + ctn[i+2][j:j+1] + ctn[i+3][j:j+1]
			if word == "XMAS" || word == "SAMX" {
				c++
			}
		}
	}
	return c
}

func checkHorizontal(ctn []string, lineLength int) int {
	maxCharPos := lineLength - 3
	c := 0
	for _, row := range ctn {
		for i := 0; i < maxCharPos; i++ {
			word := row[i : i+4]
			if word == "XMAS" || word == "SAMX" {
				c++
			}
		}
	}
	return c
}

func readFileContents(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
