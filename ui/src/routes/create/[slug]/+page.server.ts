import { error } from '@sveltejs/kit';
import * as prettier from "prettier";
import parserTypescript from "prettier/parser-typescript";

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
      console.log(category.functions);
    }
  }

  if (!category) throw error(404, "Create category not found");

  let categoryTitle: string = category.title;
  let functionDeclaration: string = category.functionDeclaration;
  let codeString: string = "\n// Write your code here\n";
  let nickname: string = "";
  let imports: string = category.imports;
  let newFunction: boolean = true;

  const functionIndex: string = url.searchParams.get("func");
  if (functionIndex) {
    let func = category.functions[+functionIndex];
    newFunction = true;
    codeString = func.code;
    nickname = func.nickname;
  }

  let code: string = await prettier.format(`${imports}\n\n${functionDeclaration}${codeString}}`, {
    parser: "typescript",
    plugins: [parserTypescript]
  });

  const populationFilePath = path.resolve(`src/lib/ts/ga/population.ts`);
  const populationContent: string = fs.readFileSync(populationFilePath, 'utf-8');

  const individualFilePath = path.resolve(`src/lib/ts/ga/individual.ts`);
  const individualContent: string = fs.readFileSync(individualFilePath, 'utf-8');

  return { individualContent, populationContent, categoryTitle, code, nickname, newFunction };
}