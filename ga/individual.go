package main

type Individual struct {
    fitness float64
    genome []byte
}

func newIndividual(targetLength int) *Individual {
    //randomBytes := make([]byte, len(ga.target))
    randomBytes := randStringRunes(targetLength)
    //    _, err := rand.Read(randomBytes)
    //    if err != nil {
    //        panic("Couldn't initialise individual")
    //    }
    return &Individual{
        genome: randomBytes,
    }
}