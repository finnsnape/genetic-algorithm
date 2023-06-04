import { Individual } from "./individual";

export class Population {
  size: number;
  individuals: Individual[];
  totalFitness: number;
  fittestIndividual: Individual;

  constructor(size: number, target: string[]) {
      this.size = size;
      this.individuals = Array.from({ length: size }, () => new Individual(target.length));
      this.totalFitness = 0.0;
      this.fittestIndividual = this.individuals[0];
  }

  print(): void {
    this.individuals.forEach(individual => {
      console.log(individual.genome)
    });
  }
}