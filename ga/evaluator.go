package main

type Evaluator interface {
    Evaluate(individual Individual, target []byte) float64
}

type ByteMatch struct {}

func (ByteMatch) Evaluate(individual Individual, target []byte) float64 {
    fitness := 0
    for i, b := range target {
        if b == individual.genome[i] {
            fitness++
        }
    }
    return float64(fitness)/float64(len(target))
}
