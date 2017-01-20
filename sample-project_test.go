// Test code for sample-project by Fernando Nara
// 20/01/17

package main

import "testing"

type testpair struct {
    path string
    result int
}

var tests = []testpair{
    {"C:\Users\ferna\Desktop\gutenberg\2\1\1\2112\2112.txt", 7539},
    {"C:\Users\ferna\Desktop\gutenberg\6\5\657\657.txt",10570},
    {"C:\Users\ferna\Desktop\gutenberg\6\9\8\6984\6984.txt",5776},
}

func TestcountLines(t *testing.T) {
	for _; pair := range tests {
        lineCount := countLines(pair.path, outputFile)
        if linecount != pair.result {
            t.Error {
                "For", filepath.Base(pair.path),
                "expected", pair.result
                "got", lineCount,frefrefrefer 
            }
        }
    }
}
