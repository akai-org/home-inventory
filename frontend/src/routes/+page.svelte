<script lang="ts">
	import { onMount } from 'svelte';
	import type { Item } from '$lib/types/models';
	import { listItems, createItem } from '$lib/api/services';

	let items: Item[] = [];
	let newItem: Partial<Item> = {};

	onMount(async () => {
		items = await listItems();
	});

	async function handleAddItem() {
		if (newItem.name && newItem.type && newItem.description && newItem.storage_id) {
			const createdItem = await createItem(newItem as Item);
			items = [...items, createdItem];
			newItem = {};
		}
	}
</script>

<h1>Home Inventory</h1>

<form on:submit|preventDefault={handleAddItem}>
	<input type="text" placeholder="Name" bind:value={newItem.name} />
	<input type="text" placeholder="Type" bind:value={newItem.type} />
	<input type="text" placeholder="Description" bind:value={newItem.description} />
	<input type="text" placeholder="Storage ID" bind:value={newItem.storage_id} />
	<button type="submit">Add Item</button>
</form>

<table>
	<thead>
		<tr>
			<th>Name</th>
			<th>Type</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		{#each items as item}
			<tr>
				<td>{item.name}</td>
				<td>{item.type}</td>
				<td>{item.description}</td>
			</tr>
		{/each}
	</tbody>
</table>
