import { Utils } from "./utils";

export class Individual {
  fitness: number;
  genome: string[];

  constructor(targetLength: number, genome?: string[]) {
      if (typeof genome == 'undefined') {
        this.genome = Utils.generateRandomChars(targetLength);
      } else {
        this.genome = genome;
      }
      this.fitness = 0.0;
  }
}