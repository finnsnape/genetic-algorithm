<script lang="ts">
    /** @type {import('./$types').PageData} */  export let data;
    import Chart from '$lib/components/run/Chart.svelte';
    import Image from '$lib/components/run/Image.svelte';
    import { GeneticAlgorithm } from '$lib/ts/ga';

    let target: string[] = new Array(1024).fill("1");
    let populationSize: number = 96;
    let maxGenerations = 1000;
    let mutatorIndex: number = 0;
    let selectorIndex: number = 0;
    let reproducerIndex: number = 0;
    let evaluatorIndex: number = 0;
    let fittestGenome: string[] = new Array(target.length).fill("0");

    let generation: number;
    let highestFitness: number;

    async function simulateGA() {
      let ga: GeneticAlgorithm = new GeneticAlgorithm(selectorIndex, evaluatorIndex, reproducerIndex, mutatorIndex, maxGenerations, populationSize, target);
      let done: boolean = false;
      while (!done) {
        done = ga.simulate();
        updatePage(ga);
        await new Promise(r => requestAnimationFrame(r)); // wait until paint has occurred (so our Image component updates)
      }
    }

    function updatePage(ga: GeneticAlgorithm) {
      fittestGenome = ga.population.fittestIndividual.genome;
      highestFitness = ga.population.fittestIndividual.fitness;
      generation = ga.generation;
      // update fitness graphs etc
    }
</script>

<svelte:head>
    <title>Run</title> 
</svelte:head>


<style lang="scss">
    form {
      display: flex;
      flex-direction: row;
      align-items: flex-start;
      justify-content: flex-start;
      column-gap: $size-070;
    }

    .form-item {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }

    button {
      align-self: flex-end;
    }

    .visual {
      display: flex;
      align-items: center;
      margin-top: $size-070;
      column-gap: $size-070;
      justify-content: flex-start;
    }

    .graph {
      width: 800px;
      height: 500px;
      border: 1px solid black;
      border-radius: $size-000;
    }


    .console {
      margin-top: $size-070;
    }

    textarea {
      width: 100%;
    }
</style>

<h1 class="heading-1">Run</h1>
<form method="POST">
    {#each data.categories as category, i}
        <div class="form-item">
            <label class="paragraph-2">{category.title}</label>
            <select>
            {#each category.functions as func, i}
                <option value={i}>{func.nickname}</option>
            {/each}
            </select>
        </div>
    {/each}
    <button class="button button--primary" on:click|preventDefault={simulateGA}>Run</button>
</form>

<div class="visual">
  <Image {fittestGenome}></Image>
  <Chart {highestFitness} {generation}></Chart>
</div>

<details class="console paragraph-2">
  <summary>Console</summary>
  <textarea name="" id="" rows="10"></textarea>
</details>