package main

import (
    "math/rand"
)

type Mutator interface {
    Mutate(individual Individual) []byte
}

type RandomReset struct {
    mutationRate uint
}

func (mutator RandomReset) Mutate(individual Individual) []byte {
    randomNumber := rand.Intn(100)
    if uint(randomNumber) > mutator.mutationRate {
        return individual.genome
    }
    randomByte := randStringRunes(1)
    randomNumber = rand.Intn(len(individual.genome))
    individual.genome[randomNumber] = randomByte[0]
    return individual.genome
}