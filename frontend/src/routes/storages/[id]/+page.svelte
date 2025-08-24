<script lang="ts">
  import { onMount } from "svelte";
  import type { Item } from "$lib/types/models";
  import { createItem, listStorageItems } from "$lib/api/services";
  import Layout from "$lib/components/Layout.svelte";
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

<Layout>
  <h1>Storage Items</h1>

  <ItemForm {storageId} on:addItem={handleAddItem} />

  <div>
    {#each items as item}
      <ItemCard {item} />
    {/each}
  </div>
</Layout>