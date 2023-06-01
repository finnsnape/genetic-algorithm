<script lang="ts">
  import LearnMenu from "$lib/components/learn/LearnMenu.svelte";

  export let title: string;
  export let image: string;
  export let paragraphs: string[];
  export let resources: string[];
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: row;
    column-gap: $size-070;
  }

  .content {
    display: flex;
    flex-direction: column;
    row-gap: $size-070;
  }

  .content__info {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    &__paragraphs {
      display: flex;
      flex-direction: column;
      row-gap: $size-020;
    }
    img {
      margin-top: $size-040;
    }
  }

  .content__resources {
  }
</style>

<div class="wrapper">
  <LearnMenu></LearnMenu>
  <div class="content">
    <div class="content__info">
      <h1 class="heading-1">{title}</h1>
      <div class="content__info__paragraphs">
        {#each paragraphs as paragraph}
        <p class="paragraph-1">{paragraph}</p>
        {/each}
      </div>
      <img alt="" src={`/learn/${image}`}>
    </div>
    <div class="content__resources">
      <details>
        <summary class="paragraph-2">External resources</summary>
        <ul>
          {#each resources as resource}
          <li class="paragraph-2"><a class="button button--tertiary" href="{resource}" target="_blank">{resource}</a></li>
          {/each}
        </ul>
      </details>
    </div>
  </div>
</div>