package main

import (
	"flag"
	"fmt"

	stack "github.com/danilobandeira29/data-structs/00_stack"
)

func ConvertDecimalToBinary(n int) []byte {
	stack := stack.NewStack[int]()
	for n != 0 {
		rest := n % 2
		stack.Push(rest)
		n = n / 2
	}
	var numbers []byte
	for !stack.IsEmpty() {
		numbers = append(numbers, byte(stack.Pop()))
	}
	return numbers
}

func main() {
	n := flag.Int("n", 0, "number to be converted to binary")
	flag.Parse()
	fmt.Printf("number %d will be converted to binary\n", *n)
	result := ConvertDecimalToBinary(*n)
	fmt.Println(result)
}
