import type { Individual } from "./individual";

export interface Evaluator {
  evaluate(individual: Individual, target: string[]): number;
}

export class CharMatch implements Evaluator {
  evaluate(individual: Individual, target: string[]): number {
    let fitness: number = 0.0;
    target.forEach((char, i) => {
      if (char == individual.genome[i]) fitness++;
    });
    return fitness/target.length;
  }
}

export class CustomEvaluator implements Evaluator {
  evaluate(individual: Individual, target: string[]): number {
    throw new Error("Custom evaluate not implemented");
  }
}