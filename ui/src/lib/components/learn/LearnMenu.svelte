<script lang="ts">
    import { page } from '$app/stores'; 
    import learn from "$lib/data/learn/learn.json";

    let menuItems = new Map<string, string>();
    for (let key in learn) {
      let value: string = learn[key];
      menuItems.set(key, value.title);
    }
</script>

<nav>
  <ol class="stack">
    <li class="paragraph-2"><a class="button button--tertiary" href="/learn" class:active={$page.url.pathname === "/learn"}>Welcome</a></li>
    {#each [...menuItems] as [key, value]}
    <li class="paragraph-2"><a class="button button--tertiary" href="/learn/{key}" class:active={$page.url.pathname.split("/")[2] === key}>{value}</a></li>
    {/each}
  </ol>
</nav>
