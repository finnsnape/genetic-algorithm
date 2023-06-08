import type { Individual } from "./individual";

export interface Evaluator {
  evaluate(individual: Individual, target: string[]): number;
}

export class CustomEvaluator implements Evaluator {
  evaluate(individual: Individual, target: string[]): number {
    throw new Error("Custom evaluate not implemented");
  }
}