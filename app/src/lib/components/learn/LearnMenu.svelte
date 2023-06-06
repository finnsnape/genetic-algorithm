<script lang="ts">
    import { page } from '$app/stores'; 
    import learn from "$lib/data/learn/learn.json";

    let menuItems = new Map<string, string>();
    learn.forEach(item => {
      menuItems.set(item.title, item.file.replace(".md", ""));
    });
    // create sub menu based on h1s
</script>
<nav>
  <ol class="stack">
    {#each [...menuItems] as [title, file]}
    <li class="paragraph-2"><a class="button button--tertiary" href="/learn/{file}" class:active={$page.url.pathname.split("/")[2] === file}>{title}</a></li>
    {/each}
  </ol>
</nav>