import { readFile } from 'fs/promises';
import learn from '$lib/data/learn/learn.json';

export async function getLearnMenu() {
  let learnMenu = [];
  for (const category of learn) {
    let subHeadings: string[] = await getHeadings(category.file);
    let menuItem = {[category.title]: subHeadings}
    learnMenu.push(menuItem);
  }
  return learnMenu;
}


async function getHeadings(file: string): Promise<string[]> {
  const data: string = await readFile(`src/lib/data/learn/${file}`, 'utf8');
  const matches = data.match(/^# (?!#)(.+)/gm);
  if (!matches) {
      return [];
  }
  const headings: string[] = matches.map(match => match.replace('# ', ''));

  return headings;
}