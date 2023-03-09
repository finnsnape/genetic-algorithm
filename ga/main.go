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
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	//tests()
	populationSize := 512
	maxGenerations := 0
	mutator := RandomReset{mutationRate: 0.002}
	//selector := Tournament{tournamentSize: 8}
	selector := Rank{selectionPressure: 1.8}
	//breeder := MultiPointCrossover{crossoverRate: 0.9, n: 1}
	breeder := UniformCrossover{crossoverRate: 0.9}
	evaluator := ByteMatch{}
	ga := newGeneticAlgorithm(selector, evaluator, breeder, mutator, maxGenerations, populationSize)
	ga.simulate()
}
