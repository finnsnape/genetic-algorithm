<script lang="ts">
  import loader from '@monaco-editor/loader';
  import { onMount } from 'svelte';
  import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
  import { createCategory } from "$lib/ts/stores";
  import { format } from "prettier";
  import parserTypescript from "prettier/parser-typescript";
  import ts from 'typescript';

  let editor: Monaco.editor.IStandaloneCodeEditor;
  let monaco: typeof Monaco;
  let editorContainer: HTMLElement;

  onMount(async () => {
    const monacoEditor = await import('monaco-editor');
    loader.config({ monaco: monacoEditor.default });
    monaco = await loader.init();

    let configDeclaration: string = `let config = ${JSON.stringify($createCategory.config)};`

    let formattedCode: string = format(`${$createCategory.functionDeclaration}${$createCategory.currentFunction.code}}`, {
      parser: 'typescript',
      plugins: [parserTypescript],
      printWidth: 100
    });

    editor = monaco.editor.create(editorContainer, {
      value: [$createCategory.imports, '', configDeclaration, '', formattedCode].join('\n'),
      language: 'typescript',
      fontSize: 18,
      minimap: {enabled: false},
      inlineSuggest: {enabled: true},
      theme: 'vs-light',
      automaticLayout: true
    });

    function formatFunctionBody(body: string): string {
    // Remove the enclosing braces
    const bodyWithoutBraces = body.slice(1, body.length - 1);

  // Split into lines, trim each line, then join them back together
  // We keep line breaks before comments to make them more readable
  const formatted = bodyWithoutBraces
    .split('\n')
    .map(line => {
      // If it's a comment line, preserve it as is
      if(line.trim().startsWith('//')) {
        return '\n' + line.trim() + '\n';
      }
      // Otherwise, trim and keep it on the same line
      return line.trim();
    })
    .join('');

  // Remove semicolon followed by whitespace except when it's followed by a newline
  return formatted.replace(/;(?=\s(?!$))/g, ';');
}

    function cleanCode(code: string): string | null {
      console.log(code);
      const sourceFile = ts.createSourceFile('temp.ts', code, ts.ScriptTarget.Latest, true);
      let functionBody: ts.Block | null = null;

      function findFunction(node: ts.Node) {
        if (ts.isFunctionDeclaration(node) && node.body) {
          functionBody = node.body;
        } else {
          ts.forEachChild(node, findFunction);
        }
      }

      findFunction(sourceFile);

      if (functionBody) {
        const printer = ts.createPrinter();
        const newCode = printer.printNode(ts.EmitHint.Unspecified, functionBody, sourceFile);
        return newCode;
      }
      return null;
    }

    // function cleanCode(code: string) {
    //   console.log("imports", $createCategory.imports);
    //   let cleanCode: string = code.replace($createCategory.imports, "").replace(configDeclaration, "").replace($createCategory.functionDeclaration, "").replaceAll("  ", "");
    //   let finalBraceIndex: number = cleanCode.lastIndexOf("}");
    //   return cleanCode.slice(0, finalBraceIndex) + cleanCode.slice(finalBraceIndex + 1);
    // }

    editor.onDidChangeModelContent(() => {
      console.log("~~~~~~~~~~~~~~~~");
      console.log("changed");
      let code: string | null = formatFunctionBody(cleanCode(editor.getValue()));
      console.log(code);
      if (code === null) return;
      createCategory.update(currentValue => {
        console.log(currentValue.currentFunction.nickname);
        currentValue.currentFunction.code = code;
        return currentValue;
      });
    });
  });
</script>

<style lang="scss">
  .wrapper {
    align-self: flex-start;
    padding: $size-000;
    border: 1px solid #000;
    border-radius: $size-000;
  }

  .editor {
    width: 1200px;
    height: 600px;
  }
</style>

<div class="wrapper">
  <div class="editor" bind:this={editorContainer}></div>
</div>