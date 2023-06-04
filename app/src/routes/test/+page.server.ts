import { Individual } from "$lib/ts/ga/individual";
import type { Reproducer } from "$lib/ts/ga";
import { CustomReproducer } from "$lib/ts/ga/reproducer";
import LoadCustomFunctions from "$lib/ts/load_functions";

/** @type {import('./$types').PageLoad} */
export async function load() {
  let individual1: Individual = new Individual(5, ["a", "b", "c", "d", "e"]);
  let individual2: Individual = new Individual(5, ["0", "1", "2", "3", "4"]);
  
  let reproducer: Reproducer = new CustomReproducer();
  reproducer.crossover = LoadCustomFunctions.getCrossover(0);
  
  let [child1, child2] = reproducer.crossover(individual1, individual2);
  
  console.log(child1);
  console.log(child2);
}
