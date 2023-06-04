<script lang="ts">
    import Create from "$lib/components/create/Create.svelte";
    /** @type {import('./$types').PageData} */  export let data;
    // TODO: do we need this to be a component?
</script>

<Create categoryTitle={data.categoryTitle} functionDeclaration={data.functionDeclaration} code={data.code} nickname={data.nickname} populationContent={data.populationContent} individualContent={data.individualContent} newFunction={data.newFunction}></Create>