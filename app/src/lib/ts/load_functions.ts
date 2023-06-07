import * as _ from "lodash";

import mutators from "$lib/data/create/mutators.json";
import reproducers from "$lib/data/create/reproducers.json";
import evaluators from "$lib/data/create/evaluators.json";
import selectors from "$lib/data/create/selectors.json";
import config from "$lib/data/run/config.json";
import { transpileModule, ModuleKind } from 'typescript';

import type { Mutator } from "$lib/ts/ga/mutator";
import type { Reproducer } from "$lib/ts/ga/reproducer";
import type { Evaluator } from "$lib/ts/ga/evaluator";
import type { Selector } from "$lib/ts/ga/selector";

type MutatorFunction = Mutator['mutate'];
type ReproducerFunction = Reproducer['crossover'];
type EvaluatorFunction = Evaluator['evaluate'];
type SelectorFunction = Selector['select'];

export default class LoadCustomFunctions {
  // gonna need a lot of try/catch here!

  static transpile(code: string): string { // convert TS to JS
    const transpiled = transpileModule(code, {
      compilerOptions: { module: ModuleKind.CommonJS }
    });
    return transpiled.outputText;
  }

  static getMutate(index: number): MutatorFunction {
    let code: string = this.transpile(mutators.functions[index].code);
    let customMutator = new Function('config', '_', 'individual', code).bind(null, config, _) as MutatorFunction;
    return customMutator;
  }

  static getCrossover(index: number): ReproducerFunction {
    let code: string = this.transpile(reproducers.functions[index].code);
    let customReproducer = new Function('config', '_', 'parent1', 'parent2', code).bind(null, config, _) as ReproducerFunction;
    return customReproducer;
  }
  
  static getEvaluate(index: number): EvaluatorFunction {
    let code: string = this.transpile(evaluators.functions[index].code);
    let customEvaluator = new Function('config', '_', 'individual', 'target', code).bind(null, config, _) as EvaluatorFunction;
    return customEvaluator;
  }

  static getSelect(index: number): SelectorFunction {
    let code: string = this.transpile(selectors.functions[index].code);
    let customSelector = new Function('config', '_', 'population', code).bind(null, config, _) as SelectorFunction;
    return customSelector;
  }
}