<script lang="ts">
  import type { Image } from "$lib/ts/types";
  import SlideMenu from "./SlideMenu.svelte";

  export let title: string;
  export let description: string[];
  export let image: Image;
</script>

<style lang="scss">
  .wrapper {
    margin: $size-040 0;
  }
</style>

<div class="wrapper paragraph-2 ch-l stack">
  <h1 class="heading-2">{title}</h1>

  {#each description as paragraph}
  <p>{paragraph}</p>
  {/each}
  
  <img src={image.src} alt={image.alt}/>
  
</div>