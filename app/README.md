https://www.npmjs.com/package/@typescript/sandbox

mention they can use _.

what if we let them define a custom `config` object in Run, and have it imported by default in each function? then they can do stuff like `export const config = {mutationRate: 0.002}` and then `import config from "...";` 

maybe we can export it to a json file? will we be able to import that in monaco?

or we can export it to a ts file?