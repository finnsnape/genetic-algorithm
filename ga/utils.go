package main

import (
	"math/rand"
)

func randStringRunes(n int) []byte {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ. ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return []byte(string(b))
}

func randFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Safely concatenates two slices
func concatSlices[A any](slice1 []A, slice2 []A) []A {
	length := len(slice1) + len(slice2)
	slice := make([]A, 0, length)
	slice = append(slice, slice1...)
	slice = append(slice, slice2...)
	return slice
}
