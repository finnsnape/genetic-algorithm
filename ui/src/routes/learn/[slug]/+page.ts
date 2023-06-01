/** @type {import('./$types').PageLoad} */
import learn from "$lib/data/learn/learn.json";

export async function load({ params }) {
  let learnItem = learn[params.slug];
  return learnItem;
}