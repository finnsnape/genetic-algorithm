<script lang="ts">
  import LearnMenu from "$lib/components/learn/LearnMenu.svelte";
  /** @type {import('./$types').PageData} */  export let data;
</script>

<svelte:head>
    <title>Learn</title> 
</svelte:head>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: row;
    column-gap: $size-070;
  }

  .learn-content {
    :global(h1:not(:first-child)) {
      padding-top: 2em;
      border-top: 1px solid #ccc;
    }
  }

</style>

<div class="wrapper">
  <LearnMenu></LearnMenu>
  <div class="learn-content stack-recursive-l">
    {@html data.markdown}
  </div>

</div>

