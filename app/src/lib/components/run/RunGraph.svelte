<script lang="ts">
  import Chart from 'chart.js/auto';
  import { onMount } from 'svelte';
  import { gaStatus } from '$lib/ts/stores';

  let chart: Chart;
  let canvas: HTMLCanvasElement;

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
      if (!chart || !value.fittestIndividual || !value.generation) return;
      if (value.generation == 1) resetChart(false);
      chart.data.labels.push(value.generation);
      chart.data.datasets[0].data.push(value.fittestIndividual.fitness);
      chart.update();
    });
  });

  function resetChart(update: boolean) {
    if (!chart) return;
    chart.data.labels = [];
    chart.data.datasets.forEach((dataset) => {
      dataset.data = [];
    });
    if (update) chart.update();
  }
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
    height: 700px;
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
  <button class="button button--tertiary paragraph-3" on:click={() => resetChart(true)}>Reset graph</button>
</div>