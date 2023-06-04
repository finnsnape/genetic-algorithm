import type { Individual } from "./individual";
import type { Population } from "./population";
// import * as _ from "lodash";

export interface Selector {
  select(population: Population): Individual;
}

export class CustomSelector implements Selector {
  select(population: Population): Individual {
    throw new Error("Custom select not implemented");
  }
}

