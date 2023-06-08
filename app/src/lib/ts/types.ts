import type { Individual } from "./ga";

export interface GAStatus {
  fittestIndividual: Individual | undefined,
  generation: number
}

export interface CreateCategory {
  title: string,
  route: string,
  imports: string,
  functionDeclaration: string,
  currentFunction: CreateFunction
  config: Record<string, unknown>,
  importContent: Map<string, string>
}

export interface CreateFunction {
  index: number | undefined,
  nickname: string,
  code: string
}