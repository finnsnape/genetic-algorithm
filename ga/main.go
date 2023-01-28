package main

import (
    "crypto/rand"
    "fmt"
)


type Genome struct {
    fitness uint
    bits []byte
}

type Population struct {
    size uint
    genomes []Genome
    totalFitness uint
    fittestGenome *Genome
}

type GeneticAlgorithm struct {
    population Population
    target []byte
    generation uint
}

func (ga *GeneticAlgorithm) initialiseGenome() Genome {
    bits := make([]byte, len(ga.target))
    _, err := rand.Read(bits)
    if err != nil {
        panic("Couldn't initialise genome")
    }
    genome := Genome{}
    genome.bits = bits
    return genome
}

func (ga *GeneticAlgorithm) initialisePopulation(size uint) {
    ga.population = Population{} // ?
    ga.population.size = size
    for i:=uint(0); i<ga.population.size; i++ {
        genome := ga.initialiseGenome()
        ga.population.genomes = append(ga.population.genomes, genome)
    }
    ga.evaluatePopulation()
}

func (ga *GeneticAlgorithm) evaluateGenome(genome Genome) uint {
    fitness := uint(0)
    for i, b := range ga.target {
        if b != genome.bits[i] {
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

func (ga *GeneticAlgorithm) mateParents() {
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
    fmt.Println("Best solution: ", ga.population.fittestGenome.bits)
    fmt.Println()
}

func (ga *GeneticAlgorithm) evolve() {
    ga.print()
    ga.evaluatePopulation()
    ga.mateParents()
    ga.mutatePopulation()
    ga.generation++
}



func main() {
}