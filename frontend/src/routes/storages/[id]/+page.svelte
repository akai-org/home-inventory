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

<h1 class="mb-4">Storage: {storageId}</h1>

<div class="row">
  <div class="col-md-4">
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">Add New Item</h5>
        <ItemForm {storageId} on:addItem={handleAddItem} />
      </div>
    </div>
  </div>
  <div class="col-md-8">
    <div class="row">
      {#each items as item}
        <div class="col-md-6 col-lg-4 mb-4">
          <ItemCard {item} />
        </div>
      {/each}
    </div>
  </div>
</div>
