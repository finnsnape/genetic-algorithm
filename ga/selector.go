package main

import (
	"math/rand"
	"sort"
)

type Selector interface {
	Select(population Population) Individual
}

type Tournament struct {
	tournamentSize int
}

func (selector Tournament) Select(population Population) Individual {
	var competitors []Individual
	for i := 0; i < selector.tournamentSize; i++ {
		randomNumber := rand.Intn(population.size) // range?
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
	panic("Couldn't select an individual")
}

type Rank struct {
	selectionPressure float64
}

func (selector Rank) Select(population Population) Individual {
	sort.Slice(population.individuals, func(i, j int) bool {
		return population.individuals[i].fitness > population.individuals[j].fitness
	})
	randomNumber := rand.Float64()
	probabilitySum := 0.0
	for i, individual := range population.individuals {
		probabilitySum += (1 / float64(population.size)) * (selector.selectionPressure - (2*selector.selectionPressure-2)*float64(i)/(float64(population.size)-1))
		if randomNumber < probabilitySum {
			return individual
		}
	}
	panic("Couldn't select an individual")
}
