import { promises as fs } from 'fs';
import { readFile } from 'fs/promises';
import path from 'path';
import { marked } from 'marked';
import { getLearnMenu } from "$lib/ts/learn_menu";

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
  const slug: string = params.slug;
  const filePath = path.resolve(`src/lib/data/learn/${slug}.md`);

  const rawMarkdown = await readFile(filePath, 'utf-8');

  // TODO: CATCH!
  const markdown = await renderMarkdown(rawMarkdown);

  const learnMenu = await getLearnMenu();

  return { markdown, learnMenu };
}

async function renderMarkdown(rawMarkdown: string): Promise<string> {
  const renderer = new marked.Renderer();

  renderer.paragraph = (text) => {
    return `<p class="paragraph-2">${text}</p>`;
  };

  renderer.link = (href, _, text) => {
    return `<a class="paragraph-3" href="${href}">${text}</a>`;
  };

  let headingCounter: number = 0;
  renderer.heading = (text, level) => {
    let headingRender: string = `<h${level} class="heading-${level}" style="margin-bottom: 0" id="heading-${headingCounter}">${text}</h${level}>`;
    if (level == 1) headingCounter++;
    return headingRender;
  };

  marked.use({ renderer });

  return marked(rawMarkdown, {mangle: false, headerIds: false});
}