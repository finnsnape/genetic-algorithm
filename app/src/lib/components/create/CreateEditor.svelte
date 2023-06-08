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

    $createCategory.importContent.forEach((key, value) =>{
      monaco.languages.typescript.typescriptDefaults.addExtraLib(value, key);
    });

    monaco.languages.typescript.typescriptDefaults.setCompilerOptions({
      target: monaco.languages.typescript.ScriptTarget.ESNext,
      allowNonTsExtensions: true,
      moduleResolution: monaco.languages.typescript.ModuleResolutionKind.NodeJs,
      module: monaco.languages.typescript.ModuleKind.ESNext,
      noEmit: true,
      typeRoots: ["/src/ts"]
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
      const bodyWithoutBraces = body.slice(1, body.length - 1);

      const formatted = bodyWithoutBraces.split('\n').map(line => {
        if(line.trim().startsWith('//')) {
          return '\n' + line.trim() + '\n';
        }
        return line.trim();
      }).join('');
      return formatted.replace(/;(?=\s(?!$))/g, ';');
    } // source: SO / OAI

    // source: SO / OAI
    function cleanCode(code: string): string | null {
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

    editor.onDidChangeModelContent(() => {
      let code: string | null = formatFunctionBody(cleanCode(editor.getValue()));
      if (code === null) return;
      createCategory.update(currentValue => {
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