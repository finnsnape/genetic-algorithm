/** @type {import('./$types').PageLoad} */

import reproducers from "$lib/data/create/reproducers.json";
import mutators from "$lib/data/create/mutators.json";
import selectors from "$lib/data/create/selectors.json";
import evaluators from "$lib/data/create/evaluators.json";

export async function load({ fetch, params }) {
  let categories = [reproducers, mutators, selectors, evaluators];
    return {categories};
}