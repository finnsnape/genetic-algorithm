<script lang="ts">
    import LearnMenu from "$lib/components/learn/LearnMenu.svelte";
    import learn from "$lib/data/learn/learn.json";
    console.log(learn);
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: row;
    column-gap: $size-070;
  }
</style>

<div class="wrapper">
  <LearnMenu></LearnMenu>
  <div class="content">
    <h1 class="heading-1">Welcome!</h1>
  </div>
</div>