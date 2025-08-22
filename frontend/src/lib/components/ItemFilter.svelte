<script lang="ts">
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  let name = "";
  let type = "";
  let tags = "";

  function handleFilter() {
    dispatch("filter", { name, type, tags });
  }
</script>

<form on:submit|preventDefault={handleFilter}>
  <input type="text" placeholder="Name" bind:value={name} />
  <input type="text" placeholder="Type" bind:value={type} />
  <input type="text" placeholder="Tags (comma separated)" bind:value={tags} />
  <button type="submit">Filter</button>
</form>