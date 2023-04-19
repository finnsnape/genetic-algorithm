<script lang="ts">
    /** @type {import('./$types').PageData} */  export let data;
</script>

<style lang="scss">
    .result {
      margin: $size-030 auto;
      width: 500px;
      height: 500px;
      border: 1px solid black;
    }
</style>

<form method="POST" class="stack">
    {#each Object.entries(data) as [name, func]}
        <div>
            <label>{func.title}</label>
            <select>
                <option>N/A></option>
            </select>
        </div>
    {/each}
    <button>Run</button>
</form>

<div class="result">

</div>