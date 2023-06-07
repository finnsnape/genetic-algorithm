import type { Individual } from "./individual";

export interface Mutator {
  mutate(individual: Individual): string[];
}

export class CustomMutator implements Mutator {
  mutate(individual: Individual): string[] {
    throw new Error("Custom mutate not implemented");
  }
}