package main

import (
	"fmt"
)

type GeneticAlgorithm struct {
	adder       Adder
	multiplier  Multiplier
	maxGenerations int
	generation     int
}

func newGeneticAlgorithm(adder Adder, multiplier Multiplier, maxGenerations int) *GeneticAlgorithm {
	return &GeneticAlgorithm{
		adder:       adder,
		multiplier:      multiplier,
		maxGenerations: maxGenerations,
		generation:     0,
	}
}

func (ga *GeneticAlgorithm) evolve() {
	added := ga.adder.Add(5)
	multiplied := ga.multiplier.Multiply(4)
	fmt.Print("Added", added)
	fmt.Println("Multiplied", multiplied)
	fmt.Println("======")
	ga.generation++
}

func (ga *GeneticAlgorithm) simulate() {
	for {
		ga.evolve()
		if ga.generation > ga.maxGenerations && ga.maxGenerations != 0 {
			fmt.Println("Maximum generations reached")
			break // exceeded max generations
		}
	}
}
