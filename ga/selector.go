package main

import (
    "math/rand"
)

type Selector interface {
    Select(population Population) Individual
}

type Tournament struct {
    tournamentSize uint
}
func (selector Tournament) Select(population Population) Individual {
    var competitors []Individual
    for i:=uint(0); i<selector.tournamentSize; i++ {
        randomNumber := rand.Intn(int(population.size))
        competitor := population.individuals[randomNumber]
        competitors = append(competitors, *competitor)
    }
    winner := competitors[0]
    for _, competitor := range competitors {
        if competitor.fitness <= winner.fitness {
            continue
        }
        winner = competitor
    }
    return winner
}

type Roulette struct {}

func (Roulette) Select(population Population) Individual {
    randomNumber := rand.Float64()
    probSum := 0.0
    for _, individual := range population.individuals {
        probSum += individual.fitness/population.totalFitness
        if randomNumber < probSum {
            return *individual
        }
    }
    panic("Couldn't select a fitness")
}