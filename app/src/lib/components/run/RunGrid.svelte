<script lang="ts">
  import { onMount } from 'svelte';
  import { gaStatus } from '$lib/ts/stores';
  // import { Individual } from '$lib/ts/ga';
  import type { GAStatus } from '$lib/ts/types';

  // let data: GAStatus;
  // onMount(async () => {
  //   gaStatus.subscribe((value) => { // do we need to do this?
  //       data = value;
  //   });
  // });
</script>

<style lang="scss">
  .wrapper {
    height: 700px;
    width: 700px;
    border: 1px solid #000;
  }

  .grid {
    display: grid;
    height: 100%;
    width: 100%;
  }

  .pixel {
    border: 1px solid #333;
    &.white {
      background-color: #fff;
    }

    &.black {
      background-color: #000;
    }
  }
</style>

<div class="wrapper">
  {#if $gaStatus && $gaStatus.fittestIndividual}
  <div class="grid" style="grid-template-columns: repeat({Math.sqrt($gaStatus.fittestIndividual.genome.length)}, 1fr);">
    {#each $gaStatus.fittestIndividual.genome as bit, _}
    <div class="pixel {bit === '0' ? 'white' : 'black'}"></div>
    {/each}
  </div>
  {/if}
</div>