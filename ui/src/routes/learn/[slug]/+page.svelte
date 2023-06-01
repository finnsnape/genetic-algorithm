<script lang="ts">
  import LearnSlide from "$lib/components/learn/LearnSlide.svelte";
  /** @type {import('./$types').PageData} */  export let data;
  console.log(data);
</script>

<LearnSlide title={data.title} image={data.image} paragraphs={data.paragraphs} resources={data.resources}></LearnSlide>