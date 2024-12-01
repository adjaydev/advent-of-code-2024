package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	listOne := []int{}
	listTwo := []int{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	counter := 0

	// Readlines
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		if counter%2 == 0 {
			listOne = append(listOne, number)
		} else {
			listTwo = append(listTwo, number)
		}
		counter += 1
	}

	slices.Sort(listOne)
	slices.Sort(listTwo)

	// Sum distances
	diff := 0
	for i := range listOne {
		number := listOne[i] - listTwo[i]
		if number < 0 {
			number = number * -1
		}
		diff += number
	}

	fmt.Printf("Total difference: %v\n", diff)

	diff = 0
	for _, v := range listOne {
		countsInList := 0
		for _, w := range listTwo {
			if v == w {
				countsInList += 1
			}
		}
		diff = diff + (v * countsInList)
	}

	fmt.Printf("Total doubles: %v\n", diff)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
