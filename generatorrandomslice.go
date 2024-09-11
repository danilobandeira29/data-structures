package main

import (
	"math/rand"
)

var random_until = 10000

func GenerateRandomSlice(size int) []int {
	result := []int{}
	for i := 0; i < size; i++ {
		n := rand.Intn(random_until)
		for Contains(result, n) {
			n = rand.Intn(random_until)
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
	result := []int{}
	for i := size; i >= 0; i-- {
		result = append(result, i)
	}
	return result
}
