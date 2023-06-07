import reproducers from "$lib/data/create/reproducers.json";
import mutators from "$lib/data/create/mutators.json";
import selectors from "$lib/data/create/selectors.json";
import evaluators from "$lib/data/create/evaluators.json";
import config from "$lib/data/run/config.json";

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
    let categories = [reproducers, mutators, selectors, evaluators];
    return {categories, config};
}