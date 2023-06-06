import { error } from '@sveltejs/kit';
import type { CreateFunction, CreateCategory } from '$lib/ts/types';
import reproducers from "$lib/data/create/reproducers.json";
import mutators from "$lib/data/create/mutators.json";
import selectors from "$lib/data/create/selectors.json";
import evaluators from "$lib/data/create/evaluators.json";

/** @type {import('./$types').PageServerLoad} */
export async function load({ url, params }) {
  let categories = [reproducers, mutators, selectors, evaluators];
  const category = categories.find((cat) => cat.route === params.slug);
  if (!category) throw error(404, "Create category not found");

  let currentFunction: CreateFunction = {
    index: undefined,
    nickname: "",
    code: "\n// Write your code here\n"
  };

  const functionIndex: string = url.searchParams.get("func");
  if (functionIndex) {
    currentFunction.index = +functionIndex;
    currentFunction.nickname = category.functions[+functionIndex].nickname;
    currentFunction.code = category.functions[+functionIndex].code;
  };

  let createCategory: CreateCategory = {
    title: category.title,
    route: category.route,
    imports: category.imports,
    functionDeclaration: category.functionDeclaration,
    currentFunction: currentFunction
  };

  return { createCategory };
}