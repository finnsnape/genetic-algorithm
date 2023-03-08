package main

import (
	"fmt"
)

type Population struct {
	size         int
	individuals  []Individual
	totalFitness float64
}

func (population Population) print() {
	for _, individual := range population.individuals {
		fmt.Println(string(individual.genome))
	}
}

func newPopulation(size int, target []byte) Population {
	var individuals []Individual
	for i := 0; i < size; i++ {
		individual := newIndividual(len(target))
		individuals = append(individuals, individual)
	}
	return Population{
		size:        size,
		individuals: individuals,
	}
}
