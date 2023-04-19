package main

type Evaluator interface {
	Evaluate(individual Individual, target []byte) float64
}

type ByteMatch struct{}

func (ByteMatch) Evaluate(individual Individual, target []byte) float64 {
	fitness := 0
	for i, b := range target {
		if b == individual.genome[i] {
			fitness++
		}
	}
	return float64(fitness) / float64(len(target))
}

type CustomEvaluator struct {
	code           string
	needsCompiling bool
	function       func(Individual, []byte) float64
}

func (evaluator *CustomEvaluator) Evaluate(individual Individual, target []byte) float64 {
	if evaluator.function == nil || evaluator.needsCompiling {
		v := codeToFunction(evaluator.code, "Evaluator")
		evaluator.function = v.Interface().(func(individual Individual, target []byte) float64)
		evaluator.needsCompiling = false
	}

	return evaluator.function(individual, target)
}
