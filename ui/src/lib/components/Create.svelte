<script lang="ts">
    import type { Function } from "../ts/types";

    export let func : Function;

    const formattedInputs = (function() {
        return func.inputs.map(x => x.name + " " + x.type).join(", ");
    })();

    const formattedOutputTypes = (function() {
        if (func.outputs.length == 1) {
            return func.outputs[0].type;
        }
        let outputTypes : string[] = func.outputs.map(x => x.type)
        return "(" + outputTypes.join(", ") + ")";
    })();

    const formattedOutputNames = (function() {
        return func.outputs.map(x => x.name).join(", ");
    })();
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

    button {
      align-self: flex-start;
    }
</style>

<form method="POST" class="stack">
    <h1 class="heading-1">Create a {func.title}</h1>
    <p class="paragraph-2">{func.description}</p>
    <label class="paragraph-2" for="name">Unique name:</label>
    <input name="name" type="text" placeholder={func.placeholder}>
    <label class="paragraph-2" for="code">Code:</label>
    <p class="paragraph-2">func {func.name}({formattedInputs}) {formattedOutputTypes} &#123;</p>
    <div class="paragraph-2" style="margin: 0 40px; display: flex; flex-direction: column">
        <textarea name="code" rows="15"></textarea>
        return {formattedOutputNames}
    </div>
    <p class="paragraph-2">}</p>
    <button class="button button--primary">Save</button>
</form>