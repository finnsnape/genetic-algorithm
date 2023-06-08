<svelte:head>
    <title>Run</title> 
</svelte:head>

<script lang="ts">
  import { gaStatus } from '$lib/ts/stores';
  import { GeneticAlgorithm } from '$lib/ts/ga';
  import RunGrid from '$lib/components/run/RunGrid.svelte';
  import RunGraph from '$lib/components/run/RunGraph.svelte';

  /** @type {import('./$types').PageData} */ export let data;

  let config = JSON.stringify(data.config);
  let target: string = "1".repeat(1024);
  let populationSize: number = 100;
  let maxGenerations = 500;
  let functionIndexes: number[] = [0, 0, 0, 0];
  let done: boolean = true;
  let ga: GeneticAlgorithm;
  let elapsed: number = 0;

  async function simulateGA() {
    elapsed = 0;
    const startTime = Date.now();
    done = false;
    ga = new GeneticAlgorithm(functionIndexes[0], functionIndexes[1], functionIndexes[2], functionIndexes[3], maxGenerations, populationSize, target.split(""));
    while (!done) {
      done = ga.simulate();
      await updatePage();
      await new Promise(r => requestAnimationFrame(r)); // wait until paint has occurred (so our Grid component updates)
    }
    const endTime = Date.now();
    elapsed = (endTime - startTime)/1000;
  }

  async function updatePage() {
    await gaStatus.update(value => {
      value.fittestIndividual = ga.population.fittestIndividual;
      value.generation = ga.generation;
      return value;
    });
  }

  async function stopGA() {
    ga.solved = true;
    done = true;
  }

  async function saveConfig() {
    const response = await fetch('/run', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: config
    });
  }
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  .form {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    row-gap: $size-040;
  }

  fieldset {
    max-width: 80%;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
    flex-wrap: wrap;
    column-gap: $size-070;
    row-gap: $size-040;
    padding: $size-020;
    border-radius: $size-000;
    border: 1px solid #000;
  }

  .form-item {
    display: flex;
    flex-direction: column;
  }

  .buttons {
    display: flex;
    column-gap: $size-020;
    align-items: center;
  }

  .visual {
    display: flex;
    align-items: center;
    margin-top: $size-070;
    column-gap: $size-070;
  }

  button[disabled] {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .btn-config {
    flex: 0 0 100%;
  }
</style>

<div class="wrapper">
  <h1 class="heading-1">Run</h1>
  <div class="form">
    <fieldset>
      <legend class="paragraph-3">Functions</legend>
      {#each data.categories as category, i}
      <div class="form-item">
        <label class="paragraph-2" for="category-{i}">{category.title}</label>
        <select class="paragraph-3" name="category-{i}" bind:value={functionIndexes[i]}>
        {#each category.functions as func, j}
            <option value={j}>{func.nickname}</option>
        {/each}
        </select>
      </div>
      {/each}
    </fieldset>
    <fieldset>
      <legend class="paragraph-3">Parameters</legend>
      <div class="form-item">
        <label class="paragraph-2" for="populationSize">Population size</label>
        <input class="paragraph-3" type="number" name="populationSize" bind:value={populationSize} placeholder=100>
      </div>
      <div class="form-item">
        <label class="paragraph-2" for="maxGenerations">Max generations</label>
        <input class="paragraph-3" type="number" name="populationSize" bind:value={maxGenerations} placeholder=500>
      </div>
      <div class="form-item">
        <label class="paragraph-2" for="target">Target bit array</label>
        <input class="paragraph-3" type="text" name="target" bind:value={target} style="width: 50ch" placeholder="111111111">
      </div>
      <div class="form-item">
        <label class="paragraph-2" for="config">Config JSON object</label>
        <input class="paragraph-3" type="text" name="config" bind:value={config} style="width: 100ch">
      </div>
      <button class="button button--tertiary paragraph-3 btn-config" on:click={saveConfig}>Save Config object</button>
    </fieldset>
    <div class="buttons">
      <button class="button button--primary" on:click={simulateGA}>Start</button>
      <button class="button button--primary" on:click={stopGA} disabled={done}>Terminate</button>
      {#if elapsed !== 0}
      <p class="paragraph-2">Executed in {elapsed}s</p>
      {/if}
    </div>
  </div>
  <div class="visual">
    <RunGrid />
    <RunGraph />
  </div>
</div>