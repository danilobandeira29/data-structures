package main

import (
	"math/rand"
)

var randomUntil = 10000

func GenerateRandomSlice(size int) []int {
	var result []int
	for i := 0; i < size; i++ {
		n := rand.Intn(randomUntil)
		for Contains(result, n) {
			n = rand.Intn(randomUntil)
		}
		result = append(result, n)
	}
	return result
}

func Contains(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func GenerateDecreasingStartFrom(size int) []int {
	var result []int
	for i := size; i >= 0; i-- {
		result = append(result, i)
	}
	return result
}
