package main

import (
	"flag"
	"fmt"
	"strings"

	stack "github.com/danilobandeira29/data-structs/00_stack"
)

func Is(s string) bool {
	stack := stack.NewStack[string]()
	for _, l := range s {
		stack.Push(string(l))
	}
	var newString string
	for !stack.IsEmpty() {
		newString += stack.Pop()
	}
	return newString == s
}

func main() {
	phrase := flag.String("p", "", "phrase to verify if is palindrome or not")
	flag.Parse()
	phraseWithoutSpace := strings.Replace(*phrase, " ", "", -1)
	fmt.Println(Is(strings.ToLower(phraseWithoutSpace)))
}
