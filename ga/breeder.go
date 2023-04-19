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

type MultiPointCrossover struct {
	crossoverRate float64
	n             int
}

func (breeder MultiPointCrossover) Breed(parent1 Individual, parent2 Individual) ([]byte, []byte) {
	// FIXME
	return parent1.genome, parent2.genome
}

type UniformCrossover struct {
	crossoverRate float64
}

func (breeder UniformCrossover) Breed(parent1 Individual, parent2 Individual) ([]byte, []byte) {
	if rand.Float64() > breeder.crossoverRate {
		return parent1.genome, parent2.genome
	}
	child1Genome := make([]byte, 0, len(parent1.genome))
	child2Genome := make([]byte, 0, len(parent1.genome))
	for i := range parent1.genome {
		if rand.Float64() > 0.5 {
			child1Genome = append(child1Genome, parent1.genome[i])
			child2Genome = append(child2Genome, parent2.genome[i])
		} else {
			child1Genome = append(child1Genome, parent2.genome[i])
			child2Genome = append(child2Genome, parent1.genome[i])
		}
	}
	return child1Genome, child2Genome
}

type CustomBreeder struct {
	code           string
	needsCompiling bool
	function       func(Individual, Individual) ([]byte, []byte)
}

func (breeder *CustomBreeder) Breed(parent1 Individual, parent2 Individual) ([]byte, []byte) {
	if breeder.function == nil || breeder.needsCompiling {
		v := codeToFunction(breeder.code, "Breeder")
		breeder.function = v.Interface().(func(Individual, Individual) ([]byte, []byte))
		breeder.needsCompiling = false
	}

	return breeder.function(parent1, parent2)
}
