<script lang="ts">
    import Create from "$lib/components/Create.svelte";
    /** @type {import('./$types').PageData} */  export let data;
</script>

<Create func={data}></Create>