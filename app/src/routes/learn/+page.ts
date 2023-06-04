import learn from "$lib/data/learn/learn.json";
import { redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */
export async function load() {
  let firstPath = learn[0].file.replace(".md", "");
  throw redirect(302, `/learn/${firstPath}`);
}