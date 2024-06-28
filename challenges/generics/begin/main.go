// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func printAny[T string|int|bool](t T) {
	fmt.Println(t)
}

// Part 2 sum function refactoring

type numeric interface{
	constraints.Integer | constraints.Float
}
// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)
	
	// call sum function
	//fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	fmt.Println(sum(1, 2, 3))
}
