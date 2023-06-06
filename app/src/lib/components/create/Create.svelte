<script lang="ts">
    import type { CreateFunction } from '$lib/ts/types';
    import { page } from '$app/stores';

    export let createFunc: CreateFunction;
    export let populationContent: string;
    export let individualContent: string;

    import loader from '@monaco-editor/loader';
    import { onMount } from 'svelte';
    import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
    export const prerender = true;
//    import evaluators from "$lib/data/create/evaluators.json";
//    console.log(evaluators);

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

      editor = monaco.editor.create(editorContainer, {
        value: [
          createFunc.code,
        ].join('\n'),
        language: 'typescript',
        fontSize: 18,
        minimap: {enabled: false},
        inlineSuggest: {enabled: true},
        theme: 'vs-light',
        automaticLayout: true
      });
    });

    async function doPost(createFunc: CreateFunction) {
		  const res = await fetch($page.url.pathname, {
			  method: 'POST',
			  body: JSON.stringify({
				  createFunc,
			  })
		});

		await res.json()
	}

    async function saveFunction() {
      createFunc.code = editor.getValue();
      await doPost(createFunc);
    }

    // const formattedInputs = (function() {
    //     return func.inputs.map(x => x.name + " " + x.type).join(", ");
    // })();

    // const formattedOutputTypes = (function() {
    //     if (func.outputs.length == 1) {
    //         return func.outputs[0].type;
    //     }
    //     let outputTypes : string[] = func.outputs.map(x => x.type)
    //     return "(" + outputTypes.join(", ") + ")";
    // })();

    // const formattedOutputNames = (function() {
    //     return func.outputs.map(x => x.name).join(", ");
    // })();
</script>

<style lang="scss">
    form {
      display: flex;
      flex-direction: column;
      label {
        font-weight: 700;
      }
    }

    input {
      align-self: flex-start;
    }

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
<form method="POST">
  <h1 class="heading-1">Create {createFunc.categoryTitle}</h1>
  <label class="paragraph-2" for="name">Unique name:</label>
  <input name="name" type="text" placeholder="Custom{createFunc.categoryTitle}1" value={createFunc.nickname}>
  <label class="paragraph-2" for="code">Code:</label>
  <div class="wrapper stack">
    <div class="editor">
      <div bind:this={editorContainer} style="width:1000px;height:600px;"></div>
    </div>
    <button class="button button--primary" on:click|preventDefault={saveFunction}>Save</button>
  </div>
</form>
