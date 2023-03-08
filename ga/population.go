package main

import (
	"sync"
)

type Population struct {
	size               uint
	individuals        []*Individual
	totalFitness       float64
	fittestIndividual  *Individual // maybe just keep sorted instead?
	leastFitIndividual *Individual
	waitGroup          *sync.WaitGroup
}

func newPopulation(size uint, target []byte) *Population {
	var individuals []*Individual
	for i := uint(0); i < size; i++ {
		individual := newIndividual(len(target))
		individuals = append(individuals, individual)
	}
	return &Population{
		size:               size,
		individuals:        individuals,
		fittestIndividual:  &Individual{fitness: 0.0},
		leastFitIndividual: &Individual{fitness: 1.0},
	}
}
