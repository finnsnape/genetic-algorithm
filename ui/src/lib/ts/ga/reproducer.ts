import type { Individual } from "./individual";

export interface Reproducer {
  crossover(parent1: Individual, parent2: Individual): [string[], string[]];
}

// export class OnePointCrossover implements Reproducer {
//   crossoverRate: number;

//   constructor(crossoverRate: number) {
//     this.crossoverRate = crossoverRate;
//   }

//   crossover(parent1: Individual, parent2: Individual): [string[], string[]] {
//     if (Math.random() > this.crossoverRate) {
//       return [parent1.genome, parent2.genome];
//     }
//     const randomIndex: number = Math.floor(Math.random() * parent1.genome.length);
//     const child1Genome: string[] = parent1.genome.slice(0, randomIndex).concat(parent2.genome.slice(randomIndex));
//     const child2Genome: string[] = parent2.genome.slice(0, randomIndex).concat(parent1.genome.slice(randomIndex));
//     return [child1Genome, child2Genome];
//   }
// }

export class CustomReproducer implements Reproducer {
  crossover(parent1: Individual, parent2: Individual): [string[], string[]] {
    throw new Error("Custom crossover not implemented");
  }
}

