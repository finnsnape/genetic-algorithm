/** @type {import('./$types').PageLoad} */

import functions from "$lib/config/functions.json";

export async function load({ fetch, params }) {
    return functions[params.slug];
}