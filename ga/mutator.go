package main

import (
	"math/rand"
)

type Mutator interface {
	Mutate(individual Individual) []byte
}

type RandomReset struct {
	mutationRate float64
}

func (mutator RandomReset) Mutate(individual Individual) []byte {
	for i := range individual.genome {
		if rand.Float64() > mutator.mutationRate {
			continue
		}
		randomByte := randStringRunes(1)
		individual.genome[i] = randomByte[0]
	}
	return individual.genome
}
