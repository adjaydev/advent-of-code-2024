package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rows := readFileContents("02.txt")

	validRows := 0
	for _, r := range rows {
		if isSafe(r) || isSafeSpecial(r) {
			validRows += 1
		}
	}

	fmt.Printf("Valid rows: %v\n", validRows)

}

func isSafe(row []int) bool {
	increaseCheck := true
	decreaseCheck := true

	for i := 0; i < len(row)-1; i++ {
		diff := row[i] - row[i+1]
		if diff < -3 || diff > +3 || row[i] == row[i+1] {
			return false
		}
		if diff < 0 {
			decreaseCheck = false
		} else if diff > 0 {
			increaseCheck = false
		}
	}
	return decreaseCheck || increaseCheck
}

func isSafeSpecial(row []int) bool {
	for i := 0; i < len(row); i++ {

		// New row without ith value
		newRow := append([]int{}, row[:i]...)
		newRow = append(newRow, row[i+1:]...)

		if isSafe(newRow) {
			return true
		}
	}
	return false
}

func readFileContents(filename string) [][]int {
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

	rows := [][]int{}

	// Read and process lines
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		rowLength := len(row)

		tempRow := make([]int, rowLength)
		for i, v := range row {
			tempRow[i], _ = strconv.Atoi(v)
		}
		rows = append(rows, tempRow)
	}
	return rows
}
