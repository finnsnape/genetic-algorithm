import type { Individual } from "./individual";
import type { Population } from "./population";
// import * as _ from "lodash";

export interface Selector {
  select(population: Population): Individual;
}

// export class Tournament implements Selector {
//   tournamentSize: number;

//   constructor(tournamentSize: number) {
//     this.tournamentSize = tournamentSize;
//   }

//   select(population: Population): Individual {
//     const competitors: Individual[] = _.sampleSize(population.individuals, this.tournamentSize);
//     const winner = competitors.reduce((currentWinner, competitor) => competitor.fitness >= currentWinner.fitness ? competitor : currentWinner);
//     return winner;
//   }
// }

export class CustomSelector implements Selector {
  select(population: Population): Individual {
    throw new Error("Custom select not implemented");
  }
}

