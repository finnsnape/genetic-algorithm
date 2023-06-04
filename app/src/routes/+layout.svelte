<script>
  import "$lib/scss/app.scss";
  import Header from "$lib/components/Header.svelte";
</script>

<style lang="scss">
  .wrapper {
      display: flex;
      flex-direction: column;
      align-items: stretch;
      background-color: #fff;
      padding: 0 $space-xl-3xl;
      max-width: 100%;
      min-height: 100vh;
      margin-bottom: $size-070;
  }
</style>

<div class="wrapper">
  <Header></Header>
  <slot />
</div>
