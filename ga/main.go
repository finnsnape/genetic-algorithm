package main

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func init() {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		panic("Cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:]))) // deprecated
}

func main() {
	//tests()
	maxGenerations := 5
	adder := AddX{x: 4}
	multiplier := MultiplyByX{x: 2}
	ga := newGeneticAlgorithm(adder, multiplier, maxGenerations)
	ga.simulate()
}
