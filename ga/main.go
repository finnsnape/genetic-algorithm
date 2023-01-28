package main

import (
    "crypto/rand"
    "fmt"
)


type Genome struct {
    fitness uint
    bytes []byte
}

type Population struct {
    size uint
    genomes []Genome
    totalFitness uint
    fittestGenome *Genome
}

type GeneticAlgorithm struct {
    population Population
    maxGenerations uint
    target []byte
    generation uint
}

func (ga *GeneticAlgorithm) initialiseGenome() Genome {
    bytes := make([]byte, len(ga.target))
    _, err := rand.Read(bytes)
    if err != nil {
        panic("Couldn't initialise genome")
    }
    genome := Genome{}
    genome.bytes = bytes
    return genome
}

func (ga *GeneticAlgorithm) initialisePopulation(populationSize uint) {
    ga.population = Population{} // ?
    ga.population.size = populationSize
    for i:=uint(0); i<ga.population.size; i++ {
        genome := ga.initialiseGenome()
        ga.population.genomes = append(ga.population.genomes, genome)
    }
}

func (ga *GeneticAlgorithm) evaluateGenome(genome Genome) uint {
    fitness := uint(0)
    for i, b := range ga.target {
        if b != genome.bytes[i] {
            fitness++
        }
    }
    return fitness
}

func (ga *GeneticAlgorithm) evaluatePopulation() {
    ga.population.totalFitness = 0
    ga.population.fittestGenome = &Genome{^uint(0), nil}
    for _, genome := range ga.population.genomes {
        genome.fitness = ga.evaluateGenome(genome)
        ga.population.totalFitness += genome.fitness
        if genome.fitness < ga.population.fittestGenome.fitness {
            ga.population.fittestGenome = &genome
        }
    }
}

func (ga *GeneticAlgorithm) selectParents() []Genome {

}

func (ga *GeneticAlgorithm) matePopulation() {
    parents := ga.selectParents()
}

func (ga *GeneticAlgorithm) mutateGenome(genome Genome) Genome {

}

func (ga *GeneticAlgorithm) mutatePopulation() {

}

func (ga *GeneticAlgorithm) print() {
    fmt.Println("Generation: ", ga.generation)
    fmt.Println("Average fitness: ", ga.population.totalFitness/ga.population.size)
    fmt.Println("Best fitness: ", ga.population.fittestGenome.fitness)
    fmt.Println("Best solution: ", ga.population.fittestGenome.bytes)
    fmt.Println()
}

func (ga *GeneticAlgorithm) evolve() {
    ga.evaluatePopulation()
    ga.print()
    ga.matePopulation()
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
        if ga.population.fittestGenome.fitness == 0 {
            break // solution found
        } else if ga.generation > ga.maxGenerations && ga.maxGenerations != 0 {
            break // exceeded max generations
        }
    }
}