// Sample project by Fernando Nara
// 21/01/17

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
func processTxtFiles(result *os.File) filepath.WalkFunc {
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
	fileNamePtr := flag.String("output", "", "file with result data")
	rootPtr := flag.String("root", "", "location of root folder")

	flag.Parse()
	root := *rootPtr
	fileName := *fileNamePtr
	if root == "" || fileName == "" {
		panic("please define all Flags")
	}

	// Create file to store results
	outputFile, err := os.Create(fileName)
	check(err)
	defer outputFile.Close()

	// Write headers
	fmt.Fprintf(outputFile, "File name,Numeber of lines\n")

	// Go through directory
	err = filepath.Walk(root, processTxtFiles(outputFile))
	check(err)

	fmt.Printf("Time elapsed: %s", time.Since(start))
}
