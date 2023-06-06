<svelte:head>
    <title>Run</title> 
</svelte:head>

<script lang="ts">
  import { onMount } from 'svelte';
  import { gaStatus } from '$lib/ts/stores';
  import RunGrid from '$lib/components/run/RunGrid.svelte';
  import RunGraph from '$lib/components/run/RunGraph.svelte';
  import { Individual } from '$lib/ts/ga/individual';

  onMount(async () => { // TODO: temp
    let fittestIndividual: Individual = new Individual(9, undefined);
    while (true) {
        fittestIndividual.genome = Array.from({ length: 9 }, () =>
          Math.round(Math.random()).toString()
        );
        await gaStatus.update(data => {
            fittestIndividual.fitness = data.generation * Math.random()/50;
            data.fittestIndividual = fittestIndividual;
            data.generation++;
            return data;
        });
        await new Promise(resolve => setTimeout(resolve, 500));
    }
  });
</script>

<style lang="scss">
  .visual {
    display: flex;
    align-items: center;
    margin-top: $size-070;
    column-gap: $size-070;
    justify-content: flex-start;
  }
</style>

<div class="visual">
  <RunGrid />
  <RunGraph />
</div>

