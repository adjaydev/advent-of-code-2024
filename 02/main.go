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
	counter := 0
	rowLength := len(row) - 1

	for counter < rowLength {
		currentValue := row[counter]
		nexValue := row[counter+1]
		maxNextValue := currentValue + 3
		if nexValue > maxNextValue || nexValue <= currentValue {
			return false
		}
		counter += 1
	}
	return true
}

func inspectRowDecrease(row []int) bool {
	counter := 0
	rowLength := len(row) - 1

	for counter < rowLength {
		currentValue := row[counter]
		nexValue := row[counter+1]
		maxNextValue := currentValue - 3
		if nexValue < maxNextValue || nexValue >= currentValue {
			return false
		}
		counter += 1
	}
	return true
}
