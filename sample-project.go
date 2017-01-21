// Sample project by Fernando Nara
// 20/01/17

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Looks for .txt extensions in all folders
func processTxtFiles(outputFile *os.File) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			lineCount := countLines(path)
			writeToFile(path, lineCount, outputFile)
		}
		return nil
	}
}

// Count lines and print results for each file found
func countLines(path string) int {
	inputFile, err := os.Open(path)
	check(err)
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}

// Write results to output file
func writeToFile(path string, lineCount int, outputFile *os.File) {
	fmt.Fprintf(outputFile, "%s, %d\n", filepath.Base(path), lineCount)
}

func main() {
	start := time.Now()
	rootPtr := flag.String("root", "", "location of root folder")
	flag.Parse()
	root := *rootPtr
	if root == "" {
		panic("please define all Flags")
	}

	// Create file to store results
	outputFile, err := os.Create("results.csv")
	check(err)
	defer outputFile.Close()

	// Write headers
	fmt.Fprintf(outputFile, "File name,Numeber of lines\n")

	// Go through directory
	err = filepath.Walk(root, processTxtFiles(outputFile))
	check(err)

	fmt.Printf("Time elapsed: %s", time.Since(start))
}
