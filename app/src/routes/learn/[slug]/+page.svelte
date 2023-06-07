<svelte:head>
  <title>Learn</title>
</svelte:head>

<script lang="ts">
  /** @type {import('./$types').PageData} */  export let data;
  import { page } from '$app/stores'; 
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
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

  nav {
    position: fixed;
    width: 20%;
  }
</style>

<div class="wrapper">
  <nav>
    <ol class="paragraph-2">
      {#each data.learnMenu as menuItem}
      {#each Object.entries(menuItem) as [title, subHeadings]}
      <li>
        <a class="button button--tertiary" class:active={$page.url.pathname.split("/")[2] === title.toLowerCase()} href="/learn/{title.toLowerCase()}">{title}</a>
        {#if subHeadings}
        <ol class="paragraph-3" type="a">
          {#each subHeadings as subHeading, i}
          <li><a class="button button--tertiary" href="/learn/{title.toLowerCase()}#heading-{i}">{subHeading}</a></li>
          {/each}
        </ol>
        {/if}
      </li>
      {/each}
      {/each}
    </ol>
  </nav>
  <div class="learn-content stack-recursive-l">
    {@html data.markdown}
  </div>
</div>