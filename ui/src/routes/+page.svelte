<script lang="ts">
    import Info from "$lib/components/Info.svelte";
import type { Image } from "$lib/ts/types";

    const description: string[] = ["Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.", "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."]

    const image: Image = {alt: "", src: "https://media.geeksforgeeks.org/wp-content/uploads/20190620121215/singleCrossover-2.png"}
</script>
<style lang="scss">
    .app {
        display: flex;
        flex-direction: column;
        align-items: center;
        position: relative;
        min-height: 100vh;
    }
</style>

<div class="app">
  <Info></Info>
</div>
