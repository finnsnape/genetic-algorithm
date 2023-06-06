<script lang="ts">
   import { onMount } from 'svelte';
   import { page } from '$app/stores';
   import { createCategory } from "$lib/ts/stores";
   import CreateEditor from '$lib/components/create/CreateEditor.svelte';
   /** @type {import('./$types').PageData} */  export let data;

   onMount(async () => {
    createCategory.set(data.createCategory);
   });

  async function saveFunction() {
    const response = await fetch($page.url.pathname, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify($createCategory)
    });
  }
</script>

<style lang="scss">
  .wrapper {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  .content {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
</style>

<div class="wrapper">
  <h1 class="heading-1">Create {data.createCategory.title}</h1>
  <div class="content stack">
    <label class="paragraph-2" for="nickname">Nickname</label>
    {#if $createCategory && $createCategory.currentFunction}
    <input type="text" name="nickname" bind:value={$createCategory.currentFunction.nickname}>
    {/if}
    <p class="paragraph-2">Code:</p>
    <div class="editor-wrapper">
      <CreateEditor />
    </div>
    <button class="button button--primary" on:click|preventDefault={saveFunction}>Save</button>
  </div>
</div>