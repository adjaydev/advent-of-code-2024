package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	removeThese := readFileContents("invalid_chars.txt")
	contents := readFileContents("03.txt")
	newContents := strings.Join(contents, " ")

	for _, v := range removeThese {
		newContents = strings.ReplaceAll(newContents, v, "")
	}

	splitContens := strings.SplitAfter(newContents, ")")

	fmt.Println(newContents)

	for _, v := range splitContens {
		fmt.Println(v)
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

	lines := []string{}

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}
	return lines
}
