<script lang="ts">
  import { createEventDispatcher, onMount } from "svelte";
  import type { Item } from "$lib/types/models";

  export let storageId: string;

  const dispatch = createEventDispatcher();

  let newItem: Partial<Item> = {};
  let tags = "";

  onMount(() => {
    newItem.storage_id = storageId;
  });

  function handleSubmit() {
    if (tags) {
      newItem.tags = tags.split(",").map(t => t.trim());
    }
    dispatch("addItem", newItem);
    newItem = {
      storage_id: storageId,
    };
    tags = "";
  }
</script>

<form on:submit|preventDefault={handleSubmit}>
  <input type="text" placeholder="Name" bind:value={newItem.name} />
  <input type="text" placeholder="Type" bind:value={newItem.type} />
  <input type="text" placeholder="Description" bind:value={newItem.description} />
  <input type="text" placeholder="Storage ID" bind:value={newItem.storage_id} disabled />
  <input type="text" placeholder="Tags (comma separated)" bind:value={tags} />
  <button type="submit">Add Item</button>
</form>