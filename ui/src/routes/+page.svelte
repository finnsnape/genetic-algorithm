<style lang="scss">
    .app {
        width: 1000px;
        height: 800px;
        border: 1px solid black;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    h1 {
        margin-top: $size-010;
    }

    form {
        display: flex;
        flex-direction: column;
        margin-top: $size-040;

        label, select, input {
          text-align: center;
        }

        input {
            width: 300px;
        }

        input[type="number"], select, button {
          padding: 5px 8px;
          border-radius: 8px;
          border: 1px solid #999;
        }
    }
</style>

<div class="app">
    <h1 class="heading-2">Config</h1>

    <form class="stack">
        <label for="mutation-rate">Mutation rate (%)</label>
        <input type="number" name="mutation-rate" min=0 max=100 value=5>

        <label for="population-size">Population size</label>
        <input type="number" name="population-size" min=0 max=10000 value=100>

         <label for="max-generations">Max generations</label>
        <input type="number" name="max-generations" min=0 value=0>

        <label for="selection-mode">Selection mode</label>
        <select name="selection-mode">
            <option value="0">Roulette</option>
            <option value="1">Tournament</option>
            <option value="1">Rank</option>
        </select>

        <label for="breeding-mode">Breeding mode</label>
        <select name="breeding-mode">
            <option value="0">Single-point crossover</option>
            <option value="1">Two-point crossover</option>
            <option value="2">Uniform crossover</option>
        </select>

        <label for="crossover-rate">Crossover rate (%)</label>
        <input type="number" name="mutation-rate" min=0 max=100 value=90>

        <button>Save</button>
        <button>Run</button>
    </form>
</div>
