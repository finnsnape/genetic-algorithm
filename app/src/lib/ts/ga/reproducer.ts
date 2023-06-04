import type { Individual } from "./individual";

export interface Reproducer {
  crossover(parent1: Individual, parent2: Individual): [string[], string[]];
}

export class CustomReproducer implements Reproducer {
  crossover(parent1: Individual, parent2: Individual): [string[], string[]] {
    throw new Error("Custom crossover not implemented");
  }
}

