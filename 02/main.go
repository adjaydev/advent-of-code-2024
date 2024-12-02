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
	f, err := os.Open("02c.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	validRows := 0
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

	for _, r := range rows {
		counter := 0
		rowLength := len(r) - 1
		increaseCheck := true
		decreaseCheck := true
		errorsFoundIncrease := 0
		errorsFoundDecrease := 0

		// Iterate for increase
		for counter < rowLength {
			currentValue := r[counter]
			nexValue := r[counter+1]
			maxNextValue := currentValue + 3
			if nexValue > maxNextValue || nexValue <= currentValue {
				errorsFoundIncrease += 1
			}
			if errorsFoundIncrease > 1 {
				increaseCheck = false
			}
			counter += 1
		}

		counter = 0
		// Iterate for decrease
		for counter < rowLength {
			currentValue := r[counter]
			nexValue := r[counter+1]
			minNextValue := currentValue - 3
			if nexValue < minNextValue || nexValue >= currentValue {
				errorsFoundDecrease += 1
			}
			if errorsFoundDecrease > 1 {
				decreaseCheck = false
			}
			counter += 1
		}

		if increaseCheck == true || decreaseCheck == true {
			validRows += 1
			fmt.Printf("%v %v %v %v %v\n", r, increaseCheck, errorsFoundIncrease, decreaseCheck, errorsFoundDecrease)
		} else {
		}
	}
	fmt.Printf("Valid rows: %v\n", validRows)

}
