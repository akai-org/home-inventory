<script lang="ts">
	import { onMount } from 'svelte';
	import type { Item } from '$lib/types/models';
	import { listItems, createItem } from '$lib/api/services';
  import Layout from '$lib/components/Layout.svelte';
  import ItemCard from '$lib/components/ItemCard.svelte';
  import ItemForm from '$lib/components/ItemForm.svelte';
  import ItemFilter from '$lib/components/ItemFilter.svelte';

	let items: Item[] = [];

	onMount(async () => {
		items = await listItems();
	});

	async function handleAddItem(event: CustomEvent<Item>) {
    const newItem = event.detail;
		if (newItem.name && newItem.type && newItem.description && newItem.storage_id) {
			const createdItem = await createItem(newItem as Item);
			items = [...items, createdItem];
		}
	}

  async function handleFilter(event: CustomEvent<{ name: string, type: string, tags: string }>) {
    const { name, type, tags } = event.detail;
    const filterParams = new URLSearchParams();
    if (name) {
      filterParams.append("name", name);
    }
    if (type) {
      filterParams.append("type", type);
    }
    if (tags) {
      filterParams.append("tags", tags);
    }
    items = await listItems(filterParams);
  }
</script>

<Layout>
  <h1>Home Inventory</h1>

  <ItemForm on:addItem={handleAddItem} />

  <ItemFilter on:filter={handleFilter} />

  <div>
    {#each items as item}
      <ItemCard {item} />
    {/each}
  </div>
</Layout>
