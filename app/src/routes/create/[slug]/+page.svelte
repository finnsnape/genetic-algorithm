<script lang="ts">
    import Create from "$lib/components/create/Create.svelte";
    /** @type {import('./$types').PageData} */  export let data;
    // TODO: do we need this to be a component?
</script>

<Create createFunc={data.createFunc} populationContent={data.populationContent} individualContent={data.individualContent}></Create>