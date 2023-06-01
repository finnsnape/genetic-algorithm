<script lang="ts">
    /** @type {import('./$types').PageData} */  export let data;
</script>

<style lang="scss">
    .result {
      margin: $size-070 auto 0 auto;
      width: 500px;
      height: 500px;
      border: 1px solid black;
    }

    form {
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: space-evenly;
    }

    div {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }

    button {
      align-self: flex-end;
    }
</style>

<h1 class="heading-1">Run</h1>
<form method="POST">
    {#each Object.entries(data) as [name, func]}
        <div>
            <label class="paragraph-2">{func.title}</label>
            <select>
                <option>N/A></option>
            </select>
        </div>
    {/each}
    <button class="button button--primary">Run</button>
</form>

<div class="result">

</div>