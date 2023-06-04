<script lang="ts">
    import CreateEditor from '$lib/components/create/CreateEditor.svelte';

    export let categoryTitle: string;
    export let nickname: string;
    export let code: string;
    export let populationContent: string;
    export let individualContent: string;
    export let newFunction: boolean;

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
</style>

<form method="POST" class="stack">
    <h1 class="heading-1">Create {categoryTitle}</h1>
    <label class="paragraph-2" for="name">Unique name:</label>
    <input name="name" type="text" placeholder="Custom{categoryTitle}1" value={nickname}>
    <label class="paragraph-2" for="code">Code:</label>
    <CreateEditor {code} {populationContent} {individualContent} {newFunction}></CreateEditor>
</form>