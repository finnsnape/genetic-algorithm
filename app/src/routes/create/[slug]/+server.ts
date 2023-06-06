import { json } from '@sveltejs/kit';
import type { CreateCategory, CreateFunction } from '$lib/ts/types';
import { promises as fs } from 'fs';
import { readFile } from 'fs/promises';
import path from 'path';

export async function POST({ request }) {
	const createCategory: CreateCategory = await request.json();
  //console.log(createCategory);
  let jsonFilePath: string = path.resolve(`./src/lib/data/create/${createCategory.route}.json`);
  let currentJSON = JSON.parse(await readFile(jsonFilePath, 'utf-8'));
  const currentFunction: CreateFunction = createCategory.currentFunction;
  if (currentFunction.index !== undefined) {
    currentJSON["functions"][currentFunction.index]["code"] = currentFunction.code;
    currentJSON["functions"][currentFunction.index]["nickname"] = currentFunction.nickname;
  } else {
    currentJSON["functions"].push({
      nickname: currentFunction.nickname,
      code: currentFunction.code
    });
  }
  
  fs.writeFile(jsonFilePath, JSON.stringify(currentJSON));

	return json({ status: 200 });
}