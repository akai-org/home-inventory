<script lang="ts">
	import { onMount } from 'svelte';
	import type { Item, Storage } from '$lib/types/models';
	import { listItems, createItem, listStorages } from '$lib/api/services';
  import ItemCard from '$lib/components/ItemCard.svelte';
  import ItemForm from '$lib/components/ItemForm.svelte';
  import ItemFilter from '$lib/components/ItemFilter.svelte';
  import StorageList from '$lib/components/StorageList.svelte';

	let items: Item[] = [];
  let storages: Storage[] = [];
  let showAddItemModal = false;

	onMount(async () => {
		items = await listItems();
    storages = await listStorages();
	});

	async function handleAddItem(event: CustomEvent<Item>) {
    const newItem = event.detail;
		if (newItem.name && newItem.type && newItem.description && newItem.storage_id) {
			const createdItem = await createItem(newItem as Item);
			items = await listItems();
      showAddItemModal = false;
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


<div class="page-layout">
  <div class="sidebar-column">
    <div class="filter-section">
      <h5 class="section-title"><i class="bi bi-funnel-fill me-2"></i>Filters</h5>
      <ItemFilter on:filter={handleFilter} />
    </div>
    <div class="storage-section">
      <h5 class="section-title"><i class="bi bi-hdd-stack-fill me-2"></i>Storages</h5>
      <StorageList />
    </div>
  </div>
  <div class="main-column">
    <div class="header">
      <h1 class="page-title">Inventory</h1>
      <div class="actions">
        <div class="search-bar">
          <i class="bi bi-search"></i>
          <input type="text" placeholder="Search by name..." on:input={(e) => handleFilter({ detail: { name: e.currentTarget.value, type: '', tags: '' }, bubbles: false, cancelable: false, composed: false,timeStamp: 0, isTrusted: false, type: ''})}>
        </div>
        <button class="add-item-button" on:click={() => showAddItemModal = true}><i class="bi bi-plus-lg me-2"></i>Add Item</button>
      </div>
    </div>
    <div class="item-grid">
      {#each items as item}
        <ItemCard {item} />
      {/each}
    </div>
  </div>
</div>

{#if showAddItemModal}
  <div class="modal-overlay" on:click={() => showAddItemModal = false}>
    <div class="modal-content" on:click|stopPropagation>
      <div class="modal-header">
        <h5 class="modal-title">Add New Item</h5>
        <button type="button" class="close-button" on:click={() => showAddItemModal = false}>&times;</button>
      </div>
      <div class="modal-body">
        <ItemForm on:addItem={handleAddItem} {storages} />
      </div>
    </div>
  </div>
{/if}

<style>
  .page-layout {
    display: grid;
    grid-template-columns: 280px 1fr;
    gap: 2rem;
  }

  .sidebar-column .section-title {
    font-size: 1.2rem;
    font-weight: 600;
    margin-bottom: 1rem;
  }

  .filter-section, .storage-section {
    background-color: #ffffff;
    padding: 1.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    margin-bottom: 1.5rem;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }

  .page-title {
    font-size: 2rem;
    font-weight: 700;
  }

  .actions {
    display: flex;
    gap: 1rem;
  }

  .search-bar {
    display: flex;
    align-items: center;
    background-color: #fff;
    border-radius: 8px;
    padding: 0.5rem 1rem;
    box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  }

  .search-bar input {
    border: none;
    outline: none;
    margin-left: 0.5rem;
  }

  .add-item-button {
    background-color: #4a90e2;
    color: white;
    border: none;
    padding: 0.8rem 1.2rem;
    border-radius: 8px;
    font-weight: 500;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .add-item-button:hover {
    background-color: #357abd;
  }

  .item-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 1.5rem;
  }

  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.6);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  .modal-content {
    background-color: #fff;
    border-radius: 12px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.2);
    width: 90%;
    max-width: 500px;
  }

  .modal-header {
    padding: 1.5rem;
    border-bottom: 1px solid #eee;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .modal-title {
    font-size: 1.25rem;
    font-weight: 600;
  }

  .close-button {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
    color: #888;
  }

  .modal-body {
    padding: 1.5rem;
  }
</style>
