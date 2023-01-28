package main

import (
    "crypto/rand"
    "fmt"
)


type Individual struct {
    fitness uint
    genome []byte
}

type Population struct {
    size uint
    individuals []Individual
    totalFitness uint
    fittestIndividual *Individual
}

type GeneticAlgorithm struct {
    population Population
    maxGenerations uint
    target []byte
    generation uint
}

func (ga *GeneticAlgorithm) initialiseIndividual() Individual {
    randomBytes := make([]byte, len(ga.target))
    _, err := rand.Read(randomBytes)
    if err != nil {
        panic("Couldn't initialise individual")
    }
    individual := Individual{}
    individual.genome = randomBytes
    return individual
}

func (ga *GeneticAlgorithm) initialisePopulation(populationSize uint) {
    ga.population = Population{} // ?
    ga.population.size = populationSize
    for i:=uint(0); i<ga.population.size; i++ {
        individual := ga.initialiseIndividual()
        ga.population.individuals = append(ga.population.individuals, individual)
    }
}

func (ga *GeneticAlgorithm) evaluateIndividual(individual Individual) uint {
    fitness := uint(0)
    for i, b := range ga.target {
        if b != individual.genome[i] {
            fitness++
        }
    }
    return fitness
}

func (ga *GeneticAlgorithm) evaluatePopulation() {
    ga.population.totalFitness = 0
    ga.population.fittestIndividual = &Individual{^uint(0), nil}
    for _, individual := range ga.population.individuals {
        individual.fitness = ga.evaluateIndividual(individual)
        ga.population.totalFitness += individual.fitness
        if individual.fitness < ga.population.fittestIndividual.fitness {
            ga.population.fittestIndividual = &individual
        }
    }
}

func (ga *GeneticAlgorithm) selectParents() []Individual {

}

func (ga *GeneticAlgorithm) breedPopulation() {
    parents := ga.selectParents()
}

func (ga *GeneticAlgorithm) mutateIndividual(individual Individual) Individual {

}

func (ga *GeneticAlgorithm) mutatePopulation() {

}

func (ga *GeneticAlgorithm) print() {
    fmt.Println("Generation: ", ga.generation)
    fmt.Println("Average fitness: ", ga.population.totalFitness/ga.population.size)
    fmt.Println("Best fitness: ", ga.population.fittestIndividual.fitness)
    fmt.Println("Best solution: ", ga.population.fittestIndividual.genome)
    fmt.Println()
}

func (ga *GeneticAlgorithm) evolve() {
    ga.evaluatePopulation()
    ga.print()
    ga.breedPopulation()
    ga.mutatePopulation()
    ga.generation++
}

func main() {
    populationSize := uint(10)
    target := []byte("The quick brown fox jumps over the lazy dog")
    ga := GeneticAlgorithm{}
    ga.initialisePopulation(populationSize)
    ga.target = target
    ga.maxGenerations = 0 // no max
    for {
        ga.evolve()
        if ga.population.fittestIndividual.fitness == 0 {
            break // solution found
        } else if ga.generation > ga.maxGenerations && ga.maxGenerations != 0 {
            break // exceeded max generations
        }
    }
}