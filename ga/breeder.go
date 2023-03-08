package main

import (
    "math/rand"
)

type Breeder interface {
    Breed(parent1 Individual, parent2 Individual) []byte
}

type OnePointCrossover struct {}

func (OnePointCrossover) Breed(parent1 Individual, parent2 Individual) []byte {
    randomNumber := rand.Intn(len(parent1.genome))
    childGenome := append(parent1.genome[0:randomNumber], parent2.genome[randomNumber:len(parent1.genome)]...)
    return childGenome
}