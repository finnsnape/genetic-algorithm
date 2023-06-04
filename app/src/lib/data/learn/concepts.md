# Overview

The general idea of a genetic algorithm is to replicate the process of natural evolution, in order to solve search and optimisation problems. It does this by iterating on itself, which is shown below.

![image info](/learn/overview.png)

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Initialisation

The beginning of the process is to initialise a population which makes up the first generation of individuals, with each individual being a random guess at a solution to the problem. The size and representation of these individuals are parameters that the developer assigns.

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Evaluation

Next, it is required to evaluate how close each individual is to a desirable solution. This can be assessed quantitatively, and the resulting score is referred to as the fitness. As such, evaluation often involves developing a fitness function that can take an individual as input and return its fitness score as output.

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Selection

After this, the individuals undergo a selection process. This is designed to mimic 'survival of the fittest', and hence typically involves individuals with a better fitness having a higher change of selection.

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Crossover

Crossover is the next process, and is similar to biological reproduction. It involves parent individuals that have been chosen in the Selection stage being combined in order to produce offspring individuals. The overall goal here is to produce offspring with better fitness than their parents.

![image info](/learn/crossover.svg)

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Mutation

Mutation is usually the final operation on the individuals in one loop, and is important for preventing the individuals from converging on some local optimum. It promotes this diversity by randomly altering some of the individuals. Mutation functions typically have a mutation rate parameter which is proportional to how likely a respective individual is to be mutated.

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Termination

At this point, a new generation has been created. As such, we can start the loop again by evaluating this new population. Often, the process is only stopped once some termination condition is reached. For example, we may choose to stop the program once we have found an individual that exactly matches our solution.

## External resources
- [https://en.wikipedia.org/wiki/Genetic_algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm)


# Parameters