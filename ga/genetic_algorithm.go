package main

import (
	"fmt"
	"sync"
)

type GeneticAlgorithm struct {
	selector       Selector
	evaluator      Evaluator
	breeder        Breeder
	mutator        Mutator
	population     Population
	maxGenerations int
	target         []byte
	generation     int
	solved         bool
	waitGroup      *sync.WaitGroup
}

func newGeneticAlgorithm(selector Selector, evaluator Evaluator, breeder Breeder, mutator Mutator, maxGenerations int, populationSize int) *GeneticAlgorithm {
	target := loadTarget()
	return &GeneticAlgorithm{
		selector:       selector,
		evaluator:      evaluator,
		breeder:        breeder,
		mutator:        mutator,
		population:     newPopulation(populationSize, target),
		maxGenerations: maxGenerations,
		target:         target,
		generation:     0,
		solved:         false,
	}
}

func loadTarget() []byte {
	img, _ := getImageFromFilePath("target1.png")
	bounds := img.Bounds()
	length := bounds.Dx() * bounds.Dx()
	target := make([]byte, 0, length)
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dx(); j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			val := []byte("0")
			if r == 0 && b == 0 && g == 0 {
				val = []byte("1")
			}
			target = append(target, val...)
		}
	}
	return target
}

func (ga *GeneticAlgorithm) evaluate() {
	fittestIndividual := ga.population.individuals[0]
	ga.population.totalFitness = 0.0
	for i, individual := range ga.population.individuals {
		fitness := ga.evaluator.Evaluate(individual, ga.target)
		ga.population.individuals[i].fitness = fitness
		ga.population.totalFitness += fitness
		if fitness > fittestIndividual.fitness {
			fittestIndividual = ga.population.individuals[i]
		}
		if fitness == 1.0 {
			ga.solved = true
		}
	}
	printImage(fittestIndividual.genome, 32)
}

func (ga *GeneticAlgorithm) breed() {
	var individuals []Individual
	for len(individuals) < ga.population.size {
		parent1 := ga.selector.Select(ga.population)
		parent2 := ga.selector.Select(ga.population)
		child1Genome, child2Genome := ga.breeder.Breed(parent1, parent2)
		child1, child2 := Individual{genome: child1Genome, fitness: 0.0}, Individual{genome: child2Genome, fitness: 0.0}
		individuals = append(individuals, child1)
		individuals = append(individuals, child2)
	}
	ga.population.individuals = individuals
}

func (ga *GeneticAlgorithm) mutate() {
	for i, individual := range ga.population.individuals {
		ga.population.individuals[i].genome = ga.mutator.Mutate(individual)
	}
}

func (ga *GeneticAlgorithm) print() {
	fmt.Printf("[%d] %.3f", ga.generation, ga.population.totalFitness/float64(ga.population.size))
	fmt.Println()
}

func (ga *GeneticAlgorithm) evolve() {
	ga.evaluate()
	if ga.solved {
		return
	}
	ga.breed()
	ga.mutate()
	ga.generation++
}

func (ga *GeneticAlgorithm) simulate() {
	for {
		ga.evolve()
		if ga.solved {
			break // solution found
		} else if ga.generation > ga.maxGenerations && ga.maxGenerations != 0 {
			fmt.Println("Maximum generations reached")
			break // exceeded max generations
		}
	}
}
