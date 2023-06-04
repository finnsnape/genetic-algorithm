<script lang="ts">
  import loader from '@monaco-editor/loader';
  import { onMount } from 'svelte';
  import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
  export let code: string;
  export let populationContent: string;
  export let individualContent: string;
  export let newFunction: boolean;

  import 
  
  // let editor;
  // let editorContainer;

  let editor: Monaco.editor.IStandaloneCodeEditor;
  let monaco: typeof Monaco;
  let editorContainer: HTMLElement;

  onMount(async () => {
    const monacoEditor = await import('monaco-editor');
    loader.config({ monaco: monacoEditor.default });

    monaco = await loader.init();

    monaco.languages.typescript.typescriptDefaults.addExtraLib(individualContent, '/ga/individual.ts');
    monaco.languages.typescript.typescriptDefaults.addExtraLib(populationContent, '/ga/population.ts');

    monaco.languages.typescript.typescriptDefaults.setCompilerOptions({
      target: monaco.languages.typescript.ScriptTarget.ESNext,
      allowNonTsExtensions: true,
      moduleResolution: monaco.languages.typescript.ModuleResolutionKind.NodeJs,
      module: monaco.languages.typescript.ModuleKind.ESNext,
      noEmit: true,
      typeRoots: ["/src/ts"]
    });


    // Your monaco instance is ready, let's display some code!
    const editor = monaco.editor.create(editorContainer, {
      value: [
        code,
      ].join('\n'),
      language: 'typescript',
      fontSize: 18,
      minimap: {enabled: false},
      inlineSuggest: {enabled: true},
      theme: 'vs-light',
      automaticLayout: true
    });
  });
  
  // onMount(() => {
  //   editor = monaco.editor.create(editorContainer, {
  //     value: [
  //       `function ${functionDeclaration}`,
  //       code,
  //       '}'
  //     ].join('\n'),
  //     language: 'typescript',
  //     fontSize: 18,
  //     minimap: {enabled: false},
  //     inlineSuggest: {enabled: true},
  //     theme: 'vs-light',
  //     automaticLayout: true,
  //   });
  //   //console.log(editor.getValue());
  // });

  function saveFunction() {
    const editorContent = editor.getValue();
    console.log(editorContent);
  }

  
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  .editor {
    align-self: flex-start;
    padding: $size-000;
    border: 1px solid #000;
    border-radius: $size-000;
  }
</style>

<div class="wrapper stack">
  <div class="editor">
    <div bind:this={editorContainer} style="width:1000px;height:600px;"></div>
  </div>
  <button class="button button--primary" on:click={saveFunction}>Save</button>
</div>
