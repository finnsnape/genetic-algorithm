import fs from 'fs';
import path from 'path';
import { marked } from 'marked';

async function renderMarkdown(rawMarkdown: string): Promise<string> {
  const renderer = new marked.Renderer();

  renderer.paragraph = (text) => {
    return `<p class="paragraph-2">${text}</p>`;
  };

  renderer.link = (href, _, text) => {
    return `<a class="paragraph-3" href="${href}">${text}</a>`;
  };

  renderer.heading = (text, level) => {
    return `<h${level} class="heading-${level}" style="margin-bottom: 0">${text}</h${level}>`;
  };

  marked.use({ renderer });

  return marked(rawMarkdown, {mangle: false, headerIds: false});
}

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
  const slug: string = params.slug;
  const filePath = path.resolve(`src/lib/data/learn/${slug}.md`);

  const rawMarkdown = fs.readFileSync(filePath, 'utf-8');

  // TODO: CATCH!
  const markdown = await renderMarkdown(rawMarkdown);
  return { markdown };
}