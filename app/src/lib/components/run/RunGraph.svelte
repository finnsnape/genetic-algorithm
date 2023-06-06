<script lang="ts">
  import Chart from 'chart.js/auto';
  import { onMount } from 'svelte';
  import { gaStatus } from '$lib/ts/stores';
  // import { Individual } from '$lib/ts/ga';
  import type { GAStatus } from '$lib/ts/types';

  let chart: Chart;
  let canvas: HTMLCanvasElement;
  let gaData: GAStatus;

  onMount(async () => {
    if (typeof window === "undefined") return; // is this needed?
    chart = new Chart(canvas, {
      type: 'line',
      data: {
        labels: [],
        datasets: [{label: 'Max fitness', data: [], fill: false, borderColor: 'rgb(0,0,0)', tension: 0.15}]
      },
      options: {
        responsive: true,
        animation: {duration: 0},
        scales: {
          y: {
            beginAtZero: true,
            max: 1,
            title: {display: true,text: "Maximum fitness",font: {size: 20,family: "Source Code Pro",}}
          },
          x: {
            beginAtZero: true,
            title: {display: true,text: "Generation",font: {size: 20,family: "Source Code Pro",}}
          }
        },
        plugins: {legend: {display: false}}
      }
    });

    gaStatus.subscribe((value) => {
        gaData = value;
        updateChart();
        // update graph
    });
  });

  function updateChart() {
    if (!chart || !gaData.fittestIndividual) return;
    let highestFitness: number = gaData.fittestIndividual.fitness;
    let generation: string = gaData.generation.toString();
    chart.data.labels.push(generation);
    chart.data.datasets[0].data.push(highestFitness);
    chart.update();
  }

  function resetChart() {
    if (!chart) return;
    chart.data.labels = [];
    chart.data.datasets.forEach((dataset) => {
      dataset.data = [];
    });
    chart.update();
  }
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
    height: 500px;
  }

  canvas {
    height: 100%;
  }

  button {
    align-self: flex-end;
  }
</style>

<div class="wrapper">
  <canvas bind:this={canvas} id="chart"></canvas>
  <button class="button button--tertiary paragraph-3" on:click={resetChart}>Reset graph</button>
</div>