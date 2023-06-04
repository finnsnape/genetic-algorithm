import { Individual } from "./individual";
import type { Selector } from "./selector";
import type { Evaluator } from "./evaluator";
import type { Reproducer } from "./reproducer";
import type { Mutator } from "./mutator";
import { Population } from "./population";
import { CustomSelector } from "./selector";
import { CustomMutator } from "./mutator";
import { CustomEvaluator } from "./evaluator";
import { CustomReproducer } from "./reproducer";
import LoadCustomFunctions from "$lib/ts/load_functions";

export default class GeneticAlgorithm {
  selector: Selector;
  evaluator: Evaluator;
  reproducer: Reproducer;
  mutator: Mutator;
  population: Population;
  maxGenerations: number;
  target: string[];
  generation: number;
  solved: boolean;

  constructor(
      selectorIndex: number,
      evaluatorIndex: number,
      reproducerIndex: number,
      mutatorIndex: number,
      maxGenerations: number,
      populationSize: number,
      target: string[],
  ) {
      this.target = target;
      this.selector = new CustomSelector();
      this.selector.select = LoadCustomFunctions.getSelect(selectorIndex);
      this.evaluator = new CustomEvaluator();
      this.evaluator.evaluate = LoadCustomFunctions.getEvaluate(evaluatorIndex);
      this.reproducer = new CustomReproducer();
      this.reproducer.crossover = LoadCustomFunctions.getCrossover(reproducerIndex);
      this.mutator = new CustomMutator();
      this.mutator.mutate = LoadCustomFunctions.getMutate(mutatorIndex);
      this.population = new Population(populationSize, target);
      this.maxGenerations = maxGenerations;
      this.target = target;
      this.generation = 0;
      this.solved = false;
  }

  evaluate(): void {
    this.population.totalFitness = 0.0;
    this.population.individuals.forEach((individual, i) => {
      let fitness: number = this.evaluator.evaluate(individual, this.target);
      this.population.individuals[i].fitness = fitness;
      this.population.totalFitness += fitness;
      if (fitness > this.population.fittestIndividual.fitness) {
        this.population.fittestIndividual = this.population.individuals[i];
      }
      if (fitness == 1.0) {
        this.solved = true;
      }
    });
  }

  reproduce(): void {
    let individuals: Individual[] = [];
    while (individuals.length < this.population.size) {
      let parent1: Individual = this.selector.select(this.population);
      let parent2: Individual = this.selector.select(this.population);
      let [child1Genome, child2Genome] = this.reproducer.crossover(parent1, parent2);
      individuals.push(new Individual(0.0, child1Genome));
      individuals.push(new Individual(0.0, child2Genome));
    };
    this.population.individuals = individuals;
  }

  mutate(): void {
    this.population.individuals.forEach((individual, i) => {
      this.population.individuals[i].genome = this.mutator.mutate(individual); // this.mutator.Mutate(individual);
    });
  }

  print(): void {
    console.log(this.generation + " " + this.population.individuals.length);
    //console.log(`[${this.generation}] ${this.population.totalFitness / this.population.size}`); // 3 dp?
  }

  evolve(): void {
    // this.print();
    this.evaluate();
    if (this.solved) {
      return;
    }
    this.reproduce();
    this.mutate();
    this.generation++;
  }

  simulate(): boolean {
    this.evolve();
    if (this.solved) {
      console.log("Solution found")
      return true; // solution found
    } else if (this.generation > this.maxGenerations && this.maxGenerations != 0) {
      console.log("Maximum generations reached");
      return true; // exceeded max generations
    }
    return false;
  }
}