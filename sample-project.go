// Sample project by Fernando Nara
// 21/01/17

package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Looks for .txt extensions in all folders
func processTxtFiles(outputFile *os.File, word string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			lineCount, wordCount := countLines(path, word)
			writeToFile(path, lineCount, wordCount, outputFile)
		}
		return nil
	}
}

// Count lines and print results for each file found
func countLines(path string, word string) (int, int) {
	inputFile, err := os.Open(path)
	check(err)
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	lineCount, wordCount := 0, 0
	var textLine string
	for fileScanner.Scan() {
		lineCount++
		textLine = fileScanner.Text()
		wordCount += countWords(textLine, word)
	}
	return lineCount, wordCount
}

// Count words occurrences one line at a time
func countWords(textLine string, word string) int {
	src := []byte(textLine)
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)
	wordCount := 0
	for {
		_, tok, lit := s.Scan()
		if strings.ToLower(lit) == word {
			wordCount++
		}
		if tok == token.EOF {
			break
		}
	}
	return wordCount
}

// Write results to output file
func writeToFile(path string, lineCount int, wordCount int, outputFile *os.File) {
	fmt.Fprintf(outputFile, "%s, %d, %d\n", filepath.Base(path), lineCount, wordCount)
}

// Write headers
func writeHeaders(outputFile *os.File, word string) {
	fmt.Fprintf(outputFile, "File name,Number of lines,Occurrences of \"%s\"\n", word)
}

func main() {
	start := time.Now()
	rootPtr := flag.String("root", "", "location of root folder")
	wordPtr := flag.String("word", "", "word to be searched")
	flag.Parse()
	root := *rootPtr
	word := *wordPtr
	if (root == "") || (word == "") {
		panic("please define all Flags")
	}

	// Create file to store results
	outputFile, err := os.Create("results.csv")
	check(err)
	defer outputFile.Close()

	writeHeaders(outputFile, word)

	// Go through directory
	err = filepath.Walk(root, processTxtFiles(outputFile, word))
	check(err)

	fmt.Printf("Time elapsed: %s", time.Since(start))
}
