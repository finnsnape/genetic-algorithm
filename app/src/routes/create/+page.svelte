<svelte:head>
  <title>Create</title> 
</svelte:head>

<script lang="ts">
  /** @type {import('./$types').PageData} */  export let data;
</script>

<style lang="scss">
.wrapper {
  display: flex;
  flex-direction: column;
}

.items {
  display: grid;
  grid-template-columns: repeat(2, max(600px));
  gap: $size-070;
}

.item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: $size-040;
  border-radius: $size-000;
  border: 1px solid #000;
  a.create {
    margin-top: $size-020;
    margin-bottom: $size-020;
  }
}
</style>

<div class="wrapper">
<h1 class="heading-1">Create</h1>
<div class="items">
  {#each data.categories as category, i}
  <div class="item">
      <h2 class="heading-2">{category.title}</h2>
      <a class="button button--secondary paragraph-2 create" href="/create/{category.route}">+ Create New</a>
      <ul>
        {#each category.functions as func, i}
          <li class="paragraph-2">{func.nickname} <a href="/create/{category.route}?func={i}" class="button button--tertiary">(Edit)</a></li>
        {/each}
      </ul>
  </div>
  {/each}
</div>
</div>