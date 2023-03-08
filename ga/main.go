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
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	//tests()
	populationSize := uint(512)
	target := []byte("The quick brown fox jumps over the lazy dog.")
	maxGenerations := uint(0)
	mutator := RandomReset{mutationRate: 0.002}
	//selector := Tournament{tournamentSize: 8}
	selector := Roulette{}
	breeder := OnePointCrossover{crossoverRate: 0.9}
	evaluator := ByteMatch{}
	ga := newGeneticAlgorithm(selector, evaluator, breeder, mutator, maxGenerations, target, populationSize)
	ga.simulate()
}
