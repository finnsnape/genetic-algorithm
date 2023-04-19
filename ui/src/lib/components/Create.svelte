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
</style>

<form method="POST" class="stack">
    <h1>Create a {func.title}</h1>
    <p>{func.description}</p>
    <label for="name">Unique name:</label>
    <input name="name" type="text" placeholder={func.placeholder}>
    <label for="code">Code:</label>
    <p>func {func.name}({formattedInputs}) {formattedOutputTypes} &#123;</p>
    <div style="margin: 0 40px; display: flex; flex-direction: column">
        <textarea name="code" rows="20"></textarea>
        return {formattedOutputNames}
    </div>
    }
    <button>Save</button>
</form>