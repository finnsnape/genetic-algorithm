<script lang="ts">
  import Chart from 'chart.js/auto';
  import { onMount } from 'svelte';
  export let highestFitness: number;
  export let generation: number;
  
  let lastHighestFitness: number;
  let lastGeneration: number;

  // let currentDataSet: number = -1;
  // let currentColour: string;

  let chart: Chart;
  let ctx;

  // let dataSetSchema = {
    
  // };

  let data = {
    labels: [],
    datasets: [{
      label: 'Max fitness',
      data: [],
      fill: false,
      borderColor: 'rgb(0,0,0)',
      tension: 0.1
    }]
  };


  // function random_rgb() {
  //     var o = Math.round, r = Math.random, s = 255;
  //     return `rgb(${o(r()*s)}, ${o(r()*s)}, ${o(r()*s)})`;
  // }

  function resetChart() {
    if (!chart) return;
    data.labels = [];
    data.datasets.forEach((dataset) => {
      dataset.data = [];
    });
    chart.update();
  }

  $: {
    if (highestFitness != lastHighestFitness && generation != lastGeneration) { // prevent infinite loop by ensuring these have changed
      if (generation == 1) {
        resetChart();
        // if (chart) { // clear previous chart
        //   data.labels = [];
        //   data.datasets[0].data = [];
        //   chart.update();
        // }
      }
      data.labels.push(generation.toString());
      data.datasets[0].data.push(highestFitness);
      if (chart) chart.update();
      lastHighestFitness = highestFitness;
      lastGeneration = generation;
    }
  }


  onMount(async () => {
    if (typeof window !== "undefined") {
      chart = new Chart(ctx, {
        type: 'line',
        data,
        options: {
          responsive: true,
          animation: {
            duration: 0
          },
          scales: {
            y: {
              beginAtZero: true,
              max: 1,
              title: {
                display: true,
                text: "Maximum fitness",
                font: {
                  size: 20,
                  family: "Source Code Pro",
                }
              }
            },
            x: {
              beginAtZero: true,
              title: {
                display: true,
                text: "Generation",
                font: {
                  size: 20,
                  family: "Source Code Pro",
                }
              }
            }
          },
          plugins: {
              legend: {
                display: false
              }
            }
        }
      });
    }
  });
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
  <canvas bind:this={ctx} id="chart"></canvas>
  <button class="button button--tertiary paragraph-3" on:click={resetChart}>Reset graph</button>
</div>