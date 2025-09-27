<script lang="ts">
  import { onMount } from "svelte";
  import type { Item } from "$lib/types/models";
  import { createItem, listStorageItems } from "$lib/api/services";
  import ItemCard from "$lib/components/ItemCard.svelte";
  import { page } from "$app/stores";
  import ItemForm from "$lib/components/ItemForm.svelte";

  let items: Item[] = [];
  const storageId = $page.params.id;

  onMount(async () => {
    items = await listStorageItems(storageId);
  });

  async function handleAddItem(event: CustomEvent<Item>) {
    const newItem = await createItem(event.detail as Item);
    items = [...items, newItem];
  }
</script>

<div class="page-header">
  <h1 class="storage-title">Storage: {storageId}</h1>
</div>

<div class="page-layout">
  <div class="form-column">
    <div class="form-container">
      <h5 class="form-title">Add New Item</h5>
      <ItemForm {storageId} on:addItem={handleAddItem} />
    </div>
  </div>
  <div class="items-column">
    <div class="item-grid">
      {#each items as item}
        <ItemCard {item} />
      {/each}
    </div>
  </div>
</div>

<style>
  .page-header {
    margin-bottom: 2rem;
  }

  .storage-title {
    font-size: 2rem;
    font-weight: 700;
  }

  .page-layout {
    display: grid;
    grid-template-columns: 320px 1fr;
    gap: 2rem;
  }

  .form-container {
    background-color: #ffffff;
    padding: 1.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  }

  .form-title {
    font-size: 1.2rem;
    font-weight: 600;
    margin-bottom: 1rem;
  }

  .item-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 1.5rem;
  }
</style>
