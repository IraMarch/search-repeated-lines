// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// delete comment
//delete comment
//add comment 3
package main

import (
	"bufio"
	"fmt"
	"os"
)

type countBody struct {
	count int
	filenames map[string]bool
}

func main() {
	countmap := make(map[string]*countBody)
	files := os.Args[1:]
	//fmt.Println (os.Args)
//	if len(files) == 0 {
//		countLines(os.Stdin, countmap)
//	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, countmap,len(files))
			f.Close()
		}
//	}
	for line, n := range countmap {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t", n.count, line)
			for f := range n.filenames {
				fmt.Printf("%s ", f)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, countmap map[string]*countBody, nfiles int) {
	input := bufio.NewScanner(f)
	fileinfo, err := f.Stat()
	filename := fileinfo.Name()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)

	}
	for input.Scan() {
		if _, ok  := countmap[input.Text()]; !ok {
			countmap[input.Text()] = &countBody{0, make (map[string]bool)}
		}
			if !countmap[input.Text()].filenames[filename] {
				countmap[input.Text()].filenames[filename]= true
			}
			countmap[input.Text()].count++

	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
