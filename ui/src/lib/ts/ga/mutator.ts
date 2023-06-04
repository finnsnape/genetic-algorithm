import type { Individual } from "./individual";
import { Utils } from "./utils";

export interface Mutator {
  mutate(individual: Individual): string[];
}

export class RandomReset implements Mutator {
  mutationRate: number;

  constructor(mutationRate: number) {
    this.mutationRate = mutationRate;
  }

  mutate(individual: Individual): string[] {
    [...individual.genome].forEach((char, i) => {
      if (Math.random() > this.mutationRate) {
        return;
      }
      individual.genome[i] = Utils.generateRandomChars(1)[0];
    });
    return individual.genome;
  }
}

export class CustomMutator implements Mutator {
  mutate(individual: Individual): string[] {
    throw new Error("Custom mutate not implemented");
  }
}