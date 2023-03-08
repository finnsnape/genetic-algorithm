package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"os"
)

func randStringRunes(n int) []byte {
	var letterRunes = []rune("01")
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

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}

func printImage(genome []byte, size int) {
	for i := 0; i < size*size; i++ {
		char := "⬜"
		if genome[i] == []byte("1")[0] {
			char = "⬛"
		}
		fmt.Print(char)
		if (i+1)%size == 0 {
			fmt.Print("\n")
		}
	}
}
