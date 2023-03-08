package main

import (
	"fmt"
	"sync"
)

type Population struct {
	size         uint
	individuals  []Individual
	totalFitness float64
	waitGroup    *sync.WaitGroup
}

func (population Population) print() {
	for _, individual := range population.individuals {
		fmt.Println(string(individual.genome))
	}
}

func newPopulation(size uint, target []byte) Population {
	var individuals []Individual
	for i := uint(0); i < size; i++ {
		individual := newIndividual(len(target))
		individuals = append(individuals, individual)
	}
	return Population{
		size:        size,
		individuals: individuals,
	}
}
