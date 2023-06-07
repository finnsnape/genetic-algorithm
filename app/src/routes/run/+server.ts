import { json } from '@sveltejs/kit';
import { promises as fs } from 'fs';
import path from 'path';

export async function POST({ request }) {
	const config: Record<string, unknown> = await request.json();
  let jsonFilePath: string = path.resolve(`./src/lib/data/run/config.json`);
  fs.writeFile(jsonFilePath, JSON.stringify(config));
	return json({ status: 200 });
}