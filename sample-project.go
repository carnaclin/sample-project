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
	// Implement word counter with line counter
	inputFile, err := os.Open(path)
	check(err)
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	//fileScanner.Split(bufio.ScanLines) //already by default
	lineCount, wordCount := 0, 0
	var textLine string
	for fileScanner.Scan() {
		lineCount++
		textLine = fileScanner.Text()
		wordCount += countWords(textLine, word)
	}
	//fmt.Printf("number of lines %d\n", lineCount)
	//fmt.Printf("number of words %d\n", wordCount)
	return lineCount, wordCount
}

// Count words
func countWords(textLine string, word string) int {
	// Here i get a text line as string to work with. I need to tokenize it correctly
	// to get all the words and do map["word"]++ to count numbers of "words"

	// src is the input to tokenize.
	src := []byte(textLine)

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to san yield the token sequence found in the input.
	wordCount := 0
	for {
		_, tok, lit := s.Scan()
		if strings.ToLower(lit) == word {
			wordCount++
			//fmt.Println("\t\t\tMatch")
			//fmt.Printf("\t\t\twordCount = %d\n", wordCount)
		}
		if tok == token.EOF {
			break
		}
		//fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
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
