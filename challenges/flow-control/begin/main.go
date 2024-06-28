// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("recovered from panic:",err)
		}
	}()

	// use os package to read the file specified as a command line argument
	fname := os.Args[1]
	data, err := os.ReadFile(fname)
	if err != nil{
		panic(fmt.Errorf("error reading file: %s",err))
	}

	// convert the bytes to a string
	fcontent := string(data)

	// initialize a map to store the counts
	tracker := map[string]int{
		"letters":0,
		"symbols":0,
		"numbers":0,
		"words":0,
	}

	// use the standard library utility package that can help us split the string into words
	words := strings.Split(fcontent, " ")

	// capture the length of the words slice
	tracker["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words{
		for _, char := range word{
			if unicode.IsLetter(char) {
				tracker["letters"]++
			} else if unicode.IsNumber(char) {
					tracker["numbers"]++
				} else {
					tracker["symbols"]++
				}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(tracker)
}
