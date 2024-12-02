package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() [][]int {
	f, err := os.Open("02.txt")
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

func main() {
	rows := readFile()

	validRows := 0
	for _, r := range rows {
		increaseCheck := inspectRowIncrease(r)
		decreaseCheck := inspectRowDecrease(r)

		if increaseCheck == true || decreaseCheck == true {
			validRows += 1
		} else {
			// fmt.Printf("%v %v %v\n", r, increaseCheck, decreaseCheck)
		}
	}
	fmt.Printf("Valid rows: %v\n", validRows)
}

func inspectRowIncrease(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		current := row[i]
		next := row[i+1]
		if next > current+3 {
			return false
		}
		if next <= current {
			return false
		}
	}
	return true
}

func inspectRowDecrease(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		current := row[i]
		next := row[i+1]
		if next < current-3 {
			return false
		}
		if next >= current {
			return false
		}
	}
	return true
}
