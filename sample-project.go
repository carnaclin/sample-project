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
func execute(result *os.File) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			countLines(path, result)
		}
		return nil
	}
}

// Count lines and print results for each file found
func countLines(path string, result *os.File) {
	inputFile, err := os.Open(path)
	check(err)
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Fprintf(result, "%s, %d\n", filepath.Base(path), lineCount)
}

func main() {
	start := time.Now()
	flag.Parse()
	root := flag.Arg(0)

	// Create file to store results
	outputFile, err := os.Create("result.csv")
	check(err)

	// Go through directory
	err = filepath.Walk(root, execute(outputFile))
	check(err)

	outputFile.Close()
	fmt.Printf("Time elapsed: %s", time.Since(start))
}
