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
    flex-direction: column;
    column-gap: $size-070;
    position: relative;
  }

  .learn-content {
  width: 80%;
  align-self: flex-end;
    :global(h1:not(:first-child)) {
      padding-top: 2em;
      border-top: 1px solid #ccc;
    }
  }

  .menu {
    position: fixed;
    width: 20%;
  }

</style>

<div class="wrapper">
<div class="menu">
  <LearnMenu></LearnMenu>
</div>
  <div class="learn-content stack-recursive-l">
    {@html data.markdown}
  </div>

</div>

