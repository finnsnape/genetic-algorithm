<script lang="ts">    
  export let fittestGenome: string[];
</script>

<style lang="scss">
  .image {
    display: grid;
    grid-template-columns: repeat(32, 1fr);
    height: 500px;
    width: 500px;
    border: 1px solid #000;
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

<div class="image">
  {#each fittestGenome as char, _}
    <div class="pixel {char === '0' ? 'white' : 'black'}"></div>
  {/each}
</div>