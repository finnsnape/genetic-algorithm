<script lang="ts">
  import loader from '@monaco-editor/loader';
  import { onMount } from 'svelte';
  import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
  import { createCategory } from "$lib/ts/stores";
  import { format } from "prettier";
  import parserTypescript from "prettier/parser-typescript";

  let editor: Monaco.editor.IStandaloneCodeEditor;
  let monaco: typeof Monaco;
  let editorContainer: HTMLElement;

  onMount(async () => {
    const monacoEditor = await import('monaco-editor');
    loader.config({ monaco: monacoEditor.default });
    monaco = await loader.init();

    let formattedCode: string = format(`${$createCategory.functionDeclaration}${$createCategory.currentFunction.code}}`, {
      parser: 'typescript',
      plugins: [parserTypescript],
      printWidth: 100
    });
    

    editor = monaco.editor.create(editorContainer, {
      value: [`${$createCategory.imports}\n`,formattedCode].join('\n'),
      language: 'typescript',
      fontSize: 18,
      minimap: {enabled: false},
      inlineSuggest: {enabled: true},
      theme: 'vs-light',
      automaticLayout: true
    });

    function cleanCode(code: string) {
      let cleanCode: string = code.replace($createCategory.imports, "").replace($createCategory.functionDeclaration, "").replaceAll("\n", "").replaceAll("  ", "");
      let finalBraceIndex: number = cleanCode.lastIndexOf("}");
      return cleanCode.slice(0, finalBraceIndex) + cleanCode.slice(finalBraceIndex + 1);
    }

    editor.onDidChangeModelContent(() => {
      let code: string = cleanCode(editor.getValue());
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