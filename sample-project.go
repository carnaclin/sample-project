package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execute(result *os.File) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			countLines(path, result)
		}
		return nil
	}
}

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
	outputFile, err := os.Create("result.csv")
	check(err)
	err = filepath.Walk(root, execute(outputFile))
	check(err)
	outputFile.Close()

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s", elapsed)
}
