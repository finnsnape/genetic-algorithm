package main

import(
    cryptoRand "crypto/rand"
    "encoding/binary"
    "math/rand"
)

func init() {
    var b [8]byte
    _, err := cryptoRand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
    //tests()
    populationSize := uint(100)
    target := []byte("Hello.")
    //target := []byte("The quick brown fox jumps over the lazy dog.")
    maxGenerations := uint(50)
    mutator := RandomReset{mutationRate: 10}
    selector := Tournament{tournamentSize: 2}
    breeder := OnePointCrossover{}
    evaluator := ByteMatch{}
    ga := newGeneticAlgorithm(selector, evaluator, breeder, mutator, maxGenerations, target, populationSize)
    ga.simulate()
}