package main

import (
	"math/rand"
)

type Breeder interface {
	Breed(parent1 Individual, parent2 Individual) ([]byte, []byte)
}

type OnePointCrossover struct {
	crossoverRate float64
}

func (breeder OnePointCrossover) Breed(parent1 Individual, parent2 Individual) ([]byte, []byte) {
	if rand.Float64() > breeder.crossoverRate {
		return parent1.genome, parent2.genome
	}
	randomNumber := rand.Intn(len(parent1.genome))
	child1Genome := concatSlices(parent1.genome[0:randomNumber], parent2.genome[randomNumber:len(parent1.genome)])
	child2Genome := concatSlices(parent2.genome[0:randomNumber], parent1.genome[randomNumber:len(parent1.genome)])
	return child1Genome, child2Genome
}
