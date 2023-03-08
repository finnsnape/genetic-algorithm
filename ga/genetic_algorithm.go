package main

import (
    "fmt"
)

type GeneticAlgorithm struct {
    selector Selector
    evaluator Evaluator
    breeder Breeder
    mutator Mutator

    population *Population
    maxGenerations uint
    target []byte
    generation uint
}

func newGeneticAlgorithm(selector Selector, evaluator Evaluator, breeder Breeder, mutator Mutator, maxGenerations uint, target []byte, populationSize uint) *GeneticAlgorithm {
    return &GeneticAlgorithm{
        selector: selector,
        evaluator: evaluator,
        breeder: breeder,
        mutator: mutator,
        population: newPopulation(populationSize, target),
        maxGenerations: maxGenerations,
        target: target,
        generation: 0,
    }
}

func (ga *GeneticAlgorithm) evaluate() {
    ga.population.totalFitness = 0.0
    for i:=uint(0); i<ga.population.size; i++ {
        individual := ga.population.individuals[i]
        individual.fitness = ga.evaluator.Evaluate(*individual, ga.target)
        ga.population.totalFitness += individual.fitness
        if individual.fitness > ga.population.fittestIndividual.fitness {
            fmt.Println("fit",individual.fitness)
            ga.population.fittestIndividual = individual
        } else if individual.fitness < ga.population.leastFitIndividual.fitness {
            fmt.Println("unfit",individual.fitness)
            ga.population.leastFitIndividual = individual
        }
    }
}

func (ga *GeneticAlgorithm) breed() {
    var individuals []*Individual
    for uint(len(individuals)) < ga.population.size {
        parent1 := ga.selector.Select(*ga.population)
        parent2 := ga.selector.Select(*ga.population)
        child := &Individual{
            genome: ga.breeder.Breed(parent1, parent2),
        }
        child.fitness = ga.evaluator.Evaluate(*child, ga.target)
        individuals = append(individuals, child)
    }
    ga.population.individuals = individuals
}

func (ga *GeneticAlgorithm) mutate() {
    for _, individual := range ga.population.individuals {
        individual.genome = ga.mutator.Mutate(*individual)
    }
}

func (ga *GeneticAlgorithm) print() {
    fmt.Println("Generation: ", ga.generation)
    fmt.Println("Average fitness: ", ga.population.totalFitness/float64(ga.population.size))
    fmt.Println("Best fitness: ", ga.population.fittestIndividual.fitness)
    fmt.Println("Best solution: ", string(ga.population.fittestIndividual.genome))
    fmt.Println()
}

func (ga *GeneticAlgorithm) evolve() {
    ga.evaluate()
    ga.print()
    ga.breed()
    ga.mutate()
    ga.generation++
}

func (ga *GeneticAlgorithm) simulate() {
    for {
        ga.evolve()
        if ga.population.fittestIndividual.fitness == 1.0 {
            break // solution found
        } else if ga.generation > ga.maxGenerations && ga.maxGenerations != 0 {
            break // exceeded max generations
        }
    }
}
