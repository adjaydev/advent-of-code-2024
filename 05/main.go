package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	updatePages := readFileContents("input1.txt")
	updateBatches := readFileContents("input2.txt")
	total := 0

	// Iterate over batches
	for _, b := range updateBatches {
		batchIsValid := true
		batch := strings.Split(b, ",")
		valueIndexes := findValueIndexes(batch)

		// Validate updatePage
		for _, p := range updatePages {
			page := strings.Split(p, "|")

			if validPage(page, batch) {
				pageAIndex := valueIndexes[page[0]]
				pageBIndex := valueIndexes[page[1]]

				if pageAIndex > pageBIndex {
					batchIsValid = false
					break
				}
			}
		}
		if batchIsValid {
			num, _ := strconv.Atoi(batch[len(batch)/2])
			total += num
		}
	}
	log.Printf("Total: %d", total)

}

// Create and return a map for the value and the indexes in the batch
func findValueIndexes(batch []string) map[string]int {
	valueIndexes := make(map[string]int, len(batch))
	for i, b := range batch {
		valueIndexes[b] = i
	}
	return valueIndexes

}

// Iterate batches over current page an check if page values are in batch
// Return: true if both page values are in batch
func validPage(page, batch []string) bool {
	checkA := false
	checkB := false
	for _, b := range batch {
		if b == page[0] {
			checkA = true
		}
		if b == page[1] {
			checkB = true
		}
	}
	return checkA && checkB
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
