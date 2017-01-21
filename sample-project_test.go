// Test code for sample-project by Fernando Nara
// 20/01/17

package main

import (
	"path/filepath"
	"testing"
)

type testdata struct {
	path  string
	count int
}

var tests = []testdata{
	{"test1.txt", 7539},
	{"test2.txt", 10570},
	{"test3.txt", 4577},
}

func TestCountLines(t *testing.T) {
	for _, pair := range tests {
		lineCount := countLines(pair.path)
		if lineCount != pair.count {
			t.Error(
				"For", filepath.Base(pair.path),
				"expected", pair.count,
				"got", lineCount,
			)
		}
	}
}
