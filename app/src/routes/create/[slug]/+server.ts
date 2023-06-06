import type { CreateFunction } from "$lib/ts/types";
import { json } from '@sveltejs/kit';
import fs from 'fs';
import path from 'path';

/** @type {import('./$types').RequestHandler} */
export async function POST({ request }) {
  const { createFunc } = await request.json();
  console.log("body", createFunc["path"]);
  const urlParams = new URLSearchParams(request.url);
  let jsonFilePath: string = path.resolve(`./src/lib/data/create/${createFunc["path"]}`);
  let currentJson = JSON.parse(fs.readFileSync(jsonFilePath, 'utf-8'));
  if (!createFunc.new) {
    const functionIndex: string = urlParams.get("func");
    currentJson["functions"][+functionIndex]["code"] = createFunc["code"];
  } else {
    currentJson.functions.push({
      nickname: createFunc["nickname"],
      code: createFunc["code"]
    });
  }
  fs.writeFile(jsonFilePath, currentJson, (err) => {
  });
  return {
    "status": 200,
  };
}