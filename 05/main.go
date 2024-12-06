package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	updatePages := readFileContents("test1.txt")
	updateBatch := readFileContents("test2.txt")

	log.Println(updatePages)
	log.Println(updateBatch)
	// TODO: Iterate update batches

	for _, b := range updateBatch[:1] {
		batch := strings.Split(b, ",")

		// Check update lines
		for i := range batch {
			activePage := batch[i : i+1][0]
			// activePage := string(row[i+1])
			preList := batch[:i+1]
			postList := batch[i+1:]
			log.Printf("pre: %v - active: %v - post: %v", preList, activePage, postList)

			potential := findUpdateLines(updatePages, activePage)
			log.Printf("Potential: %v", potential)
			// for _, updateLine := range ctn1 {
			// 	updates := strings.Split(updateLine, "|")
			// 	isValidUpdate(activePage, preList, postList, updates)
			// }
		}
	}

}

func isValidUpdate(active string, pre, post, updateLine []string) {
	log.Printf("updateline: %v %v %v", updateLine, updateLine[0], updateLine[1])
}

func findUpdateLines(lines []string, activePage string) []string {
	updates := []string{}
	for _, v := range lines {
		line := strings.Split(v, "|")
		if line[0] == activePage || line[1] == activePage {
			updates = append(updates, v)
		}
	}
	return updates
}

func checkUpdateLines(row []string) {
	for i := range row {
		preList := row[:i+1]
		postList := row[i+1:]
		log.Printf("pre: %v - post: %v", preList, postList)
	}
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
