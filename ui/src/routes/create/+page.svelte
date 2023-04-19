<script lang="ts">
    /** @type {import('./$types').PageData} */  export let data;
</script>

<div class="stack">
    {#each Object.entries(data) as [name, func]}
        <div>
            <h2>{func.title}s</h2>
            <ul>
                <li>N/A</li>
            </ul>
            <a href="/create/{name}">Create</a>
        </div>
    {/each}
</div>
