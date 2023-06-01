<script lang="ts">
    /** @type {import('./$types').PageData} */  export let data;
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
  }

  .items {
    display: flex;
    flex-direction: column;
    row-gap: $size-040;
  }
</style>

<div class="wrapper">
  <h1 class="heading-1">Create</h1>
  <div class="items">
    {#each Object.entries(data) as [name, func]}
    <div>
        <h2 class="heading-2">{func.title}s</h2>
        <ul>
            <li class="paragraph-2">N/A</li>
        </ul>
        <a class="button button--tertiary paragraph-2" href="/create/{name}">Create</a>
    </div>
{/each}
  </div>
</div>
