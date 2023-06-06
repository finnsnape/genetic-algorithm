import { error } from '@sveltejs/kit';
import * as prettier from "prettier";
import parserTypescript from "prettier/parser-typescript";
import type { CreateFunction } from "$lib/ts/types";

import fs from 'fs';
import path from 'path';

import reproducers from "$lib/data/create/reproducers.json";
import mutators from "$lib/data/create/mutators.json";
import selectors from "$lib/data/create/selectors.json";
import evaluators from "$lib/data/create/evaluators.json";


/** @type {import('./$types').PageLoad} */
export async function load({ url, params }) {
  let categories = [reproducers, mutators, selectors, evaluators];
  let category;
  for (const cat of categories) {
    if (cat.route === params.slug) {
      category = cat;
    }
  }

  if (!category) throw error(404, "Create category not found");

  let createFunc: CreateFunction = {
    categoryTitle: category.title,
    nickname: "",
    code: "",
    new: true,
    path: `${category.route}.json`,
  };

  let functionDeclaration: string = category.functionDeclaration;
  let codeString: string = "\n// Write your code here\n";
  let imports: string = category.imports;

  const functionIndex: string = url.searchParams.get("func");
  if (functionIndex) {
    let func = category.functions[+functionIndex];
    createFunc.new = false;
    codeString = func.code;
    createFunc.nickname = func.nickname;
  }

  createFunc.code = await prettier.format(`${imports}\n\n${functionDeclaration}${codeString}}`, {
    parser: "typescript",
    plugins: [parserTypescript]
  });

  const populationFilePath = path.resolve(`src/lib/ts/ga/population.ts`);
  const populationContent: string = fs.readFileSync(populationFilePath, 'utf-8');
  // TODO: make this one function
  const individualFilePath = path.resolve(`src/lib/ts/ga/individual.ts`);
  const individualContent: string = fs.readFileSync(individualFilePath, 'utf-8');

  return { individualContent, populationContent, createFunc };
}

