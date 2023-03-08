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
	for i := uint(0); i < selector.tournamentSize; i++ {
		randomNumber := rand.Intn(int(population.size)) // range?
		competitor := population.individuals[randomNumber]
		competitors = append(competitors, competitor)
	}
	winner := competitors[0]
	for _, competitor := range competitors {
		if competitor.fitness < winner.fitness {
			continue
		}
		winner = competitor
	}
	return winner
}

type Roulette struct{}

func (Roulette) Select(population Population) Individual {
	//if population.totalFitness == 0 {
	//	randomNumber := rand.Intn(len(population.individuals))
	//	return population.individuals[randomNumber]
	//}
	randomNumber := rand.Float64()
	probabilitySum := 0.0
	for _, individual := range population.individuals {
		probabilitySum += individual.fitness / population.totalFitness
		if randomNumber < probabilitySum {
			return individual
		}
	}
	panic("Couldn't select a fitness")
}
