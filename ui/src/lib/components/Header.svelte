<script lang="ts">
  import { page } from '$app/stores'; 

  let nav = [
    {
      "href": "/learn",
      "title": "Learn",
    },
    {
      "href": "/create",
      "title": "Create",
    },
    {
      "href": "/run",
      "title": "Run",
    },
  ];
</script>

<style lang="scss">
    nav {
      display: flex;
      justify-content: space-evenly;
      padding: $size-010;
      margin-bottom: $size-060;
    }
</style>

<nav>
  {#each nav as navItem}
  <a class="button button--tertiary paragraph-2" class:active={$page.url.pathname.split("/")[1] === navItem.href.split("/")[1]} href={navItem.href}>{navItem.title}</a>
  {/each}
</nav>