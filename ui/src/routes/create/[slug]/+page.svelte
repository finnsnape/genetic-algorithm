<script lang="ts">
    import Create from "$lib/components/Create.svelte";
    /** @type {import('./$types').PageData} */  export let data;
</script>

<Create categoryTitle={data.categoryTitle} functionDeclaration={data.functionDeclaration} code={data.code} nickname={data.nickname} populationContent={data.populationContent} individualContent={data.individualContent} newFunction={data.newFunction}></Create>